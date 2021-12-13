package main

import (
	"fmt"
)

type expr interface {
	Eval(values map[string]interface{}) interface{}
}
type number struct {
	n float64
}

type boolean struct {
	val *bool
}

func (b *boolean) Eval(values map[string]interface{}) interface{} {
	if b.val != nil {
		return *b.val
	}
	return false
}

func (n *number) Eval(values map[string]interface{}) interface{} {
	return n.n
}

type symbol struct {
	name string
}

func (s *symbol) Eval(values map[string]interface{}) interface{} {
	return values[s.name]
}

type ifelse struct {
	then  expr
	_else expr
}

type or struct {
	value1 expr
	value2 expr
}

type and struct {
	value1 expr
	value2 expr
}

type not struct {
	value expr
}

func (n *not) Eval(values map[string]interface{}) interface{} {
	return !(n.value.Eval(values).(bool))
}

type mapF struct {
	collection expr
	f          expr
}

type reduceF struct {
	collection expr
	initialAcc expr
	f          expr
}

type filterF struct {
	collection expr
	f          expr
}

type function struct {
	args []string
	body expr
}

func makeExpr(tokens []token) expr {
}

func main() {
	// code := `map [1 2 3] x: x * 2`
	code := `not true`
	fmt.Printf("%+v\n", lex(code))
}
