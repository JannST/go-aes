package matr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRcon(t *testing.T) {
	mat := newMatrix(4, 4)
	mat.data[0] = 0x8a
	mat.data[1] = 0x84
	mat.data[2] = 0xeb
	mat.data[3] = 0x01
	mat.RconXORColumn(0, 4/mat.length)
	assert.Equal(t, []byte{0x8b, 0x84, 0xeb, 0x01}, mat.data[0:mat.height])
}
