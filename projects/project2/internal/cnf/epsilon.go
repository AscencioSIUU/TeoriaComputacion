package cnf

import (
	"proyecto-cyk/internal/grammar"
)

// EpsilonEliminator elimina producciones epsilon
type EpsilonEliminator struct{}

// NewEpsilonEliminator crea un nuevo eliminador de epsilon
func NewEpsilonEliminator() *EpsilonEliminator {
	return &EpsilonEliminator{}
}

// Eliminate elimina todas las producciones epsilon de la gramática
func (ee *EpsilonEliminator) Eliminate(g *grammar.Grammar) *grammar.Grammar {
	// Paso 1: Encontrar símbolos anulables
	nullable := ee.findNullable(g)

	// Paso 2: Generar nuevas producciones sin epsilon
	newGrammar := ee.generateWithoutEpsilon(g, nullable)

	return newGrammar
}

// findNullable encuentra todos los símbolos anulables
// Un símbolo A es anulable si:
// - Tiene producción A -> ε, O
// - Tiene producción A -> B C D donde B, C, D son todos anulables
func (ee *EpsilonEliminator) findNullable(g *grammar.Grammar) map[string]bool {
	nullable := make(map[string]bool)

	// Iterar hasta que no haya cambios
	changed := true
	for changed {
		changed = false

		for _, prod := range g.Productions {
			// Si ya es anulable, skip
			if nullable[prod.Left.Key()] {
				continue
			}

			// Caso 1: A -> ε
			if prod.IsEpsilon() {
				nullable[prod.Left.Key()] = true
				changed = true
				continue
			}

			// Caso 2: A -> B C D donde todos son anulables
			allNullable := true
			for _, symbol := range prod.Right {
				if symbol.IsTerminal() {
					allNullable = false
					break
				}
				if !nullable[symbol.Key()] {
					allNullable = false
					break
				}
			}

			if allNullable && len(prod.Right) > 0 {
				nullable[prod.Left.Key()] = true
				changed = true
			}
		}
	}

	return nullable
}

// generateWithoutEpsilon genera todas las producciones sin epsilon
func (ee *EpsilonEliminator) generateWithoutEpsilon(g *grammar.Grammar, nullable map[string]bool) *grammar.Grammar {
	newGrammar := grammar.NewGrammar(g.StartSymbol)

	// Para cada producción, generar todas las variantes
	for _, prod := range g.Productions {
		// No agregar producciones epsilon
		if prod.IsEpsilon() {
			continue
		}

		// Generar variantes
		variants := ee.generateVariants(prod, nullable)
		for _, variant := range variants {
			// No agregar epsilon a menos que sea el símbolo inicial
			if variant.IsEpsilon() && !variant.Left.Equals(g.StartSymbol) {
				continue
			}
			// Evitar duplicados
			if !newGrammar.HasProduction(variant) {
				newGrammar.AddProduction(variant)
			}
		}
	}

	return newGrammar
}

// generateVariants genera todas las variantes de una producción
// considerando símbolos anulables
func (ee *EpsilonEliminator) generateVariants(prod *grammar.Production, nullable map[string]bool) []*grammar.Production {
	variants := make([]*grammar.Production, 0)

	// Encontrar posiciones de símbolos anulables
	nullablePos := make([]int, 0)
	for i, symbol := range prod.Right {
		if nullable[symbol.Key()] {
			nullablePos = append(nullablePos, i)
		}
	}

	// Si no hay símbolos anulables, retornar la producción original
	if len(nullablePos) == 0 {
		variants = append(variants, prod.Clone())
		return variants
	}

	// Generar todas las combinaciones (2^n variantes)
	numVariants := 1 << len(nullablePos) // 2^n

	for i := 0; i < numVariants; i++ {
		// Determinar qué símbolos incluir en esta variante
		exclude := make(map[int]bool)
		for j := 0; j < len(nullablePos); j++ {
			// Si el bit j está encendido, excluir ese símbolo
			if (i & (1 << j)) != 0 {
				exclude[nullablePos[j]] = true
			}
		}

		// Construir el lado derecho de la variante
		right := make([]grammar.Symbol, 0)
		for idx, symbol := range prod.Right {
			if !exclude[idx] {
				right = append(right, symbol)
			}
		}

		// Crear la variante
		variant := grammar.NewProduction(prod.Left, right)
		variants = append(variants, variant)
	}

	return variants
}

// Eliminate es una función helper para uso directo
func Eliminate(g *grammar.Grammar) *grammar.Grammar {
	eliminator := NewEpsilonEliminator()
	return eliminator.Eliminate(g)
}
