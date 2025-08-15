// Package graphviz provides functions to generate Graphviz DOT files and PNG images
// from a Thompson NFA.
package graphviz

import (
	"fmt"
	"lab4/thompson"
	"os"
	"os/exec"
	"sort"
)

// WriteDOT writes the NFA to a DOT file at the specified path.
func WriteDOT(nfa *thompson.NFA, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// header and graph settings
	fmt.Fprintln(f, "digraph NFA {")
	fmt.Fprintln(f, "  rankdir=LR;")
	fmt.Fprintln(f, "  node [shape=circle];")

	// invisible entry arrow to start state
	fmt.Fprintf(f, "  s [shape=point];\n")
	fmt.Fprintf(f, "  s -> q%d;\n", nfa.Start.ID)

	// accept as doublecircle node
	fmt.Fprintf(f, "  q%d [shape=doublecircle];\n", nfa.Accept.ID)

	// nodes (sorted by ID for consistency)
	ids := make([]int, 0, len(nfa.States))
	idToState := make(map[int]*thompson.State)
	for _, s := range nfa.States {
		ids = append(ids, s.ID)
		idToState[s.ID] = s
	}
	sort.Ints(ids)
	for _, id := range ids {
		if id == nfa.Accept.ID {
			continue // already declared with doublecircle
		}
		fmt.Fprintf(f, "  q%d;\n", id)
	}

	// edges (sorted by from ID, label, to ID for consistency)
	for _, id := range ids {
		s := idToState[id]
		for label, outs := range s.Trans {
			lab := string(label)
			if label == thompson.Epsilon {
				lab = "ε"
			}
			for _, t := range outs {
				fmt.Fprintf(f, "  q%d -> q%d [label=\"%s\"];\n", s.ID, t.ID, lab)
			}
		}
	}

	fmt.Fprintln(f, "}")
	return nil
}

// GeneratePNGFromDot generates a PNG image from a DOT file using the 'dot' command.
func GeneratePNGFromDot(dotPath, pngPath string) error {
	cmd := exec.Command("dot", "-Tpng", dotPath, "-o", pngPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
