package main

type ASTNode struct {
	Token    Token
	Left     *ASTNode
	Right    *ASTNode
	Children []*ASTNode // For holding statements inside blocks like if
}

func parser(tokens []Token) *ASTNode {
	if len(tokens) < 2 {
		return nil
	}

	if tokens[0].Type == SEX && tokens[1].Type == VAR && tokens[2].Type == EQUAL {
		return &ASTNode{
			Token: tokens[1], // variable name
			Left:  nil,
			Right: &ASTNode{Token: tokens[3]}, // value
		}
	}

	// Handle if condition
	if tokens[0].Type == IF {
		return &ASTNode{
			Token: tokens[0], // if keyword
			Left: &ASTNode{ // condition
				Token: tokens[1], // variable name
				Right: &ASTNode{Token: tokens[3]}, // comparison value
			},
		}
	}

	// Handle when statement
	if tokens[0].Type == WHEN {
		return &ASTNode{
			Token: tokens[0], // when keyword
			Left:  &ASTNode{Token: tokens[1]}, // variable name
		}
	}

	if tokens[0].Type == FUCK {
		return &ASTNode{
			Token: tokens[0], // Fack keyword
			Left:  &ASTNode{Token: tokens[1]}, // expression to print
		}
	}

	return nil
}
