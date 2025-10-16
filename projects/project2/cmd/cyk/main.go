package main

import (
	"flag"
	"fmt"
	"os"

	"proyecto-cyk/internal/cnf"
	"proyecto-cyk/internal/cyk"
	"proyecto-cyk/internal/parser"
	"proyecto-cyk/pkg/utils"
)

func main() {
	// Definir flags
	grammarFile := flag.String("grammar", "", "Archivo de gramática (requerido)")
	inputStr := flag.String("input", "", "Cadena de entrada (requerido)")
	verbose := flag.Bool("verbose", false, "Mostrar salida detallada")
	showTable := flag.Bool("table", false, "Mostrar tabla CYK completa")
	showOriginal := flag.Bool("show-original", false, "Mostrar gramática original")
	showCNF := flag.Bool("show-cnf", false, "Mostrar gramática en CNF")

	flag.Parse()

	// Validar argumentos
	if *grammarFile == "" || *inputStr == "" {
		fmt.Println("Uso: cyk --grammar <archivo> --input <cadena>")
		fmt.Println()
		fmt.Println("Opciones:")
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println("Ejemplos:")
		fmt.Println("  cyk --grammar examples/gramaticas/1.txt --input \"she eats a cake\"")
		fmt.Println("  cyk --grammar examples/gramaticas/1.txt --input \"she eats a cake with a fork\" --verbose")
		os.Exit(1)
	}

	// Ejecutar el programa
	if err := run(*grammarFile, *inputStr, *verbose, *showTable, *showOriginal, *showCNF); err != nil {
		utils.PrintError(err)
		os.Exit(1)
	}
}

func run(grammarFile, inputStr string, verbose, showTable, showOriginal, showCNF bool) error {
	// Paso 1: Parsear gramática
	if verbose {
		utils.PrintInfo(fmt.Sprintf("Leyendo gramática desde: %s", grammarFile))
	}

	grammarParser := parser.NewGrammarParser(grammarFile)
	originalGrammar, err := grammarParser.Parse()
	if err != nil {
		return fmt.Errorf("error al parsear gramática: %w", err)
	}

	if showOriginal {
		utils.PrintGrammar(originalGrammar, "GRAMÁTICA ORIGINAL")
	}

	// Paso 2: Convertir a CNF
	if verbose {
		utils.PrintInfo("Convirtiendo gramática a Forma Normal de Chomsky (CNF)...")
	}

	converter := cnf.NewConverter()
	cnfGrammar := converter.ConvertToCNF(originalGrammar)

	if showCNF {
		utils.PrintGrammar(cnfGrammar, "GRAMÁTICA EN CNF")
	}

	if verbose {
		utils.PrintComparison(originalGrammar, cnfGrammar)
	}

	// Verificar que la conversión fue exitosa
	if !cnfGrammar.IsCNF() {
		return fmt.Errorf("error: la gramática no está en CNF después de la conversión")
	}

	// Paso 3: Parsear entrada
	if verbose {
		utils.PrintInfo(fmt.Sprintf("Parseando entrada: \"%s\"", inputStr))
	}

	inputParser := parser.NewInputParser()
	tokens := inputParser.Parse(inputStr)

	if len(tokens) == 0 {
		return fmt.Errorf("error: la entrada está vacía después del parsing")
	}

	if verbose {
		fmt.Printf("Tokens: %v\n", tokens)
	}

	// Paso 4: Ejecutar algoritmo CYK
	if verbose {
		utils.PrintInfo("Ejecutando algoritmo CYK...")
	}

	cykAlgorithm := cyk.NewCYK(cnfGrammar)
	result, err := cykAlgorithm.Parse(tokens)
	if err != nil {
		return fmt.Errorf("error al ejecutar CYK: %w", err)
	}

	// Paso 5: Mostrar resultados
	utils.PrintResult(result, tokens, cykAlgorithm)

	if showTable {
		utils.PrintTable(result.Table, tokens, verbose)
	}

	// Resumen final
	utils.PrintHeader("RESUMEN")
	if result.Accepted {
		fmt.Println("✓ SÍ - La frase es sintácticamente correcta según la gramática")
	} else {
		fmt.Println("✗ NO - La frase NO es válida según la gramática")
	}
	fmt.Printf("\nTiempo de validación: %s\n", utils.FormatDuration(result.ExecutionTime))
	fmt.Println()

	return nil
}
