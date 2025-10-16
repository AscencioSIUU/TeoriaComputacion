package cyk

import (
	"fmt"
	"strings"
)

// Table representa la tabla dinámica del algoritmo CYK
// Es una matriz triangular superior de tamaño n×n
type Table struct {
	Size  int      // Tamaño n (longitud de la cadena)
	Cells [][]*Cell // Matriz triangular
}

// NewTable crea una nueva tabla de tamaño n×n
func NewTable(n int) *Table {
	cells := make([][]*Cell, n)
	for i := 0; i < n; i++ {
		cells[i] = make([]*Cell, n)
		for j := 0; j < n; j++ {
			cells[i][j] = NewCell()
		}
	}

	return &Table{
		Size:  n,
		Cells: cells,
	}
}

// Get obtiene la celda en posición [i][j]
func (t *Table) Get(i, j int) *Cell {
	if i < 0 || i >= t.Size || j < 0 || j >= t.Size {
		return nil
	}
	return t.Cells[i][j]
}

// Set establece la celda en posición [i][j]
func (t *Table) Set(i, j int, cell *Cell) {
	if i >= 0 && i < t.Size && j >= 0 && j < t.Size {
		t.Cells[i][j] = cell
	}
}

// String genera una representación visual de la tabla
func (t *Table) String() string {
	var sb strings.Builder

	// Encabezado
	sb.WriteString("\nTabla CYK:\n")
	sb.WriteString(strings.Repeat("=", 60) + "\n\n")

	// Imprimir la tabla de forma triangular
	// Fila j representa la diagonal j (j=0 es la diagonal principal)
	for j := 0; j < t.Size; j++ {
		// Espacios para alinear
		indent := strings.Repeat("  ", j)
		sb.WriteString(indent)

		// Imprimir celdas de esta diagonal
		for i := 0; i < t.Size-j; i++ {
			cell := t.Get(i, i+j)
			cellStr := cell.String()

			// Limitar ancho para que quepa
			if len(cellStr) > 15 {
				cellStr = cellStr[:12] + "..."
			}

			sb.WriteString(fmt.Sprintf("%-15s ", cellStr))
		}
		sb.WriteString("\n")
	}

	sb.WriteString("\n")
	return sb.String()
}

// StringDetailed genera una representación detallada de la tabla
func (t *Table) StringDetailed(tokens []string) string {
	var sb strings.Builder

	sb.WriteString("\nTabla CYK Detallada:\n")
	sb.WriteString(strings.Repeat("=", 80) + "\n\n")

	// Mostrar tokens
	sb.WriteString("Tokens: ")
	for i, token := range tokens {
		sb.WriteString(fmt.Sprintf("[%d]=%s ", i, token))
	}
	sb.WriteString("\n\n")

	// Mostrar cada celda con detalle
	for j := 0; j < t.Size; j++ {
		for i := 0; i < t.Size-j; i++ {
			cell := t.Get(i, i+j)
			if !cell.IsEmpty() {
				sb.WriteString(fmt.Sprintf("Celda [%d][%d] (tokens %d a %d): %s\n",
					i, i+j, i, i+j, cell.String()))
			}
		}
	}

	sb.WriteString("\n")
	return sb.String()
}

// GetTopCell obtiene la celda superior (resultado final)
func (t *Table) GetTopCell() *Cell {
	if t.Size > 0 {
		return t.Get(0, t.Size-1)
	}
	return nil
}

// IsValid verifica que el tamaño sea válido
func (t *Table) IsValid() bool {
	return t.Size > 0 && len(t.Cells) == t.Size
}
