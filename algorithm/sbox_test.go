package algorithm

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
	result := []byte("Mei ist das kalt doa rin")
	SubWord(result)
	assert.NotEqual(t, []byte("Mei ist das kalt doa rin"), result)
	SubWord(result)
	InvSubWord(result)
	InvSubWord(result)
	assert.Equal(t, []byte("Mei ist das kalt doa rin"), result)
}
