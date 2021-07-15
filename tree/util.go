package tree

func IsOperator(s string) bool {
	return mpOp[Operator(s)]
}

func IsOperatorRune(s rune) (Operator, bool) {
	op := mpOpRune[OperatorRune(s)]
	if op == "" {
		return None, false
	}
	return op, IsOperator(string(op))
}
