package cyk

import (
	"fmt"
	"proyecto-cyk/internal/grammar"
	"time"
)

// CYKResult representa el resultado del algoritmo CYK
type CYKResult struct {
	Accepted      bool
	Table         *Table
	ExecutionTime time.Duration
	ParseTree     *CellEntry // Raíz del árbol de parsing
}

// CYK implementa el algoritmo Cocke-Younger-Kasami
type CYK struct {
	Grammar *grammar.Grammar
}

// NewCYK crea una nueva instancia del algoritmo CYK
func NewCYK(g *grammar.Grammar) *CYK {
	return &CYK{
		Grammar: g,
	}
}

// Parse ejecuta el algoritmo CYK sobre una cadena de entrada
func (cyk *CYK) Parse(tokens []string) (*CYKResult, error) {
	startTime := time.Now()

	// Verificar que la gramática esté en CNF
	if !cyk.Grammar.IsCNF() {
		return nil, fmt.Errorf("la gramática debe estar en Forma Normal de Chomsky (CNF)")
	}

	// Verificar entrada válida
	if len(tokens) == 0 {
		return nil, fmt.Errorf("la entrada no puede estar vacía")
	}

	n := len(tokens)
	table := NewTable(n)

	// Paso 1: Llenar diagonal (caso base)
	cyk.fillDiagonal(table, tokens)

	// Paso 2: Llenar tabla (caso recursivo)
	cyk.fillTable(table, tokens)

	// Paso 3: Verificar aceptación
	topCell := table.GetTopCell()
	accepted := topCell.Contains(cyk.Grammar.StartSymbol)

	executionTime := time.Since(startTime)

	// Obtener parse tree si fue aceptada
	var parseTree *CellEntry
	if accepted {
		parseTree = topCell.Get(cyk.Grammar.StartSymbol)
	}

	result := &CYKResult{
		Accepted:      accepted,
		Table:         table,
		ExecutionTime: executionTime,
		ParseTree:     parseTree,
	}

	return result, nil
}

// fillDiagonal llena la diagonal de la tabla (caso base)
// Table[i][i] contiene símbolos que generan el i-ésimo terminal
func (cyk *CYK) fillDiagonal(table *Table, tokens []string) {
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		cell := table.Get(i, i)

		// Buscar producciones A -> token
		prods := cyk.Grammar.GetProductionsGenerating(token)
		for _, prod := range prods {
			entry := NewCellEntry(prod.Left, prod)
			cell.Add(entry)
		}
	}
}

// fillTable llena el resto de la tabla (caso recursivo)
func (cyk *CYK) fillTable(table *Table, tokens []string) {
	n := len(tokens)

	// Para cada longitud de subcadena (de 2 hasta n)
	for length := 2; length <= n; length++ {
		// Para cada posición inicial
		for i := 0; i <= n-length; i++ {
			j := i + length - 1 // Posición final

			// Para cada punto de división k
			for k := i; k < j; k++ {
				cyk.processPartition(table, i, j, k)
			}
		}
	}
}

// processPartition procesa una partición específica
// Busca producciones A -> BC donde:
//   - B está en Table[i][k]
//   - C está en Table[k+1][j]
// Y agrega A a Table[i][j]
func (cyk *CYK) processPartition(table *Table, i, j, k int) {
	leftCell := table.Get(i, k)
	rightCell := table.Get(k+1, j)
	currentCell := table.Get(i, j)

	// Para cada símbolo B en la celda izquierda
	for _, leftSymbol := range leftCell.GetSymbols() {
		leftEntry := leftCell.Get(leftSymbol)

		// Para cada símbolo C en la celda derecha
		for _, rightSymbol := range rightCell.GetSymbols() {
			rightEntry := rightCell.Get(rightSymbol)

			// Buscar producciones A -> B C
			prods := cyk.Grammar.GetProductionsWith(leftSymbol, rightSymbol)
			for _, prod := range prods {
				// Agregar A a la celda actual si no existe
				if !currentCell.Contains(prod.Left) {
					entry := NewCellEntryWithChildren(
						prod.Left,
						prod,
						k,
						leftEntry,
						rightEntry,
					)
					currentCell.Add(entry)
				}
			}
		}
	}
}

// GetParseTree retorna el árbol de parsing en formato string
func (cyk *CYK) GetParseTree(entry *CellEntry, tokens []string, depth int) string {
	if entry == nil {
		return ""
	}

	indent := ""
	for i := 0; i < depth; i++ {
		indent += "  "
	}

	result := fmt.Sprintf("%s%s", indent, entry.Symbol.Value)

	if entry.IsLeaf() {
		// Es una hoja (terminal)
		if entry.Production != nil && len(entry.Production.Right) > 0 {
			result += fmt.Sprintf(" -> %s", entry.Production.Right[0].Value)
		}
		result += "\n"
	} else {
		// Es un nodo interno
		result += "\n"
		if entry.LeftChild != nil {
			result += cyk.GetParseTree(entry.LeftChild, tokens, depth+1)
		}
		if entry.RightChild != nil {
			result += cyk.GetParseTree(entry.RightChild, tokens, depth+1)
		}
	}

	return result
}

// ValidateGrammar verifica que la gramática sea válida para CYK
func (cyk *CYK) ValidateGrammar() error {
	if cyk.Grammar == nil {
		return fmt.Errorf("la gramática no puede ser nula")
	}

	if !cyk.Grammar.IsCNF() {
		return fmt.Errorf("la gramática debe estar en CNF")
	}

	if cyk.Grammar.ProductionCount() == 0 {
		return fmt.Errorf("la gramática no tiene producciones")
	}

	return nil
}
