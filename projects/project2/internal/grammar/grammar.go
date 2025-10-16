package grammar

import (
	"fmt"
	"sort"
	"strings"
)

// Grammar representa una gramática libre de contexto
type Grammar struct {
	StartSymbol   Symbol
	Terminals     map[string]Symbol
	NonTerminals  map[string]Symbol
	Productions   []*Production
	productionMap map[string][]*Production // Índice para búsqueda rápida
}

// NewGrammar crea una nueva gramática
func NewGrammar(start Symbol) *Grammar {
	g := &Grammar{
		StartSymbol:   start,
		Terminals:     make(map[string]Symbol),
		NonTerminals:  make(map[string]Symbol),
		Productions:   make([]*Production, 0),
		productionMap: make(map[string][]*Production),
	}
	g.NonTerminals[start.Key()] = start
	return g
}

// AddProduction agrega una producción a la gramática
func (g *Grammar) AddProduction(prod *Production) {
	// Agregar la producción
	g.Productions = append(g.Productions, prod)

	// Registrar el símbolo izquierdo
	g.NonTerminals[prod.Left.Key()] = prod.Left

	// Registrar símbolos del lado derecho
	for _, symbol := range prod.Right {
		if symbol.IsTerminal() {
			g.Terminals[symbol.Key()] = symbol
		} else {
			g.NonTerminals[symbol.Key()] = symbol
		}
	}

	// Actualizar índice
	key := prod.Left.Key()
	g.productionMap[key] = append(g.productionMap[key], prod)
}

// AddProductions agrega múltiples producciones
func (g *Grammar) AddProductions(prods []*Production) {
	for _, prod := range prods {
		g.AddProduction(prod)
	}
}

// GetProductionsFor obtiene todas las producciones que empiezan con un símbolo
// Ejemplo: GetProductionsFor(E) retorna todas las "E -> ..."
func (g *Grammar) GetProductionsFor(symbol Symbol) []*Production {
	return g.productionMap[symbol.Key()]
}

// GetProductionsWith obtiene producciones con lado derecho específico
// Ejemplo: GetProductionsWith(B, C) retorna todas las "A -> B C"
func (g *Grammar) GetProductionsWith(symbols ...Symbol) []*Production {
	result := make([]*Production, 0)
	for _, prod := range g.Productions {
		if len(prod.Right) == len(symbols) {
			match := true
			for i, sym := range symbols {
				if !prod.Right[i].Equals(sym) {
					match = false
					break
				}
			}
			if match {
				result = append(result, prod)
			}
		}
	}
	return result
}

// GetProductionsGenerating obtiene producciones que generan un terminal
// Ejemplo: GetProductionsGenerating("id") retorna "F -> id"
func (g *Grammar) GetProductionsGenerating(terminal string) []*Production {
	result := make([]*Production, 0)
	for _, prod := range g.Productions {
		if len(prod.Right) == 1 && prod.Right[0].IsTerminal() && prod.Right[0].Value == terminal {
			result = append(result, prod)
		}
	}
	return result
}

// IsCNF verifica si toda la gramática está en CNF
func (g *Grammar) IsCNF() bool {
	for _, prod := range g.Productions {
		if !prod.IsCNF() {
			return false
		}
	}
	return true
}

// GetNonTerminals retorna lista de no-terminales
func (g *Grammar) GetNonTerminals() []Symbol {
	result := make([]Symbol, 0, len(g.NonTerminals))
	for _, nt := range g.NonTerminals {
		result = append(result, nt)
	}
	// Ordenar para consistencia
	sort.Slice(result, func(i, j int) bool {
		return result[i].Value < result[j].Value
	})
	return result
}

// GetTerminals retorna lista de terminales
func (g *Grammar) GetTerminals() []Symbol {
	result := make([]Symbol, 0, len(g.Terminals))
	for _, t := range g.Terminals {
		result = append(result, t)
	}
	// Ordenar para consistencia
	sort.Slice(result, func(i, j int) bool {
		return result[i].Value < result[j].Value
	})
	return result
}

// Clone crea una copia profunda de la gramática
func (g *Grammar) Clone() *Grammar {
	newGrammar := NewGrammar(g.StartSymbol.Clone())

	// Copiar producciones
	for _, prod := range g.Productions {
		newGrammar.AddProduction(prod.Clone())
	}

	return newGrammar
}

// String convierte la gramática a string
func (g *Grammar) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Start Symbol: %s\n", g.StartSymbol))
	sb.WriteString(fmt.Sprintf("Non-Terminals: %v\n", g.GetNonTerminalsString()))
	sb.WriteString(fmt.Sprintf("Terminals: %v\n", g.GetTerminalsString()))
	sb.WriteString("Productions:\n")

	// Agrupar producciones por símbolo izquierdo
	grouped := g.groupProductionsByLeft()
	for _, left := range g.getSortedNonTerminals() {
		prods := grouped[left]
		if len(prods) > 0 {
			sb.WriteString(fmt.Sprintf("  %s -> ", left))
			rights := make([]string, len(prods))
			for i, prod := range prods {
				rights[i] = SymbolsToString(prod.Right)
			}
			sb.WriteString(strings.Join(rights, " | "))
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

// GetNonTerminalsString retorna string de no-terminales
func (g *Grammar) GetNonTerminalsString() string {
	nts := g.GetNonTerminals()
	parts := make([]string, len(nts))
	for i, nt := range nts {
		parts[i] = nt.Value
	}
	return strings.Join(parts, ", ")
}

// GetTerminalsString retorna string de terminales
func (g *Grammar) GetTerminalsString() string {
	ts := g.GetTerminals()
	parts := make([]string, len(ts))
	for i, t := range ts {
		parts[i] = t.Value
	}
	return strings.Join(parts, ", ")
}

// groupProductionsByLeft agrupa producciones por símbolo izquierdo
func (g *Grammar) groupProductionsByLeft() map[string][]*Production {
	grouped := make(map[string][]*Production)
	for _, prod := range g.Productions {
		key := prod.Left.Value
		grouped[key] = append(grouped[key], prod)
	}
	return grouped
}

// getSortedNonTerminals retorna lista ordenada de no-terminales
func (g *Grammar) getSortedNonTerminals() []string {
	keys := make([]string, 0, len(g.NonTerminals))
	for k := range g.NonTerminals {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// HasProduction verifica si existe una producción específica
func (g *Grammar) HasProduction(prod *Production) bool {
	for _, p := range g.Productions {
		if p.Equals(prod) {
			return true
		}
	}
	return false
}

// RemoveProduction elimina una producción
func (g *Grammar) RemoveProduction(prod *Production) {
	newProds := make([]*Production, 0)
	for _, p := range g.Productions {
		if !p.Equals(prod) {
			newProds = append(newProds, p)
		}
	}
	g.Productions = newProds
	g.rebuildIndex()
}

// rebuildIndex reconstruye el índice de producciones
func (g *Grammar) rebuildIndex() {
	g.productionMap = make(map[string][]*Production)
	for _, prod := range g.Productions {
		key := prod.Left.Key()
		g.productionMap[key] = append(g.productionMap[key], prod)
	}
}

// ProductionCount retorna el número de producciones
func (g *Grammar) ProductionCount() int {
	return len(g.Productions)
}
