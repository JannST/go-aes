package matr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotateVertical(t *testing.T) {
	mat := newMatrix(4, 4)
	mat.setData([]byte("1234abcd5678efgh"))
	mat.RotateVertical(0, 1)
	mat.RotateVertical(3, -2)
	mat.RotateVertical(1, 4)
	mat.RotateVertical(2, -1)

	result := []byte("4123abcd6785ghef")
	assert.Equal(t, result, mat.getData())
}

func TestRotateHorizontal(t *testing.T) {
	mat := newMatrix(4, 4)
	mat.setData([]byte("1234abcd5678efgh"))
	mat.RotateHorizontal(0, 1)
	mat.RotateHorizontal(1, -2)
	mat.RotateHorizontal(3, 4)

	result := []byte("e6341fcda2785bgh")
	assert.Equal(t, result, mat.getData())
}
