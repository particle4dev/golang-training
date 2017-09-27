package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlus(t *testing.T) {

	// assert equality
	assert.Equal(t, 3, Plus(1, 2), "they should be equal")

}

func TestPlusPlus(t *testing.T) {

	assert.Equal(t, 6, PlusPlus(1, 2, 3), "they should be equal")

}

func TestVals(t *testing.T) {

	a, b := Vals()
	assert.Equal(t, 3, a, "they should be equal")
	assert.Equal(t, 7, b, "they should be equal")
	_, c := Vals()
	assert.Equal(t, 7, c, "they should be equal")

}

func TestSum(t *testing.T) {

	var a = Sum(1, 2, 3)
	assert.Equal(t, 6, a, "they should be equal")

}
