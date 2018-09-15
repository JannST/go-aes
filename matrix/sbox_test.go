package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSBox(t *testing.T) {
	var buff byte
	for i := 0; i < 256; i++ {
		buff = sBox[i]
		assert.Equal(t, byte(i), rSBox[buff])
	}
}

func TestSBoxMatrix(t *testing.T) {
	mat := NewMatrix(4, 6)
	result := []byte("Mei ist das kalt doa rin")
	mat.setData(result)
	mat.SubColumn(0)
	assert.NotEqual(t, result, mat.getData())
	mat.SubColumn(0)
	mat.SubColumn(5)
	mat.InvSubColumn(1)
	mat.InvSubColumn(0)
	mat.InvSubColumn(0)
	mat.InvSubColumn(5)
	mat.SubColumn(1)
	assert.Equal(t, result, mat.getData())
}
