package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type postfixTest struct {
	postfix string
	expect  int64
}

func TestParseFromPostfix(t *testing.T) {
	cases := []postfixTest{
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
		tree, err := parser.FromPostfix(test.postfix)
		assert.Nil(t, err)
		res, err := tree.Calc()
		assert.Nil(t, err)
		assert.Equal(t, test.expect, res)
	}
}
