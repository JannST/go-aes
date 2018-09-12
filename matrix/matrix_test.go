package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	mat := NewMatrix(4, 6)
	dat := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	assert.Equal(t, mat.data, dat)

	res := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	assert.Equal(t, mat.getData(), res)

	mat = NewMatrix(2, 3)
	dat = []byte{0, 0, 0, 0, 0, 0}
	assert.Equal(t, mat.data, dat)

	res = []byte{0, 0, 0, 0, 0, 0}
	assert.Equal(t, mat.getData(), res)
}


