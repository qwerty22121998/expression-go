package parser

import (
	"fmt"
	"github.com/qwerty22121998/expression-go/tree"
	"strconv"
	"strings"
	"unicode"
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
	var now *tree.Node
	var err error
	for _, sym := range syms {
		now, err = p.parse(sym)
		if err != nil {
			return nil, err
		}
		switch now.Type {
		case tree.TypeNumber:
			stack.Push(now)
		case tree.TypeOperator:
			r := stack.Pop()
			l := stack.Pop()
			l.Val.Root = now
			r.Val.Root = now
			now.L = l.Val
			now.R = r.Val
			stack.Push(now)
		}
	}
	return now, err
}

func (p *Parser) FromExpression(expression string) (*tree.Node, error) {
	eRunes := []rune(expression)

	root := new(tree.Node)
	root.L = new(tree.Node)
	root.L.Root = root
	root = root.L

	var carry strings.Builder
	for _, c := range eRunes {
		if unicode.IsDigit(c) {
			carry.WriteRune(c)
			continue
		}
		if carry.Len() > 0 {
			fmt.Println("number", carry.String())
			num, err := strconv.ParseInt(carry.String(), 10, 64)
			if err != nil {
				return nil, err
			}
			root.Type = tree.TypeNumber
			root.Val = num
			root = root.Root
			carry.Reset()
		}

		if c == ' ' {
			continue
		}
		if c == '(' {
			fmt.Println("begin (")
			now := &tree.Node{
				Root: root,
			}
			root.L = now
			root = now
			continue
		}
		if c == ')' {
			fmt.Println("end )")
			root = root.Root
		}
		if op, ok := tree.IsOperatorRune(c); ok {
			fmt.Println("op", op)
			root.Type = tree.TypeOperator
			root.Op = op
			now := &tree.Node{
				Root: root,
			}
			root.R = now
			root = now
			continue
		}
	}
	if carry.Len() > 0 {
		num, err := strconv.ParseInt(carry.String(), 10, 64)
		if err != nil {
			return nil, err
		}
		root.Type = tree.TypeNumber
		root.Val = num
		root = root.Root
	}
	for root.Root != nil {
		root = root.Root
	}
	return root, nil
}
