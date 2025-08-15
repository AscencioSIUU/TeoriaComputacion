// Package nfa provides functionality to simulate a non-deterministic finite automaton (NFA).
package nfa

import (
	"lab4/thompson"
	"unicode/utf8"
)

// stateSet represents a set of NFA states.
type stateSet map[*thompson.State]struct{}

// add adds a state to the state set.
func add(m stateSet, s *thompson.State) { m[s] = struct{}{} }

// epsilonClosure computes the epsilon closure of a set of states.
func epsilonClosure(start stateSet) stateSet {
	stack := make([]*thompson.State, 0, len(start))
	seen := make(stateSet)

	// initialize stack and seen with start states
	for s := range start {
		stack = append(stack, s)
		seen[s] = struct{}{}
	}

	for len(stack) > 0 {
		// pop
		s := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// explore epsilon transitions
		for _, nxt := range s.Trans[thompson.Epsilon] {
			if _, ok := seen[nxt]; !ok {
				seen[nxt] = struct{}{}
				stack = append(stack, nxt)
			}
		}
	}
	return seen
}

// move computes the set of states reachable from 'from' on input 'sym'.
func move(from stateSet, sym rune) stateSet {
	out := make(stateSet)
	for s := range from {
		for _, nxt := range s.Trans[sym] {
			add(out, nxt)
		}
	}
	return out
}

// Simulate returns true if input is accepted by the NFA.
func Simulate(nfa *thompson.NFA, input string) bool {
	current := make(stateSet)
	add(current, nfa.Start)
	current = epsilonClosure(current)

	// iterate runes (supports UTF-8, including '0','1','a','b', etc.)
	for len(input) > 0 {
		r, size := utf8.DecodeRuneInString(input)
		input = input[size:]

		next := move(current, r)
		current = epsilonClosure(next)
	}

	_, ok := current[nfa.Accept]
	return ok
}
