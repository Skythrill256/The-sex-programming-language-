package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.sex")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read and process each line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Tokenize
		tokens := lexer(line)

		// Parse into AST
		ast := parser(tokens)

		// Interpret (execute)
		evalAST(ast)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// Example of when operator
	cases := map[string]string{
		"1":   "One",
		"2":   "Two",
		"else": "Unknown",
	}
	value := "1"
	fmt.Println("When result:", evaluateWhen(value, cases))

	value = "3"
	fmt.Println("When result:", evaluateWhen(value, cases))
}
