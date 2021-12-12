package main

const(
	ExprType_Reduce = "reduce"
	ExprType_Map = "map"
	ExprType_Let = "let"
	ExprType_Function = "function"
	ExprType_Data = "data"
)
const DataType (
	DataType_String = "string"
	DataType_List = "list"
	DataType_Hash = "hash"
	DataType_Number = "number"
)

type Expression struct {
	ExprType string
	Args []*Expression
}

func parseExpression(code string) *Expression {
	i := strings.Index(code, " ")
	if i == -1 {
		panic("no expression found")
	}
	expressionType := code[0:i]
	code = strings.Trim(code, " ")
}

func main() {
	code = `map [1,2,3]`

}