package main

import (
	"bufio"
	"fmt"
	"lab2/config"
	"log"
	"os"
)

func main() {
	file, err := os.Open("expressions3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("⚪ Expresión Infix: ", line)
		expanded := config.ExpandRegexExtensions(line)
		formatted := config.FormatRegex(expanded)
		postfix := config.InfixToPostfix(formatted)
		fmt.Println("🟢 Expresión: ", postfix)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
