package test

import (
	"proyecto-cyk/internal/grammar"
	"proyecto-cyk/internal/parser"
	"testing"
)

func TestSymbolCreation(t *testing.T) {
	// Test terminal
	term := grammar.NewTerminal("id")
	if !term.IsTerminal() {
		t.Error("Expected terminal symbol")
	}
	if term.IsNonTerminal() {
		t.Error("Terminal should not be non-terminal")
	}

	// Test non-terminal
	nonTerm := grammar.NewNonTerminal("E")
	if !nonTerm.IsNonTerminal() {
		t.Error("Expected non-terminal symbol")
	}
	if nonTerm.IsTerminal() {
		t.Error("Non-terminal should not be terminal")
	}
}

func TestSymbolAutoDetection(t *testing.T) {
	tests := []struct {
		value       string
		shouldBeNT  bool
		description string
	}{
		{"E", true, "Single uppercase letter"},
		{"S", true, "Start symbol"},
		{"VP", true, "Multiple uppercase letters"},
		{"Det", true, "Capitalized word"},
		{"id", false, "Lowercase word"},
		{"a", false, "Single lowercase letter"},
		{"+", false, "Symbol"},
		{"123", false, "Number"},
	}

	for _, test := range tests {
		sym := grammar.NewSymbol(test.value)
		if sym.IsNonTerminal() != test.shouldBeNT {
			t.Errorf("%s: expected non-terminal=%v, got=%v",
				test.description, test.shouldBeNT, sym.IsNonTerminal())
		}
	}
}

func TestProductionCreation(t *testing.T) {
	left := grammar.NewNonTerminal("E")
	right := []grammar.Symbol{
		grammar.NewNonTerminal("T"),
		grammar.NewNonTerminal("X"),
	}

	prod := grammar.NewProduction(left, right)

	if prod.IsEpsilon() {
		t.Error("Non-epsilon production marked as epsilon")
	}
	if prod.IsUnit() {
		t.Error("Binary production marked as unit")
	}
	if prod.Length() != 2 {
		t.Errorf("Expected length 2, got %d", prod.Length())
	}
}

func TestEpsilonProduction(t *testing.T) {
	left := grammar.NewNonTerminal("A")
	right := []grammar.Symbol{} // Epsilon

	prod := grammar.NewProduction(left, right)

	if !prod.IsEpsilon() {
		t.Error("Epsilon production not detected")
	}
	if prod.Length() != 0 {
		t.Errorf("Epsilon production should have length 0, got %d", prod.Length())
	}
}

func TestUnitProduction(t *testing.T) {
	left := grammar.NewNonTerminal("A")
	right := []grammar.Symbol{grammar.NewNonTerminal("B")}

	prod := grammar.NewProduction(left, right)

	if !prod.IsUnit() {
		t.Error("Unit production not detected")
	}
}

func TestCNFProduction(t *testing.T) {
	// Test A -> a (terminal)
	prod1 := grammar.NewProduction(
		grammar.NewNonTerminal("A"),
		[]grammar.Symbol{grammar.NewTerminal("a")},
	)
	if !prod1.IsCNF() {
		t.Error("A -> a should be CNF")
	}

	// Test A -> BC (two non-terminals)
	prod2 := grammar.NewProduction(
		grammar.NewNonTerminal("A"),
		[]grammar.Symbol{
			grammar.NewNonTerminal("B"),
			grammar.NewNonTerminal("C"),
		},
	)
	if !prod2.IsCNF() {
		t.Error("A -> BC should be CNF")
	}

	// Test A -> BCD (not CNF)
	prod3 := grammar.NewProduction(
		grammar.NewNonTerminal("A"),
		[]grammar.Symbol{
			grammar.NewNonTerminal("B"),
			grammar.NewNonTerminal("C"),
			grammar.NewNonTerminal("D"),
		},
	)
	if prod3.IsCNF() {
		t.Error("A -> BCD should not be CNF")
	}

	// Test A -> aB (not CNF)
	prod4 := grammar.NewProduction(
		grammar.NewNonTerminal("A"),
		[]grammar.Symbol{
			grammar.NewTerminal("a"),
			grammar.NewNonTerminal("B"),
		},
	)
	if prod4.IsCNF() {
		t.Error("A -> aB should not be CNF")
	}
}

func TestGrammarCreation(t *testing.T) {
	start := grammar.NewNonTerminal("S")
	g := grammar.NewGrammar(start)

	// Add production S -> A B
	prod := grammar.NewProduction(
		start,
		[]grammar.Symbol{
			grammar.NewNonTerminal("A"),
			grammar.NewNonTerminal("B"),
		},
	)
	g.AddProduction(prod)

	if g.ProductionCount() != 1 {
		t.Errorf("Expected 1 production, got %d", g.ProductionCount())
	}

	prods := g.GetProductionsFor(start)
	if len(prods) != 1 {
		t.Errorf("Expected 1 production for S, got %d", len(prods))
	}
}

func TestGrammarParser(t *testing.T) {
	grammarText := `
E -> T X
X -> + T X | e
T -> F Y
Y -> * F Y | e
F -> id
`

	g, err := parser.ParseFromString(grammarText)
	if err != nil {
		t.Fatalf("Error parsing grammar: %v", err)
	}

	if g.StartSymbol.Value != "E" {
		t.Errorf("Expected start symbol E, got %s", g.StartSymbol.Value)
	}

	if g.ProductionCount() < 5 {
		t.Errorf("Expected at least 5 productions, got %d", g.ProductionCount())
	}
}

func TestGrammarParserFromFile(t *testing.T) {
	gp := parser.NewGrammarParser("testdata/1.txt")
	g, err := gp.Parse()
	if err != nil {
		t.Fatalf("Error parsing grammar file: %v", err)
	}

	if g.StartSymbol.Value != "S" {
		t.Errorf("Expected start symbol S, got %s", g.StartSymbol.Value)
	}

	if g.ProductionCount() == 0 {
		t.Error("Grammar should have productions")
	}

	t.Logf("Loaded grammar with %d productions", g.ProductionCount())
}
