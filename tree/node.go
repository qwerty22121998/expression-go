package tree

import (
	"fmt"
	"math"
)

type Operator string

type OperatorRune rune

type NodeType int64

const (
	Add  Operator = "+"
	Sub  Operator = "-"
	Mul  Operator = "*"
	Div  Operator = "/"
	Mod  Operator = "%"
	None Operator = ""

	RAdd OperatorRune = '+'
	RSub OperatorRune = '-'
	RMul OperatorRune = '*'
	RDiv OperatorRune = '/'
	RMod OperatorRune = '%'
)

var mpOp map[Operator]bool = map[Operator]bool{
	Add: true,
	Sub: true,
	Mul: true,
	Div: true,
	Mod: true,
}

var mpOpRune map[OperatorRune]Operator = map[OperatorRune]Operator{
	RAdd: Add,
	RSub: Sub,
	RMul: Mul,
	RDiv: Div,
	RMod: Mod,
}

const (
	TypeOperator NodeType = iota
	TypeNumber
)

type Node struct {
	level int64
	Root  *Node
	L, R  *Node
	Type  NodeType
	Op    Operator
	Val   int64
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func (n *Node) maxLevel() int64 {
	var l, r int64
	if n.L != nil {
		l = n.L.maxLevel()
	}
	if n.R != nil {
		r = n.R.maxLevel()
	}
	return max(n.level, max(l, r))
}

func (n *Node) Calc() (int64, error) {
	if n.Root != nil {
		n.level = n.Root.level + 1
	}
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

func Visualize(tree *Node) {
	tree.Calc()
	level := tree.maxLevel()
	fmt.Println(level)
}
