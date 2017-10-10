package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRect(t *testing.T) {

	r := rect{width: 10, height: 5}
	assert.Equal(t, r.area(), 50, "they should be equal")
	assert.Equal(t, r.perim(), 30, "they should be equal")

	// Go automatically handles conversion between values and pointers
	// for method calls. You may want to use a pointer receiver type to
	// avoid copying on method calls or to allow the method to mutate
	// the receiving struct.
	rp := &r
	assert.Equal(t, rp.area(), 50, "they should be equal")
	assert.Equal(t, rp.perim(), 30, "they should be equal")

}
