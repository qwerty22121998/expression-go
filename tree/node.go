package tree

import (
	"fmt"
	"math"
)

type Operator string

type NodeType int64

const (
	Add  Operator = "+"
	Sub  Operator = "-"
	Mul  Operator = "*"
	Div  Operator = "/"
	Mod  Operator = "%"
	None Operator = ""
)

var mpOp map[Operator]bool = map[Operator]bool{
	Add: true,
	Sub: true,
	Mul: true,
	Div: true,
	Mod: true,
}

const (
	TypeOperator NodeType = iota
	TypeNumber
)

type Node struct {
	L, R *Node
	Type NodeType
	Op   Operator
	Val  int64
}

func (n *Node) Calc() (int64, error) {
	var err error
	var l, r int64

	if n.L != nil {
		l, err = n.L.Calc()
		if err != nil {
			return math.MinInt64, err
		}
	}

	if n.R != nil {
		r, err = n.R.Calc()
		if err != nil {
			return math.MinInt64, err
		}
	}

	switch n.Type {
	case TypeOperator:
		{
			switch n.Op {
			case Add:
				return l + r, nil
			case Sub:
				return l - r, nil
			case Mul:
				return l * r, nil
			case Div:
				return l / r, nil
			case Mod:
				return l % r, nil
			default:
				return math.MinInt64, fmt.Errorf("operator %v not supported", n.Op)
			}
		}
	case TypeNumber:
		return n.Val, nil
	default:
		return math.MinInt64, fmt.Errorf("node type %v not supported", n.Type)
	}
}
