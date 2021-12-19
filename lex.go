package main

import (
	"fmt"
	"unicode"
)

const (
	tokenType_Symbol         = "symbol"
	tokenType_Bool           = "boolean"
	tokenType_Number         = "number"
	tokenType_Colon          = "colon"
	tokenType_Asterix        = "asterix"
	tokenType_Plus           = "plus"
	tokenType_Minus          = "minus"
	tokenType_Div            = "div"
	tokenType_Mod            = "modulus"
	tokenType_OpenSqBracket  = "open_square_bracket"
	tokenType_CloseSqBracket = "close_squate_bracket"
	tokenType_OpenParen      = "open_paren"
	tokenType_CloseParen     = "close_paren"
)

type token struct {
	typ   string
	value string
}

func (t token) String() string {
	return fmt.Sprintf("<%s %s>", t.typ, t.value)
}

func flush(tokens *[]token, buffType *string, buff *string) {
	if *buff != "" {
		if *buffType == "" {
			if *buff == "true" || *buff == "false" {
				*buffType = "boolean"
			} else {
				*buffType = "symbol"
			}
		}
		*tokens = append(*tokens, token{typ: *buffType, value: *buff})
	}
	*buff = ""
	*buffType = ""

}
func lex(code string) []token {
	var tokens []token
	var buffType string
	var buff string
	for _, char := range code {
		if char == '[' {
			flush(&tokens, &buffType, &buff)
			tokens = append(tokens, token{typ: tokenType_OpenSqBracket, value: "["})
		} else if char == ']' {
			flush(&tokens, &buffType, &buff)
			tokens = append(tokens, token{typ: tokenType_CloseSqBracket, value: "]"})
		} else if char == '*' {
			flush(&tokens, &buffType, &buff)
			tokens = append(tokens, token{typ: tokenType_Asterix, value: "*"})
		} else if char == '+' {
			flush(&tokens, &buffType, &buff)
			tokens = append(tokens, token{typ: tokenType_Plus, value: "+"})
		} else if char == '-' {
			flush(&tokens, &buffType, &buff)
			tokens = append(tokens, token{typ: tokenType_Minus, value: "-"})
		} else if char == '/' {
			flush(&tokens, &buffType, &buff)
			tokens = append(tokens, token{typ: tokenType_Div, value: "/"})
		} else if char == '%' {
			flush(&tokens, &buffType, &buff)
			tokens = append(tokens, token{typ: tokenType_Mod, value: "%"})
		} else if char == '(' {
			flush(&tokens, &buffType, &buff)
			tokens = append(tokens, token{typ: tokenType_OpenParen, value: "("})
		} else if char == ')' {
			flush(&tokens, &buffType, &buff)
			tokens = append(tokens, token{typ: tokenType_CloseParen, value: ")"})
		} else if unicode.IsDigit(char) {
			buffType = tokenType_Number
			buff = buff + string(char)
		} else if unicode.IsLetter(char) {
			buff = buff + string(char)
		} else if char == ':' {
			flush(&tokens, &buffType, &buff)
			tokens = append(tokens, token{typ: tokenType_Colon, value: ":"})
		} else if char == ' ' {
			flush(&tokens, &buffType, &buff)
		}
	}
	flush(&tokens, &buffType, &buff)
	return tokens
}

func isMathOp(t token) bool {
	return (t.typ == tokenType_Plus ||
		t.typ == tokenType_Minus ||
		t.typ == tokenType_Div ||
		t.typ == tokenType_Asterix ||
		t.typ == tokenType_Mod)
}
