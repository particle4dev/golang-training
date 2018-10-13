package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStructs(t *testing.T) {

	s := person{name: "Sean", age: 50}
	assert.Equal(t, s.name, "Sean", "they should be equal")
	assert.Equal(t, s.age, 50, "they should be equal")

	// Structs are mutable.
	sp := &s
	assert.Equal(t, sp.name, "Sean", "they should be equal")
	sp.age = 51
	assert.Equal(t, sp.age, 51, "they should be equal")

}
