package grammar

import (
	"fmt"
	"strings"
)

// Production representa una producción en la gramática (A -> α)
type Production struct {
	Left  Symbol   // Lado izquierdo (siempre no-terminal)
	Right []Symbol // Lado derecho (puede ser vacío para ε)
}

// NewProduction crea una nueva producción
func NewProduction(left Symbol, right []Symbol) *Production {
	return &Production{
		Left:  left,
		Right: right,
	}
}

// IsEpsilon verifica si es una producción epsilon (A → ε)
func (p *Production) IsEpsilon() bool {
	return len(p.Right) == 0
}

// IsUnit verifica si es una producción unitaria (A → B)
func (p *Production) IsUnit() bool {
	return len(p.Right) == 1 && p.Right[0].IsNonTerminal()
}

// IsCNF verifica si la producción está en Forma Normal de Chomsky
// CNF permite:
// - A → a (un terminal)
// - A → BC (dos no-terminales)
func (p *Production) IsCNF() bool {
	// No puede ser epsilon
	if p.IsEpsilon() {
		return false
	}

	// Caso 1: A → a (un terminal)
	if len(p.Right) == 1 {
		return p.Right[0].IsTerminal()
	}

	// Caso 2: A → BC (exactamente dos no-terminales)
	if len(p.Right) == 2 {
		return p.Right[0].IsNonTerminal() && p.Right[1].IsNonTerminal()
	}

	// Cualquier otra forma no es CNF
	return false
}

// Length retorna la longitud del lado derecho
func (p *Production) Length() int {
	return len(p.Right)
}

// String convierte la producción a string (formato: A -> B C)
func (p *Production) String() string {
	right := SymbolsToString(p.Right)
	return fmt.Sprintf("%s -> %s", p.Left.String(), right)
}

// Equals compara dos producciones
func (p *Production) Equals(other *Production) bool {
	if other == nil {
		return false
	}
	if !p.Left.Equals(other.Left) {
		return false
	}
	return SymbolsEqual(p.Right, other.Right)
}

// Clone crea una copia profunda de la producción
func (p *Production) Clone() *Production {
	rightCopy := make([]Symbol, len(p.Right))
	copy(rightCopy, p.Right)
	return &Production{
		Left:  p.Left.Clone(),
		Right: rightCopy,
	}
}

// HasSymbol verifica si el lado derecho contiene un símbolo específico
func (p *Production) HasSymbol(symbol Symbol) bool {
	for _, s := range p.Right {
		if s.Equals(symbol) {
			return true
		}
	}
	return false
}

// HasTerminal verifica si el lado derecho contiene algún terminal
func (p *Production) HasTerminal() bool {
	for _, s := range p.Right {
		if s.IsTerminal() {
			return true
		}
	}
	return false
}

// HasOnlyTerminals verifica si el lado derecho solo contiene terminales
func (p *Production) HasOnlyTerminals() bool {
	if len(p.Right) == 0 {
		return false
	}
	for _, s := range p.Right {
		if !s.IsTerminal() {
			return false
		}
	}
	return true
}

// HasOnlyNonTerminals verifica si el lado derecho solo contiene no-terminales
func (p *Production) HasOnlyNonTerminals() bool {
	if len(p.Right) == 0 {
		return false
	}
	for _, s := range p.Right {
		if !s.IsNonTerminal() {
			return false
		}
	}
	return true
}

// RightAsString retorna el lado derecho como string concatenado
func (p *Production) RightAsString() string {
	parts := make([]string, len(p.Right))
	for i, s := range p.Right {
		parts[i] = s.Value
	}
	return strings.Join(parts, " ")
}

// Key retorna una clave única para la producción
func (p *Production) Key() string {
	return fmt.Sprintf("%s->%s", p.Left.Key(), p.RightAsString())
}
