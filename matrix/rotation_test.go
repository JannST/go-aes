package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModulo(t *testing.T) {
	assert.Equal(t, 3, modulo(-1, 4))
	assert.Equal(t, 5, modulo(-2, 7))
	assert.Equal(t, 0, modulo(3, 3))
	assert.Equal(t, 2, modulo(6, 4))
}

func TestRotateVertical(t *testing.T) {
	mat := NewMatrix(4, 4)
	mat.setData([]byte("1234abcd5678efgh"))
	mat.RotateVertical(0, 1)
	mat.RotateVertical(3, -2)
	mat.RotateVertical(1, 4)
	mat.RotateVertical(2, -1)

	result := []byte("4123abcd6785ghef")
	assert.Equal(t, result, mat.getData())
}

func TestRotateHorizontal(t *testing.T) {
	mat := NewMatrix(4, 4)
	mat.setData([]byte("1234abcd5678efgh"))
	mat.RotateHorizontal(0, 1)
	mat.RotateHorizontal(1, -2)
	mat.RotateHorizontal(3, 4)

	result := []byte("e6341fcda2785bgh")
	assert.Equal(t, result, mat.getData())
}
