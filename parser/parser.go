package parser

import (
	"fmt"
	"github.com/qwerty22121998/expression-go/tree"
	"strconv"
	"strings"
)

type Parser struct {
}

func (p *Parser) parse(s string) (*tree.Node, error) {
	if tree.IsOperator(s) {
		return &tree.Node{
			Type: tree.TypeOperator,
			Op:   tree.Operator(s),
		}, nil
	}
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return &tree.Node{
			Type: tree.TypeNumber,
			Op:   tree.None,
			Val:  num,
		}, nil
	}
	return nil, fmt.Errorf("string value %s could not reconigzed", s)
}

func (p *Parser) FromPostfix(expression string) (*tree.Node, error) {
	syms := strings.Split(expression, " ")

	stack := new(tree.Stack)
	var n *tree.Node
	var err error
	for _, sym := range syms {
		n, err = p.parse(sym)
		if err != nil {
			return nil, err
		}
		switch n.Type {
		case tree.TypeNumber:
			stack.Push(n)
		case tree.TypeOperator:
			r := stack.Pop()
			l := stack.Pop()
			n.L = l.Val
			n.R = r.Val
			stack.Push(n)
		}
	}
	return n, err
}
