package main

import (
	"bufio"
	"fmt" // Format the input and output of the data
	"log" //
	"os"  //
	"strings"
)

// Stack that represent the pile structure.
type Stack []interface{}

// Append the element to the stack.
func (s *Stack) Push(value interface{}) {
	fmt.Println("Pila: push: ", value)
	*s = append(*s, value)
}

// Delete and return the last item in the stack.
// If is empty, return nil and false.
func (s *Stack) Pop() (interface{}, bool) {
	if len(*s) == 0 {
		return nil, false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	fmt.Println("Pila: pop", element)
	return element, true
}

// Returns the TOP element without deleting it
func (s *Stack) Peek() (interface{}, bool) {
	if len(*s) == 0 {
		return nil, false
	}
	return (*s)[len(*s)-1], true
}

func containsChar(slice []string, targetChar string) bool {
	for _, char := range slice {
		if char == targetChar {
			return true // Character found
		}
	}
	return false // Character not found
}

// Check if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Function Main
func main() {
	var open = []string{"(", "{", "["}
	var close = []string{")", "}", "]"}
	var pairs = map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	file, err := os.Open("expressions2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var stack Stack

		line := scanner.Text()
		fmt.Println("⚪ Expresión: ", line)
		char := strings.Split(line, "")

		for i, c := range char {
			if containsChar(open, c) { // For opening characters
				stack.Push(c)
			}
			if containsChar(close, c) { // For closing characters
				if stack.IsEmpty() {
					fmt.Printf("Pila vacía al encontrar %s — 🔴 No balanceada\n", c)
					break
				}
				top, _ := stack.Peek()
				if pairs[c] == top {
					stack.Pop()
				} else {
					fmt.Printf("🔴 Pila: se esperaba %s pero se encontró %s\n", pairs[c], top)
					break
				}
			}
			if len(char)-1 == i {
				if stack.IsEmpty() {
					fmt.Println("🟢 Expresión balanceada")
				}
			}
		}
		fmt.Println("< — — — — — — — — — — — - - - - - >")
		fmt.Println("\n")
	}
	// If is an error in scann
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
