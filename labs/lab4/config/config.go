// Package config provides utilities for regex formatting and conversion to postfix notation.
package config

import (
	"fmt"
	"strings"
	"unicode"
)

// OperatorPrecedence defines precedence for regex operators.
var OperatorPrecedence = map[rune]int{
	'(': 1,
	'|': 2,
	'.': 3,
	'*': 4,
}

// Keep only if you use them elsewhere; otherwise you can remove these slices.
var (
	AllOperators    = []rune{'|', '.', '*'}
	BinaryOperators = []rune{'|', '.'}
)

// IsAlphanumeric returns true if r is a letter, digit, or epsilon.
func IsAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == 'ε'
}

// ContainsRune checks if a slice contains a specific rune.
func ContainsRune(slice []rune, r rune) bool {
	for _, x := range slice {
		if x == r {
			return true
		}
	}
	return false
}

// shouldInsertConcat returns true if a '.' should be inserted between c1 and c2.
func shouldInsertConcat(c1, c2 rune) bool {
	// concat when: (symbol or '*' or ')') followed by (symbol or '(')
	if (IsAlphanumeric(c1) || c1 == '*' || c1 == ')') &&
		(IsAlphanumeric(c2) || c2 == '(') {
		return true
	}
	return false
}

// FormatRegex inserts explicit concatenation operators '.' where needed.
func FormatRegex(regex string) string {
	var b strings.Builder
	chars := []rune(regex)
	i := 0

	for i < len(chars) {
		c1 := chars[i]
		// preserve escapes
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

// normalizeEpsilon replaces all variants of epsilon with 'ε'.
func normalizeEpsilon(s string) string {
	return strings.ReplaceAll(s, "𝜀", "ε")
}

// ExpandRegexExtensions expands '+' and '?' in the regex to their basic equivalents.
func ExpandRegexExtensions(expr string) string {
	expr = normalizeEpsilon(expr)
	in := []rune(expr)
	out := make([]rune, 0, len(in))

	for i := 0; i < len(in); i++ {
		c := in[i]

		// preserve escapes
		if c == '\\' && i+1 < len(in) {
			out = append(out, c, in[i+1])
			i++
			continue
		}
		// handle '+' and '?'
		if (c == '+') || (c == '?') {
			start, end := lastOperandBounds(out)
			X := string(out[start:end])

			// replace last operand with its expansion
			// '+' -> X.X*
			// '?' -> (X|ε)
			tmp := make([]rune, 0, len(out))
			tmp = append(tmp, out[:start]...)

			if c == '+' {
				tmp = append(tmp, []rune(X)...)
				tmp = append(tmp, '.')
				tmp = append(tmp, []rune(X)...)
				tmp = append(tmp, '*')
			} else {
				tmp = append(tmp, '(')
				tmp = append(tmp, []rune(X)...)
				tmp = append(tmp, '|', 'ε', ')')
			}
			out = tmp
			continue
		}

		out = append(out, c)
	}

	return string(out)
}

// lastOperandBounds finds the start and end indices of the last operand in out.
// An operand can be a single symbol, an escaped symbol, or a parenthesized group
func lastOperandBounds(out []rune) (int, int) {
	if len(out) == 0 {
		return 0, 0
	}
	j := len(out) - 1

	// case 1: parenthesized group
	if out[j] == ')' {
		depth := 0
		for j >= 0 {
			if out[j] == ')' {
				depth++
			} else if out[j] == '(' {
				depth--
				if depth == 0 {
					return j, len(out)
				}
			}
			j--
		}
		return 0, len(out) // fallback if unbalanced
	}

	// case 2: escaped rune
	if j > 0 && out[j-1] == '\\' {
		return j - 1, j + 1
	}

	// case 3: single rune
	return j, j + 1
}

// InfixToPostfix converts an infix regex expression to postfix notation using the Shunting Yard algorithm.
func InfixToPostfix(rawRegex string) string {
	expr := rawRegex
	var output strings.Builder
	var stack []rune

	for _, c := range expr {
		switch {
		case IsAlphanumeric(c):
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
