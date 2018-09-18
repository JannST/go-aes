package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	mat := NewMatrix()
	dat := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	assert.Equal(t, dat, mat.data)
}
