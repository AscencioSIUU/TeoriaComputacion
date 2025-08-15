// labs/lab4/main.go
// This file is part of the lab4 project for the course on Theory of Computation.
// It implements a command-line tool to read regexes from an input file,
// build their NFAs using Thompson's construction, and generate DOT and PNG files
// for visualization. It also simulates the NFA with a given string to check acceptance.
// It supports regex extensions like Kleene star, union, concatenation, and more.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"lab4/config"
	"lab4/graphviz"
	"lab4/nfa"
	"lab4/regex"
	"lab4/thompson"
)

func main() {
	// Command-line flags for input and output directories
	inPath := flag.String("in", "input.txt", "path to input file")
	dotDir := flag.String("dotout", "dotout", "output directory for DOT files")
	pngDir := flag.String("pngout", "pngout", "output directory for PNG files")
	flag.Parse()

	f, err := os.Open(*inPath)
	if err != nil {
		log.Fatalf("cannot open input file: %v", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	lineNo := 0

	for sc.Scan() {
		lineNo++
		raw := strings.TrimSpace(sc.Text())
		if raw == "" || strings.HasPrefix(raw, "#") {
			continue
		}

		// Enforce "regex;w"
		parts := strings.SplitN(raw, ";", 2)
		if len(parts) != 2 {
			log.Printf("Line %d: invalid format. Expected 'regex;w'. Got: %q\n", lineNo, raw)
			continue
		}
		r := strings.TrimSpace(parts[0])
		w := strings.TrimSpace(parts[1])
		if r == "" { // Check if regex is empty
			log.Printf("Line %d: empty regex before ';'\n", lineNo)
			continue
		}
		if w == "" { // Check if w is empty
			log.Printf("Line %d: empty w after ';'\n", lineNo)
			continue
		}

		// Expand and format the regex
		// This handles extensions like Kleene star, union, concatenation, etc.
		// It also formats the regex to a standard form.
		// The regex is expected to be in infix notation.
		expanded := config.ExpandRegexExtensions(r)
		formatted := config.FormatRegex(expanded)
		postfix := config.InfixToPostfix(formatted)

		fmt.Printf("Line %d\n", lineNo)
		fmt.Printf("  raw       : %s\n", r)
		fmt.Printf("  expanded  : %s\n", expanded)
		fmt.Printf("  formatted : %s\n", formatted)
		fmt.Printf("  postfix   : %s\n", postfix)

		// Build AST from postfix regex
		ast, err := regex.BuildAST(postfix)
		if err != nil {
			log.Printf("  AST error: %v\n\n", err)
			continue
		}

		// Build NFA using Thompson's construction
		nfaObj, err := thompson.Build(ast)
		if err != nil {
			log.Printf("  Thompson error: %v\n\n", err)
			continue
		}

		// Save DOT and PNG
		dotPath := filepath.Join(*dotDir, fmt.Sprintf("nfa_%03d.dot", lineNo))
		pngPath := filepath.Join(*pngDir, fmt.Sprintf("nfa_%03d.png", lineNo))

		if err := graphviz.WriteDOT(nfaObj, dotPath); err != nil {
			log.Printf("  DOT error: %v\n\n", err)
			continue
		}
		fmt.Printf("  DOT saved: %s\n", dotPath)

		if err := graphviz.GeneratePNGFromDot(dotPath, pngPath); err != nil {
			log.Printf("  PNG error (is Graphviz installed?): %v\n\n", err)
		} else {
			fmt.Printf("  PNG saved: %s\n", pngPath)
		}

		// Simulate NFA with the string w
		accepted := nfa.Simulate(nfaObj, w)
		ans := map[bool]string{true: "sí", false: "no"}[accepted]
		fmt.Printf("  w ∈ L(r)? %s   (w = %q)\n\n", ans, w)
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
}
