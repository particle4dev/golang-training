package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberValAndNumberPoi(t *testing.T) {

	i := 5
	assert.Equal(t, 5, i, "they should be equal")

	NumberVal(i)
	assert.Equal(t, 5, i, "they should be equal")

	NumberPoi(&i)
	assert.Equal(t, 10, i, "they should be equal")
}
