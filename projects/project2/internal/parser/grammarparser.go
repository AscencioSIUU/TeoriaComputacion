package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"proyecto-cyk/internal/grammar"
)

// GrammarParser parsea archivos de gramática
type GrammarParser struct {
	filename string
}

// NewGrammarParser crea un nuevo parser de gramática
func NewGrammarParser(filename string) *GrammarParser {
	return &GrammarParser{
		filename: filename,
	}
}

// Parse parsea el archivo de gramática
func (gp *GrammarParser) Parse() (*grammar.Grammar, error) {
	file, err := os.Open(gp.filename)
	if err != nil {
		return nil, fmt.Errorf("error al abrir archivo: %w", err)
	}
	defer file.Close()

	var g *grammar.Grammar
	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		// Ignorar líneas vacías y comentarios
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") {
			continue
		}

		// Parsear la línea
		prods, err := gp.parseLine(line)
		if err != nil {
			return nil, fmt.Errorf("error en línea %d: %w", lineNum, err)
		}

		// La primera producción define el símbolo inicial
		if g == nil {
			if len(prods) == 0 {
				return nil, fmt.Errorf("línea %d: no se encontraron producciones", lineNum)
			}
			g = grammar.NewGrammar(prods[0].Left)
		}

		// Agregar todas las producciones
		for _, prod := range prods {
			g.AddProduction(prod)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error al leer archivo: %w", err)
	}

	if g == nil {
		return nil, fmt.Errorf("el archivo no contiene producciones válidas")
	}

	return g, nil
}

// parseLine parsea una línea de gramática
// Formato: "E -> T X | F Y"
func (gp *GrammarParser) parseLine(line string) ([]*grammar.Production, error) {
	// Dividir por "->"
	parts := strings.Split(line, "->")
	if len(parts) != 2 {
		return nil, fmt.Errorf("formato inválido: debe contener '->'")
	}

	// Lado izquierdo
	leftStr := strings.TrimSpace(parts[0])
	if leftStr == "" {
		return nil, fmt.Errorf("lado izquierdo vacío")
	}
	left := grammar.NewSymbol(leftStr)
	if !left.IsNonTerminal() {
		return nil, fmt.Errorf("lado izquierdo debe ser no-terminal: %s", leftStr)
	}

	// Lado derecho (puede tener múltiples alternativas con |)
	rightStr := strings.TrimSpace(parts[1])
	if rightStr == "" {
		return nil, fmt.Errorf("lado derecho vacío")
	}

	// Dividir por "|" para alternativas
	alternatives := strings.Split(rightStr, "|")
	productions := make([]*grammar.Production, 0, len(alternatives))

	for _, alt := range alternatives {
		right, err := gp.parseRightSide(alt)
		if err != nil {
			return nil, err
		}
		prod := grammar.NewProduction(left, right)
		productions = append(productions, prod)
	}

	return productions, nil
}

// parseRightSide parsea el lado derecho de una producción
// Ejemplos: "T X", "id", "e" (epsilon)
func (gp *GrammarParser) parseRightSide(text string) ([]grammar.Symbol, error) {
	text = strings.TrimSpace(text)
	if text == "" {
		return nil, fmt.Errorf("lado derecho vacío")
	}

	// Caso especial: epsilon
	if text == "e" || text == "ε" || text == "epsilon" {
		return []grammar.Symbol{}, nil
	}

	// Dividir por espacios
	tokens := strings.Fields(text)
	symbols := make([]grammar.Symbol, 0, len(tokens))

	for _, token := range tokens {
		token = strings.TrimSpace(token)
		if token == "" {
			continue
		}
		// Crear símbolo (detecta automáticamente tipo)
		symbols = append(symbols, grammar.NewSymbol(token))
	}

	if len(symbols) == 0 {
		return nil, fmt.Errorf("no se encontraron símbolos válidos")
	}

	return symbols, nil
}

// ParseFromString parsea una gramática desde un string (útil para tests)
func ParseFromString(text string) (*grammar.Grammar, error) {
	var g *grammar.Grammar
	lines := strings.Split(text, "\n")

	gp := &GrammarParser{}

	for lineNum, line := range lines {
		line = strings.TrimSpace(line)

		// Ignorar líneas vacías y comentarios
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") {
			continue
		}

		// Parsear la línea
		prods, err := gp.parseLine(line)
		if err != nil {
			return nil, fmt.Errorf("error en línea %d: %w", lineNum+1, err)
		}

		// La primera producción define el símbolo inicial
		if g == nil {
			if len(prods) == 0 {
				return nil, fmt.Errorf("línea %d: no se encontraron producciones", lineNum+1)
			}
			g = grammar.NewGrammar(prods[0].Left)
		}

		// Agregar todas las producciones
		for _, prod := range prods {
			g.AddProduction(prod)
		}
	}

	if g == nil {
		return nil, fmt.Errorf("el texto no contiene producciones válidas")
	}

	return g, nil
}
