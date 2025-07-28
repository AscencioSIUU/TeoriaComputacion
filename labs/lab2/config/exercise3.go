package config

import (
	"fmt"
	"strings"
	"unicode"
)

// Precedencias de operadores
var OperatorPrecedence = map[rune]int{
	'(': 1,
	'|': 2,
	'.': 3,
	'*': 4,
	'+': 4,
	'?': 4,
	'^': 5,
}

var (
	AllOperators    = []rune{'|', '*', '+', '?', '^'}
	BinaryOperators = []rune{'|', '^'}
)

func IsAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func ContainsRune(slice []rune, r rune) bool {
	for _, x := range slice {
		if x == r {
			return true
		}
	}
	return false
}

func shouldInsertConcat(c1, c2 rune) bool {
	// Evita concatenar si alguno ya es un punto explícito
	if c1 == '.' || c2 == '.' {
		return false
	}
	if c1 == '(' || c2 == ')' {
		return false
	}
	if c2 == '*' || c2 == '+' || c2 == '?' {
		return false
	}
	if c1 == '|' || c1 == '^' {
		return false
	}
	if ContainsRune(AllOperators, c1) || ContainsRune(AllOperators, c2) {
		return false
	}
	return true
}

func FormatRegex(regex string) string {
	var b strings.Builder
	chars := []rune(regex)
	i := 0

	for i < len(chars) {
		c1 := chars[i]
		if c1 == '\\' && i+1 < len(chars) {
			b.WriteRune(c1)
			b.WriteRune(chars[i+1])
			i += 2
			if i < len(chars) && shouldInsertConcat(chars[i-1], chars[i]) {
				b.WriteRune('.')
			}
			continue
		}
		b.WriteRune(c1)
		if i+1 < len(chars) && shouldInsertConcat(c1, chars[i+1]) {
			b.WriteRune('.')
		}
		i++
	}
	return b.String()
}

func ExpandRegexExtensions(expr string) string {
	var result strings.Builder
	chars := []rune(expr)

	for i := 0; i < len(chars); i++ {
		c := chars[i]
		if c == '\\' && i+1 < len(chars) {
			result.WriteRune(c)
			result.WriteRune(chars[i+1])
			i++
			continue
		}
		if (c == '+' || c == '?') && i > 0 {
			previous := chars[i-1]
			if c == '+' {
				result.WriteRune(previous)
				result.WriteRune('.')
				result.WriteRune(previous)
				result.WriteRune('*')
			} else {
				result.WriteRune('(')
				result.WriteRune(previous)
				result.WriteRune('|')
				result.WriteRune('ε')
				result.WriteRune(')')
			}
		} else {
			result.WriteRune(c)
		}
	}
	return result.String()
}

func InfixToPostfix(rawRegex string) string {
	expr := rawRegex
	var output strings.Builder
	var stack []rune

	for _, c := range expr {
		switch {
		case IsAlphanumeric(c) || c == '[' || c == ']':
			fmt.Printf("Append operando '%c' → output = %s\n", c, output.String())
			output.WriteRune(c)
		case c == '(':
			fmt.Printf("Push '(': stack = %q\n", stack)
			stack = append(stack, c)
		case c == ')':
			fmt.Println("Encontrado ')', pop hasta '('")
			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				output.WriteRune(top)
				fmt.Printf("  Pop '%c' → output = %s\n", top, output.String())
			}
			if len(stack) > 0 {
				fmt.Printf("  Pop '(': stack = %q\n", stack)
				stack = stack[:len(stack)-1]
			}
		default:
			precC := OperatorPrecedence[c]
			fmt.Printf("Operador '%c' (precedencia %d) encontrado\n", c, precC)
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				precTop := OperatorPrecedence[top]
				if precTop >= precC {
					stack = stack[:len(stack)-1]
					output.WriteRune(top)
					fmt.Printf("  Pop '%c' (prec %d ≥ %d) → output = %s\n", top, precTop, precC, output.String())
					continue
				}
				break
			}
			stack = append(stack, c)
			fmt.Printf("Push '%c': stack = %q\n", c, stack)
		}
	}

	fmt.Println("Fin de input, vaciando pila:")
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		output.WriteRune(top)
		fmt.Printf("  Pop '%c' → output = %s\n", top, output.String())
	}
	return output.String()
}
