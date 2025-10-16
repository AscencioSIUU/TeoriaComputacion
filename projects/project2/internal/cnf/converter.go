package cnf

import (
	"fmt"
	"proyecto-cyk/internal/grammar"
)

// Converter convierte gramáticas a Forma Normal de Chomsky (CNF)
type Converter struct {
	symbolCounter int
}

// NewConverter crea un nuevo convertidor
func NewConverter() *Converter {
	return &Converter{
		symbolCounter: 0,
	}
}

// ConvertToCNF convierte una gramática a CNF
// Pasos:
// 1. Eliminar producciones epsilon
// 2. Eliminar producciones unitarias
// 3. Eliminar símbolos inútiles
// 4. Aislar terminales
// 5. Dividir producciones largas
func (c *Converter) ConvertToCNF(g *grammar.Grammar) *grammar.Grammar {
	// Paso 1: Eliminar epsilon
	g = Eliminate(g)

	// Paso 2: Eliminar unitarias
	g = EliminateUnit(g)

	// Paso 3: Eliminar inútiles
	g = EliminateUseless(g)

	// Paso 4: Aislar terminales
	g = c.isolateTerminals(g)

	// Paso 5: Dividir producciones largas
	g = c.splitLongProductions(g)

	return g
}

// isolateTerminals aísla terminales mezclados con no-terminales
// Ejemplo: A -> a B se convierte en:
//   A -> C1 B
//   C1 -> a
func (c *Converter) isolateTerminals(g *grammar.Grammar) *grammar.Grammar {
	newGrammar := grammar.NewGrammar(g.StartSymbol)

	// Mapa para reutilizar símbolos creados para terminales
	terminalSymbols := make(map[string]grammar.Symbol)

	for _, prod := range g.Productions {
		// Si la producción ya está en CNF, mantenerla
		if prod.IsCNF() {
			newGrammar.AddProduction(prod.Clone())
			continue
		}

		// Si es solo un terminal, mantenerla
		if len(prod.Right) == 1 && prod.Right[0].IsTerminal() {
			newGrammar.AddProduction(prod.Clone())
			continue
		}

		// Necesitamos aislar terminales
		newRight := make([]grammar.Symbol, 0)
		for _, symbol := range prod.Right {
			if symbol.IsTerminal() {
				// Crear o reutilizar símbolo para este terminal
				newSym, exists := terminalSymbols[symbol.Key()]
				if !exists {
					// Crear nuevo símbolo
					newSym = c.createNewSymbol("C")
					terminalSymbols[symbol.Key()] = newSym

					// Agregar producción NewSym -> terminal
					newProd := grammar.NewProduction(newSym, []grammar.Symbol{symbol})
					newGrammar.AddProduction(newProd)
				}
				newRight = append(newRight, newSym)
			} else {
				newRight = append(newRight, symbol)
			}
		}

		// Agregar la producción modificada
		newProd := grammar.NewProduction(prod.Left, newRight)
		newGrammar.AddProduction(newProd)
	}

	return newGrammar
}

// splitLongProductions divide producciones largas en binarias
// Ejemplo: A -> B C D se convierte en:
//   A -> B X1
//   X1 -> C D
func (c *Converter) splitLongProductions(g *grammar.Grammar) *grammar.Grammar {
	newGrammar := grammar.NewGrammar(g.StartSymbol)

	for _, prod := range g.Productions {
		// Si ya es CNF o tiene 2 o menos símbolos, mantenerla
		if len(prod.Right) <= 2 {
			newGrammar.AddProduction(prod.Clone())
			continue
		}

		// Dividir en producciones binarias
		// A -> B C D E se convierte en:
		//   A -> B X1
		//   X1 -> C X2
		//   X2 -> D E
		current := prod.Left
		for i := 0; i < len(prod.Right)-2; i++ {
			// Crear nuevo símbolo
			newSym := c.createNewSymbol("X")

			// Crear producción: current -> Right[i] newSym
			newProd := grammar.NewProduction(
				current,
				[]grammar.Symbol{prod.Right[i], newSym},
			)
			newGrammar.AddProduction(newProd)

			// El siguiente símbolo izquierdo es el nuevo símbolo
			current = newSym
		}

		// Última producción: current -> Right[n-2] Right[n-1]
		lastProd := grammar.NewProduction(
			current,
			[]grammar.Symbol{
				prod.Right[len(prod.Right)-2],
				prod.Right[len(prod.Right)-1],
			},
		)
		newGrammar.AddProduction(lastProd)
	}

	return newGrammar
}

// createNewSymbol genera un nuevo símbolo único
func (c *Converter) createNewSymbol(base string) grammar.Symbol {
	c.symbolCounter++
	name := fmt.Sprintf("%s%d", base, c.symbolCounter)
	return grammar.NewNonTerminal(name)
}

// ConvertToCNF es una función helper para uso directo
func ConvertToCNF(g *grammar.Grammar) *grammar.Grammar {
	converter := NewConverter()
	return converter.ConvertToCNF(g)
}
