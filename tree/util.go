package tree

func IsOperator(s string) bool {
	return mpOp[Operator(s)]
}
