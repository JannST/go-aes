package tables

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
	test := []byte("Mei ist das kalt doa rin")
	SubWord(test)
	assert.NotEqual(t, []byte("Mei ist das kalt doa rin"), test)
	SubWord(test)
	InvSubWord(test)
	InvSubWord(test)
	assert.Equal(t, []byte("Mei ist das kalt doa rin"), test)
}
