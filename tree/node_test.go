package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNodeCalc1Plus2(t *testing.T) {
	n := &Node{
		L: &Node{
			Type: TypeNumber,
			Val:  1,
		},
		R: &Node{
			Type: TypeNumber,
			Val:  2,
		},
		Type: TypeOperator,
		Op:   Add,
	}
	res, err := n.Calc()
	assert.Nil(t, err)
	assert.Equal(t, int64(3), res)
}
