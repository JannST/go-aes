package math

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModulo(t *testing.T) {
	assert.Equal(t, 3, Modulo(-1, 4))
	assert.Equal(t, 5, Modulo(-2, 7))
	assert.Equal(t, 0, Modulo(3, 3))
	assert.Equal(t, 2, Modulo(6, 4))
}

func TestXor(t *testing.T) {
	d1, _ := hex.DecodeString("8b84eb01")
	d2, _ := hex.DecodeString("2b7e1516")
	result, _ := hex.DecodeString("a0fafe17")

	assert.Equal(t, result, Xor(d1, d2))

}
