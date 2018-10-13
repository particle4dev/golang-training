package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArithmeticSequence(t *testing.T) {

	// assert equality
	// (1, 2, 5) => 1 + 3 + 5 + 7 + 9 = > 25
	assert.Equal(t, 25, ArithmeticSequence(9, 2), "they should be equal")

	// (5, 1, 3) => 5 + 6 + 7  = > 18
	assert.Equal(t, 6, ArithmeticSequence(3, 1), "they should be equal")
}
