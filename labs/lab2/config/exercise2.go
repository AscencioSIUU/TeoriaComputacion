// lab2/config/config.go
package config

import "fmt"

// Stack representa la pila usada en el ejercicio 2.
type Stack []interface{}

// Push apila un valor y lo imprime.
func (s *Stack) Push(value interface{}) {
	fmt.Println("Pila: push:", value)
	*s = append(*s, value)
}

// Pop desapila y devuelve el valor; si está vacía, devuelve (nil,false).
func (s *Stack) Pop() (interface{}, bool) {
	if len(*s) == 0 {
		return nil, false
	}
	i := len(*s) - 1
	v := (*s)[i]
	*s = (*s)[:i]
	fmt.Println("Pila: pop", v)
	return v, true
}

// Peek devuelve el tope sin desapilar.
func (s *Stack) Peek() (interface{}, bool) {
	if len(*s) == 0 {
		return nil, false
	}
	return (*s)[len(*s)-1], true
}

// IsEmpty chequea si la pila está vacía.
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Chequea si el carracter esta dentro del array de strings.
func ContainsChar(slice []string, targetChar string) bool {
	for _, char := range slice {
		if char == targetChar {
			return true // Character found
		}
	}
	return false // Character not found
}

// Símbolos para exercise2.go
var (
	OpenBrackets  = []string{"(", "{", "["}
	CloseBrackets = []string{")", "}", "]"}
	Pairs         = map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
)
