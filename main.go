package main

import (
	"fmt"
)

func getPrecedence(tok token) int {
	if tok.value == "+" {
		return 0
	}
	if tok.value == "-" {
		return 0
	}
	if tok.value == "/" {
		return 1
	}
	if tok.value == "*" {
		return 1
	}
	if tok.value == "!" {
		return 2
	}
	if tok.value == "^" {
		return 2
	}
	return 0
}

func shuntingYard(tokens []token) expr {
	var output queue[token]
	var operators stack[token]
	for _, tok := range tokens {
		if tok.typ == "number" {
			output.push(tok)
		}
		if isMathOp(tok) {
			for operators.top().value == "(" && getPrecedence(operators.top()) >= getPrecedence(tok) {
				output.push(operators.pop())
			}
			operators.push(tok)
		}
		if tok.value == "(" {
			operators.push(tok)
		}
		if tok.value == ")" {
			for operators.top().value != "(" {
				if len(operators.data) == 0 {
					panic("unmatched paren")
				}
				output.push(operators.pop())
			}
			operators.pop()
		}

	}
	for len(operators.data) != 0 {
		if operators.top().value == "(" {
			panic("unmatched paren")
		}
		output.push(operators.pop())
	}
	fmt.Println(output)
	return nil
}

func main() {
	code := `2+3*2`
	tokens := lex(code)
	fmt.Printf("%+v\n", tokens)
	shuntingYard(tokens)

}
