package config

import (
	"fmt"
	"os"
	"os/exec"
)

type Node struct {
	ID    int
	Label string
	Left  *Node
	Right *Node
}

var nodeIdCounter int = 0

func PostfixToTree(postfix string) *Node {
	var stack []*Node

	for _, i := range postfix {
		if IsAlphanumeric(i) { // Alfanuméricos
			node := &Node{
				ID:    nodeIdCounter,
				Label: string(i),
			}
			nodeIdCounter++
			stack = append(stack, node)
		} else if i == '|' || i == '.' { // operadores binarios
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			node := &Node{
				ID:    nodeIdCounter,
				Label: string(i),
				Left:  left,
				Right: right,
			}
			nodeIdCounter++
			stack = append(stack, node)
		} else if i == '*' { // Operador unitario
			child := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			node := &Node{
				ID:    nodeIdCounter,
				Label: string(i),
				Left:  child,
			}
			nodeIdCounter++
			stack = append(stack, node)
		}
	}
	if len(stack) != 1 {
		fmt.Println("Error: expresión inválida, pila final no tiene un único árbol")
		return nil
	}
	return stack[0]
}

func GenerateDotFile(root *Node, fileName int) {
	file, err := os.Create("linea" + string(fileName) + ".dot")
	if err != nil {
		fmt.Println("Error creando archivo .dot:", err)
		return
	}
	defer file.Close()
	// Encabezado
	fmt.Fprintln(file, "digraph SyntaxTree {")
	fmt.Fprintln(file, "    node [shape=circle];")

	var writeNode func(n *Node)
	writeNode = func(n *Node) {
		if n == nil {
			return
		}

		fmt.Fprintf(file, "    %d [label=\"%s\"];\n", n.ID, n.Label)
		// Escribr nodo izquierdo
		if n.Left != nil {
			fmt.Fprintf(file, "    %d -> %d;\n", n.ID, n.Left.ID)
			writeNode(n.Left)
		}
		if n.Right != nil {
			fmt.Fprintf(file, "    %d -> %d;\n", n.ID, n.Right.ID)
			writeNode(n.Right)
		}
	}
	writeNode(root)
	fmt.Fprintln(file, "}")

	fmt.Println("Archivo Dot creado:", "linea"+string(fileName)+".dot")
}

func GeneratePNGFromDot(fileName int) {
	cmd := exec.Command("dot", "-Tpng", fileName+".dot", "-o", "linea"+string(fileName)+".png")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error generando PNG:", err)
		return
	}
}
