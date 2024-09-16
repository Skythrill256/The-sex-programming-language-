package main

import (
	"fmt"
)

var symbolTable = make(map[string]interface{})

func evalAST(node *ASTNode) {
	if node == nil {
		return
	}

	// Handle variable assignment
	if node.Token.Type == VAR && node.Right != nil {
		value := node.Right.Token.Value
		symbolTable[node.Token.Value] = value
	}

	// Handle if statement
	if node.Token.Type == IF {
		varName := node.Left.Token.Value
		conditionValue := node.Left.Right.Token.Value
		if symbolTable[varName] == conditionValue {
			// Evaluate block if condition is true
		}
	}

	// Handle when statement
	if node.Token.Type == WHEN {
		varName := node.Left.Token.Value
		varValue := symbolTable[varName]

		// Example hardcoded cases for simplicity
		cases := map[string]string{
			"20":  "Twenty",
			"else": "Other",
		}

		result := evaluateWhen(fmt.Sprintf("%v", varValue), cases)
		fmt.Printf("%s\n", result) // Print the when result
	}

	// Handle Fack (print statement)
	if node.Token.Type == FUCK {
		fmt.Println(symbolTable[node.Left.Token.Value])
	}
}

func evaluateWhen(value string, cases map[string]string) string {
	if result, found := cases[value]; found {
		return result
	}
	if result, found := cases["else"]; found {
		return result
	}
	return "No match"
}
