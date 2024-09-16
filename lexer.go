package main

import (
	"strings"
	"fmt"
)

type TokenType string

const (
	IF      TokenType = "IF"
	FOR     TokenType = "FOR"
	WHILE   TokenType = "WHILE"
	SEX     TokenType = "SEX"
	WHEN    TokenType = "WHEN"
	VAR     TokenType = "VAR"
	EQUAL   TokenType = "EQUAL"
	INT     TokenType = "INT"
	FUCK    TokenType = "FUCK"   // Print statement
	STRING  TokenType = "STRING"
	UNKNOWN TokenType = "UNKNOWN"
)

type Token struct {
	Type  TokenType
	Value string
}

func lexer(input string) []Token {
	var tokens []Token
	words := strings.Fields(input)

	for _, word := range words {
		switch word {
		case "if":
			tokens = append(tokens, Token{Type: IF, Value: word})
		case "for":
			tokens = append(tokens, Token{Type: FOR, Value: word})
		case "while":
			tokens = append(tokens, Token{Type: WHILE, Value: word})
		case "sex":
			tokens = append(tokens, Token{Type: SEX, Value: word})
		case "when":
			tokens = append(tokens, Token{Type: WHEN, Value: word})
		case "Fuck":  // Add support for Fack
			tokens = append(tokens, Token{Type: FUCK, Value: word})
		case "=":
			tokens = append(tokens, Token{Type: EQUAL, Value: word})
		default:
			if isNumeric(word) {
				tokens = append(tokens, Token{Type: INT, Value: word})
			} else {
				tokens = append(tokens, Token{Type: VAR, Value: word})
			}
		}
	}
	return tokens
}

func isNumeric(s string) bool {
	_, err := fmt.Sscanf(s, "%d", new(int))
	return err == nil
}
