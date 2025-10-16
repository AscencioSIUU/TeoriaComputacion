package test

import (
	"proyecto-cyk/internal/cnf"
	"proyecto-cyk/internal/grammar"
	"proyecto-cyk/internal/parser"
	"testing"
)

func TestEpsilonElimination(t *testing.T) {
	// Create grammar with epsilon productions
	grammarText := `
S -> A B
A -> a | e
B -> b
`

	g, err := parser.ParseFromString(grammarText)
	if err != nil {
		t.Fatalf("Error parsing grammar: %v", err)
	}

	// Eliminate epsilon
	eliminator := cnf.NewEpsilonEliminator()
	newG := eliminator.Eliminate(g)

	// Verify no epsilon productions remain (except possibly for start symbol)
	for _, prod := range newG.Productions {
		if prod.IsEpsilon() && !prod.Left.Equals(g.StartSymbol) {
			t.Errorf("Found epsilon production: %s", prod.String())
		}
	}

	t.Logf("Original: %d productions, After epsilon elimination: %d productions",
		g.ProductionCount(), newG.ProductionCount())
}

func TestUnitElimination(t *testing.T) {
	// Create grammar with unit productions
	grammarText := `
S -> A
A -> B
B -> a
`

	g, err := parser.ParseFromString(grammarText)
	if err != nil {
		t.Fatalf("Error parsing grammar: %v", err)
	}

	// Eliminate unit productions
	eliminator := cnf.NewUnitEliminator()
	newG := eliminator.Eliminate(g)

	// Verify no unit productions remain
	for _, prod := range newG.Productions {
		if prod.IsUnit() {
			t.Errorf("Found unit production: %s", prod.String())
		}
	}

	// Should have S -> a and A -> a and B -> a
	t.Logf("Original: %d productions, After unit elimination: %d productions",
		g.ProductionCount(), newG.ProductionCount())
}

func TestUselessElimination(t *testing.T) {
	// Create grammar with useless symbols
	grammarText := `
S -> A B
A -> a
B -> C
C -> D
`

	g, err := parser.ParseFromString(grammarText)
	if err != nil {
		t.Fatalf("Error parsing grammar: %v", err)
	}

	// Eliminate useless symbols
	eliminator := cnf.NewUselessEliminator()
	newG := eliminator.Eliminate(g)

	// C and D are not generating, B is not generating, so they should be removed
	t.Logf("Original: %d productions, After useless elimination: %d productions",
		g.ProductionCount(), newG.ProductionCount())

	// Should only have S -> A B and A -> a (or less if B was removed)
	if newG.ProductionCount() > 2 {
		t.Errorf("Expected at most 2 productions, got %d", newG.ProductionCount())
	}
}

func TestCNFConversion(t *testing.T) {
	// Load grammar from file
	gp := parser.NewGrammarParser("testdata/1.txt")
	g, err := gp.Parse()
	if err != nil {
		t.Fatalf("Error parsing grammar file: %v", err)
	}

	t.Logf("Original grammar: %d productions", g.ProductionCount())

	// Convert to CNF
	converter := cnf.NewConverter()
	cnfG := converter.ConvertToCNF(g)

	t.Logf("CNF grammar: %d productions", cnfG.ProductionCount())

	// Verify all productions are in CNF
	if !cnfG.IsCNF() {
		t.Error("Grammar is not in CNF after conversion")

		// Show which productions are not CNF
		for _, prod := range cnfG.Productions {
			if !prod.IsCNF() {
				t.Logf("Not CNF: %s", prod.String())
			}
		}
	}

	// Verify we still have the start symbol
	if cnfG.StartSymbol.Value != g.StartSymbol.Value {
		t.Errorf("Start symbol changed from %s to %s",
			g.StartSymbol.Value, cnfG.StartSymbol.Value)
	}
}

func TestCNFConversionSteps(t *testing.T) {
	// Simple grammar to test all steps
	grammarText := `
S -> A B C
A -> a | e
B -> b
C -> c
`

	g, err := parser.ParseFromString(grammarText)
	if err != nil {
		t.Fatalf("Error parsing grammar: %v", err)
	}

	t.Logf("Step 0 - Original: %d productions", g.ProductionCount())

	// Step 1: Eliminate epsilon
	g = cnf.Eliminate(g)
	t.Logf("Step 1 - After epsilon elimination: %d productions", g.ProductionCount())

	// Step 2: Eliminate unit
	g = cnf.EliminateUnit(g)
	t.Logf("Step 2 - After unit elimination: %d productions", g.ProductionCount())

	// Step 3: Eliminate useless
	g = cnf.EliminateUseless(g)
	t.Logf("Step 3 - After useless elimination: %d productions", g.ProductionCount())

	// Step 4: Full conversion
	converter := cnf.NewConverter()
	g = converter.ConvertToCNF(grammar.NewGrammar(g.StartSymbol))

	// Re-parse and convert properly
	g, _ = parser.ParseFromString(grammarText)
	g = converter.ConvertToCNF(g)

	t.Logf("Step 4 - Final CNF: %d productions", g.ProductionCount())

	if !g.IsCNF() {
		t.Error("Final grammar is not in CNF")
	}
}
