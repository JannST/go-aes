package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotWord(t *testing.T) {
	test := []byte("12345")
	RotWord(test, -1)
	assert.Equal(t, []byte("23451"), test)

	test = []byte("12345")
	RotWord(test, 5)
	assert.Equal(t, []byte("12345"), test)

	test = []byte("12345")
	RotWord(test, 2)
	assert.Equal(t, []byte("45123"), test)

	test = []byte("12345")
	RotWord(test, 4)
	assert.Equal(t, []byte("23451"), test)
}
