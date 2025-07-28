package main

import (
	"bufio"
	"fmt"
	"lab2/config"
	"log"
	"os"
	"strings"
)

// Function Main
func main() {
	file, err := os.Open("expressions2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var stack config.Stack

		line := scanner.Text()
		fmt.Println("⚪ Expresión: ", line)
		char := strings.Split(line, "")

		for i, c := range char {
			if config.ContainsChar(config.OpenBrackets, c) { // For opening characters
				stack.Push(c)
			}
			if config.ContainsChar(config.CloseBrackets, c) { // For closing characters
				if stack.IsEmpty() {
					fmt.Printf("Pila vacía al encontrar %s — 🔴 No balanceada\n", c)
					break
				}
				top, _ := stack.Peek()
				if config.Pairs[c] == top {
					stack.Pop()
				} else {
					fmt.Printf("🔴 Pila: se esperaba %s pero se encontró %s\n", config.Pairs[c], top)
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
