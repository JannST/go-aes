package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModulo(t *testing.T) {
	assert.Equal(t, 3, Modulo(-1, 4))
	assert.Equal(t, 5, Modulo(-2, 7))
	assert.Equal(t, 0, Modulo(3, 3))
	assert.Equal(t, 2, Modulo(6, 4))
}
