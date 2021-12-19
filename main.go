package main

import (
	"fmt"
	"strconv"
)

func parse(tokens []token) expr {
    exprStack := &stack[expr]{}
	var currentExpr expr

	for _, tok := range tokens {
		if currentExpr == nil {
			if tok.typ == tokenType_Number {
                fmt.Println("no currentExpr and token is number")
				x, err := strconv.Atoi(tok.value)
				if err != nil {
					panic(err)
				}
				exprStack.push(&number{
					n: float64(x),
				})
			}
			if isMathOp(tok) {
                fmt.Println("no currentExpr and token is mathOP")
				var args []expr
				args = append(args, exprStack.pop())
				currentExpr = &arithmeticExpr {
					op: tok.value,
					args: args,
				}
			}
		} else {
			if tok.typ == tokenType_Number {
				x, err := strconv.Atoi(tok.value)
				if err != nil {
					panic(err)
				}
				e := currentExpr.(*arithmeticExpr)
				e.args = append(e.args, &number{n: float64(x)})
                exprStack.push(e)
                currentExpr = nil
			}
		}
	}
	return exprStack.pop()
}

func main() {
	code := `123+2+3+8+9+1`
	tokens := lex(code)
	fmt.Printf("%+v\n", tokens)
    parsed := parse(tokens).(*arithmeticExpr)
    
    result := parsed.Eval(nil)
    fmt.Println(result)
}
