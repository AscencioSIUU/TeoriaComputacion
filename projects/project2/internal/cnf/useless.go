package cnf

import (
	"proyecto-cyk/internal/grammar"
)

// UselessEliminator elimina símbolos inútiles
type UselessEliminator struct{}

// NewUselessEliminator crea un nuevo eliminador de símbolos inútiles
func NewUselessEliminator() *UselessEliminator {
	return &UselessEliminator{}
}

// Eliminate elimina símbolos inútiles de la gramática
// Símbolos inútiles:
// 1. No generadores: no pueden derivar a terminales
// 2. No alcanzables: no se pueden alcanzar desde el símbolo inicial
func (ue *UselessEliminator) Eliminate(g *grammar.Grammar) *grammar.Grammar {
	// Paso 1: Eliminar símbolos no generadores
	generating := ue.findGenerating(g)
	g = ue.removeNonGenerating(g, generating)

	// Paso 2: Eliminar símbolos no alcanzables
	reachable := ue.findReachable(g)
	g = ue.removeNonReachable(g, reachable)

	return g
}

// findGenerating encuentra todos los símbolos generadores
// Un símbolo A es generador si:
// - A -> a (deriva directamente a terminal), O
// - A -> B C donde B y C son generadores
func (ue *UselessEliminator) findGenerating(g *grammar.Grammar) map[string]bool {
	generating := make(map[string]bool)

	// Todos los terminales son generadores
	for _, t := range g.GetTerminals() {
		generating[t.Key()] = true
	}

	// Iterar hasta que no haya cambios
	changed := true
	for changed {
		changed = false

		for _, prod := range g.Productions {
			// Si ya es generador, skip
			if generating[prod.Left.Key()] {
				continue
			}

			// Verificar si todos los símbolos del lado derecho son generadores
			allGenerating := true
			for _, symbol := range prod.Right {
				if !generating[symbol.Key()] {
					allGenerating = false
					break
				}
			}

			// Si todos son generadores (o es epsilon), este símbolo es generador
			if allGenerating {
				generating[prod.Left.Key()] = true
				changed = true
			}
		}
	}

	return generating
}

// removeNonGenerating elimina símbolos no generadores
func (ue *UselessEliminator) removeNonGenerating(g *grammar.Grammar, generating map[string]bool) *grammar.Grammar {
	newGrammar := grammar.NewGrammar(g.StartSymbol)

	// Solo mantener producciones donde todos los símbolos son generadores
	for _, prod := range g.Productions {
		// Verificar lado izquierdo
		if !generating[prod.Left.Key()] {
			continue
		}

		// Verificar lado derecho
		allGenerating := true
		for _, symbol := range prod.Right {
			if !generating[symbol.Key()] {
				allGenerating = false
				break
			}
		}

		if allGenerating {
			newGrammar.AddProduction(prod.Clone())
		}
	}

	return newGrammar
}

// findReachable encuentra todos los símbolos alcanzables desde el símbolo inicial
func (ue *UselessEliminator) findReachable(g *grammar.Grammar) map[string]bool {
	reachable := make(map[string]bool)

	// El símbolo inicial es alcanzable
	reachable[g.StartSymbol.Key()] = true

	// Iterar hasta que no haya cambios
	changed := true
	for changed {
		changed = false

		for _, prod := range g.Productions {
			// Si el lado izquierdo no es alcanzable, skip
			if !reachable[prod.Left.Key()] {
				continue
			}

			// Marcar todos los símbolos del lado derecho como alcanzables
			for _, symbol := range prod.Right {
				if !reachable[symbol.Key()] {
					reachable[symbol.Key()] = true
					changed = true
				}
			}
		}
	}

	return reachable
}

// removeNonReachable elimina símbolos no alcanzables
func (ue *UselessEliminator) removeNonReachable(g *grammar.Grammar, reachable map[string]bool) *grammar.Grammar {
	newGrammar := grammar.NewGrammar(g.StartSymbol)

	// Solo mantener producciones donde todos los símbolos son alcanzables
	for _, prod := range g.Productions {
		// Verificar lado izquierdo
		if !reachable[prod.Left.Key()] {
			continue
		}

		// Verificar lado derecho
		allReachable := true
		for _, symbol := range prod.Right {
			if !reachable[symbol.Key()] {
				allReachable = false
				break
			}
		}

		if allReachable {
			newGrammar.AddProduction(prod.Clone())
		}
	}

	return newGrammar
}

// EliminateUseless es una función helper para uso directo
func EliminateUseless(g *grammar.Grammar) *grammar.Grammar {
	eliminator := NewUselessEliminator()
	return eliminator.Eliminate(g)
}
