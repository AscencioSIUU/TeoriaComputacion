package test

import (
	"proyecto-cyk/internal/cnf"
	"proyecto-cyk/internal/cyk"
	"proyecto-cyk/internal/parser"
	"testing"
)

func TestCYKSimpleAccepted(t *testing.T) {
	// Simple grammar: S -> a
	grammarText := `
S -> a
`

	g, err := parser.ParseFromString(grammarText)
	if err != nil {
		t.Fatalf("Error parsing grammar: %v", err)
	}

	// Already in CNF
	cykAlg := cyk.NewCYK(g)
	result, err := cykAlg.Parse([]string{"a"})
	if err != nil {
		t.Fatalf("Error running CYK: %v", err)
	}

	if !result.Accepted {
		t.Error("Expected string to be accepted")
	}

	t.Logf("Execution time: %v", result.ExecutionTime)
}

func TestCYKSimpleRejected(t *testing.T) {
	// Simple grammar: S -> a
	grammarText := `
S -> a
`

	g, err := parser.ParseFromString(grammarText)
	if err != nil {
		t.Fatalf("Error parsing grammar: %v", err)
	}

	// Test with "b" - should be rejected
	cykAlg := cyk.NewCYK(g)
	result, err := cykAlg.Parse([]string{"b"})
	if err != nil {
		t.Fatalf("Error running CYK: %v", err)
	}

	if result.Accepted {
		t.Error("Expected string to be rejected")
	}
}

func TestCYKBinaryProduction(t *testing.T) {
	// Grammar: S -> A B, A -> a, B -> b
	grammarText := `
S -> A B
A -> a
B -> b
`

	g, err := parser.ParseFromString(grammarText)
	if err != nil {
		t.Fatalf("Error parsing grammar: %v", err)
	}

	cykAlg := cyk.NewCYK(g)

	// Test "a b" - should be accepted
	result, err := cykAlg.Parse([]string{"a", "b"})
	if err != nil {
		t.Fatalf("Error running CYK: %v", err)
	}

	if !result.Accepted {
		t.Error("Expected 'a b' to be accepted")
	}

	// Test "a" - should be rejected
	result2, err := cykAlg.Parse([]string{"a"})
	if err != nil {
		t.Fatalf("Error running CYK: %v", err)
	}

	if result2.Accepted {
		t.Error("Expected 'a' to be rejected")
	}
}

func TestCYKWithRealGrammar(t *testing.T) {
	// Load grammar from file
	gp := parser.NewGrammarParser("testdata/1.txt")
	g, err := gp.Parse()
	if err != nil {
		t.Fatalf("Error parsing grammar file: %v", err)
	}

	// Convert to CNF
	converter := cnf.NewConverter()
	cnfG := converter.ConvertToCNF(g)

	if !cnfG.IsCNF() {
		t.Fatal("Grammar is not in CNF")
	}

	t.Logf("CNF Grammar has %d productions", cnfG.ProductionCount())

	cykAlg := cyk.NewCYK(cnfG)

	// Test cases
	tests := []struct {
		input    string
		expected bool
		desc     string
	}{
		{"she eats", true, "simple valid sentence (VP can be just a verb)"},
		{"she eats a cake", true, "valid sentence"},
		{"she eats a cake with a fork", true, "valid sentence with PP"},
		{"he drinks the beer", true, "valid sentence"},
		{"the cat eats the cake", true, "valid sentence with Det"},
		{"a dog drinks a juice", true, "valid sentence"},
		{"she he eats", false, "invalid sentence"},
		{"eats", false, "only verb"},
		{"cake", false, "only noun"},
	}

	inputParser := parser.NewInputParser()

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			tokens := inputParser.Parse(test.input)
			result, err := cykAlg.Parse(tokens)
			if err != nil {
				t.Fatalf("Error running CYK: %v", err)
			}

			if result.Accepted != test.expected {
				t.Errorf("Input '%s': expected accepted=%v, got=%v",
					test.input, test.expected, result.Accepted)
			}

			t.Logf("'%s': accepted=%v, time=%v",
				test.input, result.Accepted, result.ExecutionTime)
		})
	}
}

func TestCYKTable(t *testing.T) {
	// Simple grammar to check table
	grammarText := `
S -> A B
A -> a
B -> b
`

	g, err := parser.ParseFromString(grammarText)
	if err != nil {
		t.Fatalf("Error parsing grammar: %v", err)
	}

	cykAlg := cyk.NewCYK(g)
	result, err := cykAlg.Parse([]string{"a", "b"})
	if err != nil {
		t.Fatalf("Error running CYK: %v", err)
	}

	// Check table structure
	if result.Table.Size != 2 {
		t.Errorf("Expected table size 2, got %d", result.Table.Size)
	}

	// Check diagonal [0][0] contains A
	cell00 := result.Table.Get(0, 0)
	if !cell00.ContainsKey("A") {
		t.Error("Cell [0][0] should contain A")
	}

	// Check diagonal [1][1] contains B
	cell11 := result.Table.Get(1, 1)
	if !cell11.ContainsKey("B") {
		t.Error("Cell [1][1] should contain B")
	}

	// Check top cell [0][1] contains S
	topCell := result.Table.Get(0, 1)
	if !topCell.ContainsKey("S") {
		t.Error("Cell [0][1] should contain S")
	}
}

func TestCYKParseTree(t *testing.T) {
	grammarText := `
S -> A B
A -> a
B -> b
`

	g, err := parser.ParseFromString(grammarText)
	if err != nil {
		t.Fatalf("Error parsing grammar: %v", err)
	}

	cykAlg := cyk.NewCYK(g)
	result, err := cykAlg.Parse([]string{"a", "b"})
	if err != nil {
		t.Fatalf("Error running CYK: %v", err)
	}

	if !result.Accepted {
		t.Fatal("String should be accepted")
	}

	if result.ParseTree == nil {
		t.Fatal("Parse tree should not be nil for accepted string")
	}

	// Get parse tree as string
	tree := cykAlg.GetParseTree(result.ParseTree, []string{"a", "b"}, 0)
	t.Logf("Parse tree:\n%s", tree)

	if tree == "" {
		t.Error("Parse tree string should not be empty")
	}
}

func TestCYKNonCNFGrammar(t *testing.T) {
	// Grammar not in CNF
	grammarText := `
S -> A B C
A -> a
B -> b
C -> c
`

	g, err := parser.ParseFromString(grammarText)
	if err != nil {
		t.Fatalf("Error parsing grammar: %v", err)
	}

	// Try to run CYK without converting to CNF
	cykAlg := cyk.NewCYK(g)
	_, err = cykAlg.Parse([]string{"a", "b", "c"})

	if err == nil {
		t.Error("Expected error when running CYK on non-CNF grammar")
	}
}

func TestCYKEmptyInput(t *testing.T) {
	grammarText := `
S -> a
`

	g, err := parser.ParseFromString(grammarText)
	if err != nil {
		t.Fatalf("Error parsing grammar: %v", err)
	}

	cykAlg := cyk.NewCYK(g)
	_, err = cykAlg.Parse([]string{})

	if err == nil {
		t.Error("Expected error when parsing empty input")
	}
}
