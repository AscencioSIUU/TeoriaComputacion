// Package regex implements a simple regular expression parser that builds an
// abstract syntax tree (AST) from a postfix expression.
// It supports literals, concatenation, union, and Kleene star operations.
package regex

import (
	"fmt"
	"lab4/config"
)

type Kind int

const (
	Literal Kind = iota
	Concat
	Union
	Star
)

// Node represents a node in the regex AST.
type Node struct {
	Kind        Kind
	Val         rune
	Left, Right *Node
}

// BuildAST constructs an AST from a postfix regex expression.
func BuildAST(postfix string) (*Node, error) {
	var stack []*Node
	// Helper to pop one or two nodes from the stack.
	pop1 := func() (*Node, error) {
		if len(stack) < 1 {
			return nil, fmt.Errorf("unary operator requires 1 operand")
		}
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return n, nil
	}
	// pop2 pops two nodes from the stack, returning left and right.
	pop2 := func() (*Node, *Node, error) {
		if len(stack) < 2 {
			return nil, nil, fmt.Errorf("binary operator requires 2 operands")
		}
		r := stack[len(stack)-1]
		l := stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		return l, r, nil
	}
	// Process each rune in the postfix expression.
	for _, c := range postfix {
		switch {
		case config.IsAlphanumeric(c):
			stack = append(stack, &Node{Kind: Literal, Val: c})
		case c == '*':
			x, err := pop1()
			if err != nil {
				return nil, err
			}
			stack = append(stack, &Node{Kind: Star, Left: x})
		case c == '.':
			l, r, err := pop2()
			if err != nil {
				return nil, err
			}
			stack = append(stack, &Node{Kind: Concat, Left: l, Right: r})
		case c == '|':
			l, r, err := pop2()
			if err != nil {
				return nil, err
			}
			stack = append(stack, &Node{Kind: Union, Left: l, Right: r})
		case c == ' ', c == '\t', c == '\n', c == '\r':
			continue
		default:
			return nil, fmt.Errorf("unexpected token %q", string(c))
		}
	}

	if len(stack) != 1 {
		return nil, fmt.Errorf("invalid postfix, final stack size = %d", len(stack))
	}
	return stack[0], nil
}
