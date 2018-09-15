package matr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	mat := newMatrix(4, 6)
	dat := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	assert.Equal(t, mat.data, dat)

	res := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	assert.Equal(t, mat.getData(), res)

	mat = newMatrix(2, 3)
	dat = []byte{0, 0, 0, 0, 0, 0}
	assert.Equal(t, mat.data, dat)

	res = []byte{0, 0, 0, 0, 0, 0}
	assert.Equal(t, mat.getData(), res)
}
