package main
import (
	"fmt"
	"unicode"

)
const (
	tokenType_Symbol = "symbol"
	tokenType_String = "string"
	tokenType_Number = "number"
	tokenType_Arrow = "arrow"
	tokenType_Asterix = "asterix"
	tokenType_Plus = "plus"
	tokenType_Minus = "minus"
	tokenType_Div = "div"
	tokenType_Mod = "modulus"
	tokenType_OpenSqBracket = "["
	tokenType_CloseSqBracket = "]"
)

type token struct {
	typ string
	value string
}

func lex(code string) []token {
	var tokens []token
	var buffType string
	var buff string
	for _, char := range code  {
		if char == '[' {
			tokens = append(tokens, token{typ: tokenType_OpenSqBracket, value: "["})
		} else if char == ']' {
			tokens = append(tokens, token{typ: buffType, value: buff})
			tokens = append(tokens, token{typ: tokenType_CloseSqBracket, value: "]"})
			buff = ""
			buffType = ""
		} else if char == '*' {
			tokens = append(tokens, token{typ: tokenType_Asterix, value: "*"})
		} else if char == '+' {
			tokens = append(tokens, token{typ: tokenType_Plus, value: "+"})
		} else if char == '-' {
			tokens = append(tokens, token{typ: tokenType_Minus, value: "-"})
		} else if char == '/' {
			tokens = append(tokens, token{typ: tokenType_Div, value: "/"})
		} else if char == '%' {
			tokens = append(tokens, token{typ: tokenType_Mod, value: "%"})
		} else if unicode.IsDigit(char) {
			buffType = tokenType_Number
			buff = buff + string(char)
		} else if unicode.IsLetter(char) {
			buff = buff + string(char)
		} else if char == '"' {
			buffType = tokenType_String
		} else if char == ' ' {
			if buff != "" {
				tokens = append(tokens, token{typ: buffType, value: buff})
			}

			buff = ""
			buffType = ""
		}
	}
	return tokens
}
func main() {
	code := `map [1,2,3] { x -> x * 2 }`
	fmt.Printf("%+v\n", lex(code))
}