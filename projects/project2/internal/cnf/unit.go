package cnf

import (
	"proyecto-cyk/internal/grammar"
)

// UnitEliminator elimina producciones unitarias
type UnitEliminator struct{}

// NewUnitEliminator crea un nuevo eliminador de producciones unitarias
func NewUnitEliminator() *UnitEliminator {
	return &UnitEliminator{}
}

// Eliminate elimina todas las producciones unitarias de la gramática
// Una producción unitaria tiene la forma A -> B
func (ue *UnitEliminator) Eliminate(g *grammar.Grammar) *grammar.Grammar {
	// Paso 1: Calcular pares unitarios (A, B) donde A ->* B
	unitPairs := ue.computeUnitPairs(g)

	// Paso 2: Generar nuevas producciones
	newGrammar := grammar.NewGrammar(g.StartSymbol)

	// Para cada par (A, B), si B -> α es no-unitaria, agregar A -> α
	for _, nonTerminal := range g.GetNonTerminals() {
		A := nonTerminal.Key()

		// Para cada símbolo B alcanzable desde A
		for B := range unitPairs[A] {
			// Obtener producciones no-unitarias de B
			prods := g.GetProductionsFor(grammar.NewNonTerminal(B))
			for _, prod := range prods {
				// Solo agregar si NO es unitaria
				if !prod.IsUnit() {
					newProd := grammar.NewProduction(nonTerminal, prod.Right)
					if !newGrammar.HasProduction(newProd) {
						newGrammar.AddProduction(newProd)
					}
				}
			}
		}
	}

	return newGrammar
}

// computeUnitPairs calcula todos los pares (A, B) donde A ->* B
// usando cerradura transitiva
func (ue *UnitEliminator) computeUnitPairs(g *grammar.Grammar) map[string]map[string]bool {
	pairs := make(map[string]map[string]bool)

	// Inicializar: cada símbolo puede derivar a sí mismo
	for _, nt := range g.GetNonTerminals() {
		pairs[nt.Key()] = make(map[string]bool)
		pairs[nt.Key()][nt.Key()] = true
	}

	// Agregar pares directos de producciones unitarias
	for _, prod := range g.Productions {
		if prod.IsUnit() {
			A := prod.Left.Key()
			B := prod.Right[0].Key()
			pairs[A][B] = true
		}
	}

	// Calcular cerradura transitiva (Floyd-Warshall)
	// Si A ->* B y B ->* C, entonces A ->* C
	nonTerminals := g.GetNonTerminals()
	for range nonTerminals {
		for _, ntI := range nonTerminals {
			i := ntI.Key()
			for _, ntJ := range nonTerminals {
				j := ntJ.Key()
				for _, ntK := range nonTerminals {
					k := ntK.Key()
					if pairs[i][j] && pairs[j][k] {
						pairs[i][k] = true
					}
				}
			}
		}
	}

	return pairs
}

// EliminateUnit es una función helper para uso directo
func EliminateUnit(g *grammar.Grammar) *grammar.Grammar {
	eliminator := NewUnitEliminator()
	return eliminator.Eliminate(g)
}
