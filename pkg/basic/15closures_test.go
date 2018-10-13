package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntSeq(t *testing.T) {

	nextInt1 := IntSeq()

	// assert equality
	assert.Equal(t, 1, nextInt1(), "they should be equal")
	assert.Equal(t, 2, nextInt1(), "they should be equal")
	assert.Equal(t, 3, nextInt1(), "they should be equal")

	nextInt2 := IntSeq()
	assert.Equal(t, 1, nextInt2(), "they should be equal")

}
