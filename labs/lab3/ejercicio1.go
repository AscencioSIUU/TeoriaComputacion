package main

import (
	"bufio"
	"fmt"
	"lab3/config"
	"log"
	"os"
)

func main() {
	file, err := os.Open("expressions1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		line := scanner.Text()
		expanded := config.ExpandRegexExtensions(line)
		formatted := config.FormatRegex(expanded)
		postfix := config.InfixToPostfix(formatted)
		fmt.Println("Postfix: ", postfix)
		root := config.PostfixToTree(postfix)
		config.GenerateDotFile(root, i)
		config.GenerateDotFile(i)

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
