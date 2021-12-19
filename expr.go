package main

type expr interface {
	Eval(values map[string]interface{}) interface{}
}

type number struct {
	n float64
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

type arithmeticExpr struct {
	op   string
	args []expr
}

func (e *arithmeticExpr) Eval(m map[string]interface{}) interface{} {
	lhs := e.args[0].Eval(m)
	rhs := e.args[1].Eval(m)
	if e.op == "+" {
		return lhs.(float64) + rhs.(float64)
	}
	return nil
}
