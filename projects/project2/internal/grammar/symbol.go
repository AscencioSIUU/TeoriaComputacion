package grammar

import "strings"

// SymbolType define el tipo de símbolo
type SymbolType int

const (
	Terminal SymbolType = iota
	NonTerminal
)

// Symbol representa un símbolo en la gramática
type Symbol struct {
	Value string
	Type  SymbolType
}

// NewTerminal crea un nuevo símbolo terminal
func NewTerminal(value string) Symbol {
	return Symbol{
		Value: value,
		Type:  Terminal,
	}
}

// NewNonTerminal crea un nuevo símbolo no-terminal
func NewNonTerminal(value string) Symbol {
	return Symbol{
		Value: value,
		Type:  NonTerminal,
	}
}

// NewSymbol crea un símbolo detectando automáticamente su tipo
// Regla: Mayúsculas o múltiples caracteres mayúsculas = no-terminal
func NewSymbol(value string) Symbol {
	if isNonTerminalValue(value) {
		return NewNonTerminal(value)
	}
	return NewTerminal(value)
}

// isNonTerminalValue determina si un valor es no-terminal
// No-terminales: empiezan con mayúscula o son nombres como "Det", "VP"
func isNonTerminalValue(value string) bool {
	if len(value) == 0 {
		return false
	}
	// Si empieza con mayúscula, es no-terminal
	firstChar := rune(value[0])
	return firstChar >= 'A' && firstChar <= 'Z'
}

// IsTerminal verifica si el símbolo es terminal
func (s Symbol) IsTerminal() bool {
	return s.Type == Terminal
}

// IsNonTerminal verifica si el símbolo es no-terminal
func (s Symbol) IsNonTerminal() bool {
	return s.Type == NonTerminal
}

// String convierte el símbolo a string
func (s Symbol) String() string {
	return s.Value
}

// Equals compara dos símbolos
func (s Symbol) Equals(other Symbol) bool {
	return s.Value == other.Value && s.Type == other.Type
}

// Key retorna una clave única para usar en maps
func (s Symbol) Key() string {
	return s.Value
}

// Clone crea una copia del símbolo
func (s Symbol) Clone() Symbol {
	return Symbol{
		Value: s.Value,
		Type:  s.Type,
	}
}

// IsEpsilon verifica si el símbolo representa epsilon
func (s Symbol) IsEpsilon() bool {
	return s.Value == "e" || s.Value == "ε" || s.Value == "epsilon"
}

// SymbolsToString convierte un slice de símbolos a string
func SymbolsToString(symbols []Symbol) string {
	if len(symbols) == 0 {
		return "ε"
	}
	parts := make([]string, len(symbols))
	for i, s := range symbols {
		parts[i] = s.String()
	}
	return strings.Join(parts, " ")
}

// SymbolsEqual compara dos slices de símbolos
func SymbolsEqual(a, b []Symbol) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !a[i].Equals(b[i]) {
			return false
		}
	}
	return true
}
