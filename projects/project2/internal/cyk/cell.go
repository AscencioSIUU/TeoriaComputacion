package cyk

import (
	"proyecto-cyk/internal/grammar"
	"strings"
)

// CellEntry representa una entrada en una celda de la tabla CYK
type CellEntry struct {
	Symbol     grammar.Symbol
	Production *grammar.Production
	SplitPoint int         // Punto de división (k) para reconstruir el árbol
	LeftChild  *CellEntry  // Puntero al hijo izquierdo
	RightChild *CellEntry  // Puntero al hijo derecho
}

// NewCellEntry crea una nueva entrada de celda
func NewCellEntry(symbol grammar.Symbol, production *grammar.Production) *CellEntry {
	return &CellEntry{
		Symbol:     symbol,
		Production: production,
		SplitPoint: -1,
		LeftChild:  nil,
		RightChild: nil,
	}
}

// NewCellEntryWithChildren crea una entrada con hijos (para producciones binarias)
func NewCellEntryWithChildren(symbol grammar.Symbol, production *grammar.Production, splitPoint int, left, right *CellEntry) *CellEntry {
	return &CellEntry{
		Symbol:     symbol,
		Production: production,
		SplitPoint: splitPoint,
		LeftChild:  left,
		RightChild: right,
	}
}

// IsLeaf verifica si es una hoja (producción terminal)
func (ce *CellEntry) IsLeaf() bool {
	return ce.LeftChild == nil && ce.RightChild == nil
}

// Cell representa una celda en la tabla CYK
type Cell struct {
	Entries map[string]*CellEntry // Mapa de símbolo -> entrada
}

// NewCell crea una nueva celda vacía
func NewCell() *Cell {
	return &Cell{
		Entries: make(map[string]*CellEntry),
	}
}

// Add agrega un símbolo a la celda
func (c *Cell) Add(entry *CellEntry) {
	c.Entries[entry.Symbol.Key()] = entry
}

// Contains verifica si la celda contiene un símbolo
func (c *Cell) Contains(symbol grammar.Symbol) bool {
	_, exists := c.Entries[symbol.Key()]
	return exists
}

// ContainsKey verifica si la celda contiene un símbolo por su clave
func (c *Cell) ContainsKey(key string) bool {
	_, exists := c.Entries[key]
	return exists
}

// Get obtiene la entrada de un símbolo
func (c *Cell) Get(symbol grammar.Symbol) *CellEntry {
	return c.Entries[symbol.Key()]
}

// GetKey obtiene la entrada por clave
func (c *Cell) GetKey(key string) *CellEntry {
	return c.Entries[key]
}

// GetSymbols retorna lista de todos los símbolos en la celda
func (c *Cell) GetSymbols() []grammar.Symbol {
	symbols := make([]grammar.Symbol, 0, len(c.Entries))
	for _, entry := range c.Entries {
		symbols = append(symbols, entry.Symbol)
	}
	return symbols
}

// IsEmpty verifica si la celda está vacía
func (c *Cell) IsEmpty() bool {
	return len(c.Entries) == 0
}

// Size retorna el número de símbolos en la celda
func (c *Cell) Size() int {
	return len(c.Entries)
}

// String convierte la celda a string
func (c *Cell) String() string {
	if c.IsEmpty() {
		return "∅"
	}

	symbols := make([]string, 0, len(c.Entries))
	for key := range c.Entries {
		symbols = append(symbols, key)
	}
	return "{" + strings.Join(symbols, ", ") + "}"
}

// Clone crea una copia de la celda
func (c *Cell) Clone() *Cell {
	newCell := NewCell()
	for key, entry := range c.Entries {
		newCell.Entries[key] = entry
	}
	return newCell
}
