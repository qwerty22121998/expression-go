package parser

import (
	"github.com/qwerty22121998/expression-go/tree"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	express string
	expect  int64
}

func TestParser_FromPostfix(t *testing.T) {
	cases := []testCase{
		{
			"1 2 +",
			3,
		},
		{
			"2 3 4 * +",
			14,
		},
		{
			"1 2 + 7 *",
			21,
		},
	}

	parser := new(Parser)
	for _, test := range cases {
		expression, err := parser.FromPostfix(test.express)
		assert.Nil(t, err)
		res, err := expression.Calc()
		assert.Nil(t, err)
		assert.Equal(t, test.expect, res)
	}
}

func TestVisualise(t *testing.T) {
	express, err := new(Parser).FromExpression("1+2+3+4+(3*4)+11-3")
	assert.Nil(t, err)
	tree.Visualize(express)

}

func TestParser_FromExpression(t *testing.T) {
	cases := []testCase{
		//{"1+2", 3},
		//{"(2*3)+5", 11},
		{"1+2+3+4+(3*4)+11-3", 30},
	}

	parse := new(Parser)
	for _, test := range cases {
		expression, err := parse.FromExpression(test.express)
		assert.Nil(t, err)
		res, err := expression.Calc()
		assert.Nil(t, err)
		assert.Equal(t, test.expect, res)
	}
}
