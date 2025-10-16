package utils

import (
	"fmt"
	"strings"

	"proyecto-cyk/internal/cyk"
	"proyecto-cyk/internal/grammar"
)

// PrintHeader imprime un encabezado con formato
func PrintHeader(title string) {
	separator := strings.Repeat("=", 80)
	fmt.Println()
	fmt.Println(separator)
	fmt.Println(title)
	fmt.Println(separator)
	fmt.Println()
}

// PrintSubHeader imprime un subencabezado
func PrintSubHeader(title string) {
	separator := strings.Repeat("-", 60)
	fmt.Println()
	fmt.Println(title)
	fmt.Println(separator)
}

// PrintGrammar imprime una gramática de forma legible
func PrintGrammar(g *grammar.Grammar, title string) {
	PrintHeader(title)

	fmt.Printf("Símbolo Inicial: %s\n\n", g.StartSymbol.Value)

	fmt.Println("No-Terminales:")
	nts := g.GetNonTerminals()
	for i, nt := range nts {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(nt.Value)
	}
	fmt.Println()

	fmt.Println("\nTerminales:")
	ts := g.GetTerminals()
	for i, t := range ts {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(t.Value)
	}
	fmt.Println()

	fmt.Printf("\nTotal de Producciones: %d\n\n", g.ProductionCount())

	PrintSubHeader("Producciones")

	// Agrupar producciones por símbolo izquierdo
	grouped := make(map[string][]*grammar.Production)
	for _, prod := range g.Productions {
		key := prod.Left.Value
		grouped[key] = append(grouped[key], prod)
	}

	// Ordenar e imprimir
	for _, nt := range nts {
		prods := grouped[nt.Value]
		if len(prods) > 0 {
			fmt.Printf("  %s -> ", nt.Value)
			rights := make([]string, len(prods))
			for i, prod := range prods {
				rights[i] = grammar.SymbolsToString(prod.Right)
			}
			fmt.Println(strings.Join(rights, " | "))
		}
	}

	fmt.Println()
}

// PrintComparison compara gramática original vs CNF
func PrintComparison(original, cnf *grammar.Grammar) {
	PrintHeader("COMPARACIÓN: GRAMÁTICA ORIGINAL vs CNF")

	fmt.Printf("%-30s | %-30s\n", "ORIGINAL", "CNF")
	fmt.Println(strings.Repeat("-", 62))
	fmt.Printf("%-30s | %-30s\n",
		fmt.Sprintf("Producciones: %d", original.ProductionCount()),
		fmt.Sprintf("Producciones: %d", cnf.ProductionCount()))
	fmt.Printf("%-30s | %-30s\n",
		fmt.Sprintf("No-Terminales: %d", len(original.GetNonTerminals())),
		fmt.Sprintf("No-Terminales: %d", len(cnf.GetNonTerminals())))
	fmt.Printf("%-30s | %-30s\n",
		fmt.Sprintf("En CNF: %v", original.IsCNF()),
		fmt.Sprintf("En CNF: %v", cnf.IsCNF()))

	fmt.Println()
}

// PrintResult imprime el resultado del algoritmo CYK
func PrintResult(result *cyk.CYKResult, tokens []string, cykAlg *cyk.CYK) {
	PrintHeader("RESULTADO DEL ANÁLISIS CYK")

	fmt.Printf("Cadena analizada: %s\n", strings.Join(tokens, " "))
	fmt.Printf("Número de tokens: %d\n\n", len(tokens))

	if result.Accepted {
		fmt.Println("✓ ACEPTADA - La cadena pertenece al lenguaje")
	} else {
		fmt.Println("✗ RECHAZADA - La cadena NO pertenece al lenguaje")
	}

	fmt.Printf("\nTiempo de ejecución: %s\n", FormatDuration(result.ExecutionTime))

	if result.Accepted && result.ParseTree != nil {
		PrintSubHeader("Árbol de Parsing")
		tree := cykAlg.GetParseTree(result.ParseTree, tokens, 0)
		fmt.Println(tree)
	}

	fmt.Println()
}

// PrintTable imprime la tabla CYK
func PrintTable(table *cyk.Table, tokens []string, detailed bool) {
	PrintSubHeader("Tabla CYK")

	if detailed {
		fmt.Print(table.StringDetailed(tokens))
	} else {
		fmt.Print(table.String())
	}
}

// PrintError imprime un mensaje de error
func PrintError(err error) {
	fmt.Println()
	fmt.Println("❌ ERROR:")
	fmt.Println(err.Error())
	fmt.Println()
}

// PrintSuccess imprime un mensaje de éxito
func PrintSuccess(message string) {
	fmt.Println()
	fmt.Printf("✓ %s\n", message)
	fmt.Println()
}

// PrintWarning imprime una advertencia
func PrintWarning(message string) {
	fmt.Println()
	fmt.Printf("⚠ ADVERTENCIA: %s\n", message)
	fmt.Println()
}

// PrintInfo imprime información general
func PrintInfo(message string) {
	fmt.Println()
	fmt.Printf("ℹ %s\n", message)
	fmt.Println()
}
