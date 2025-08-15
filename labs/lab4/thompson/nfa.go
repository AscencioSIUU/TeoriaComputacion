// Package thompson implements Thompson's construction algorithm to build a
// non-deterministic finite automaton (NFA) from a regular expression AST.
// It supports literals, concatenation, union, and Kleene star operations.
package thompson

import (
	"fmt"
	"lab4/regex"
)

const Epsilon rune = 'ε'

// State represents a state in the NFA.
type State struct {
	ID    int
	Trans map[rune][]*State
}

// NFA represents a non-deterministic finite automaton.
type NFA struct {
	Start  *State
	Accept *State
	States []*State
}

// builder helps in constructing the NFA.
type builder struct{ next int }

// newState creates a new state with a unique ID.
func (b *builder) newState() *State {
	s := &State{
		ID:    b.next,
		Trans: make(map[rune][]*State),
	}
	b.next++
	return s
}

// addEdge adds a transition from 'from' to 'to' on symbol 'sym'.
func (b *builder) addEdge(from *State, sym rune, to *State) {
	from.Trans[sym] = append(from.Trans[sym], to)
}

type frag struct {
	start, accept *State
}

// Build constructs an NFA from the given regex AST using Thompson's construction.
func Build(ast *regex.Node) (*NFA, error) {
	if ast == nil {
		return nil, fmt.Errorf("nil AST")
	}
	b := &builder{}
	f := b.buildRec(ast)

	// collect all reachable states
	seen := map[int]*State{}
	var dfs func(*State)
	dfs = func(s *State) {
		if s == nil {
			return
		}
		if _, ok := seen[s.ID]; ok {
			return
		}
		seen[s.ID] = s
		for _, outs := range s.Trans {
			for _, t := range outs {
				dfs(t)
			}
		}
	}
	dfs(f.start)

	states := make([]*State, 0, len(seen))
	for _, s := range seen {
		states = append(states, s)
	}

	return &NFA{
		Start:  f.start,
		Accept: f.accept,
		States: states,
	}, nil
}

// buildRec is a recursive helper to build NFA fragments from AST nodes.
// It returns the start and accept states of the fragment.
func (b *builder) buildRec(n *regex.Node) frag {
	switch n.Kind {
	case regex.Literal:
		// Single literal: create two states and connect them
		s := b.newState()
		t := b.newState()
		b.addEdge(s, n.Val, t)
		return frag{start: s, accept: t}

	case regex.Concat:
		// Concatenation: connect two fragments
		f1 := b.buildRec(n.Left)
		f2 := b.buildRec(n.Right)
		b.addEdge(f1.accept, Epsilon, f2.start)
		return frag{start: f1.start, accept: f2.accept}

	case regex.Union:
		// Union: create new start and accept states, connect with Epsilon
		// to the two fragments and from their accept states to the new accept state
		s := b.newState()
		t := b.newState()
		f1 := b.buildRec(n.Left)
		f2 := b.buildRec(n.Right)
		b.addEdge(s, Epsilon, f1.start)
		b.addEdge(s, Epsilon, f2.start)
		b.addEdge(f1.accept, Epsilon, t)
		b.addEdge(f2.accept, Epsilon, t)
		return frag{start: s, accept: t}

	case regex.Star:
		// Kleene star: create new start and accept states, connect with Epsilon
		// to the fragment and from its accept state back to its start and to the new accept state
		s := b.newState()
		t := b.newState()
		f := b.buildRec(n.Left)
		b.addEdge(s, Epsilon, f.start)
		b.addEdge(s, Epsilon, t)
		b.addEdge(f.accept, Epsilon, f.start)
		b.addEdge(f.accept, Epsilon, t)
		return frag{start: s, accept: t}

	default:
		panic("unknown node kind")
	}
}
