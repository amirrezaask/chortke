package main

import (
	"fmt"
	"strconv"
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
func isFunction(tok token) bool {
	return false
}
func shuntingYard(tokens []token) queue[token] {
	var output queue[token]
	var operators stack[token]
	for _, tok := range tokens {
		if tok.typ == "number" {
			output.push(tok)
		}
		if isMathOp(tok) {
			for operators.top() != nil && operators.top().value != "(" && getPrecedence(*operators.top()) >= getPrecedence(tok) {
				output.push(*operators.pop())
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
				output.push(*operators.pop())
			}
			operators.pop()
			if isFunction(*operators.top()) {
				output.push(*operators.pop())
			}
		}

	}
	for len(operators.data) != 0 {
		if operators.top().value == "(" {
			panic("unmatched paren")
		}
		output.push(*operators.pop())
	}
	fmt.Println(output)
	return output
}

func evalMath(op string, arg1, arg2 float64) float64 {

	if op == "+" {
		return arg1 + arg2
	}
	if op == "-" {
		return arg1 - arg2
	}
	if op == "*" {
		return arg1 * arg2
	}
	if op == "/" {
		return arg1 / arg2
	}
	return 0
}

func evalPostfixQueue[T any](q queue[token]) interface{} {
	var values stack[interface{}]
	for _, elem := range q.data {
		if elem.typ == "number" {
			values.push(elem.value)
			continue
		}
		if isMathOp(elem) {
			fmt.Println(values.data)
			// we should evaluate now :)
			arg2, _ := strconv.ParseFloat((*values.pop()).(string), 10)
			arg1, _ := strconv.ParseFloat((*values.pop()).(string), 10)
			fmt.Println(elem.value, arg1, arg2)

			values.push(fmt.Sprint(evalMath(elem.value, float64(arg1), float64(arg2))))
		}
	}
	fmt.Println(values.data)
	return values.data[len(values.data)-1]
}

func main() {
	code := `2+3*2-1/2`
	tokens := lex(code)
	e := evalPostfixQueue[token](shuntingYard(tokens)).(string)
	fmt.Println(e)

}
