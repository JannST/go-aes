package matr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatrix_setData(t *testing.T) {
	testBytes := []byte("abcdefghi")
	resultData := []byte("abcdefghi")
	mat := newMatrix(3, 3)
	mat.setData(testBytes)
	assert.Equal(t, resultData, mat.data)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	testBytes = []byte("abcdefgh")
	mat = newMatrix(3, 3)
	mat.setData(testBytes)
}

func TestMatrix_GetData(t *testing.T) {
	mat := newMatrix(3, 3)

	assert.Equal(t, make([]byte, 3*3), mat.getData())
	mat.data = []byte("abcdefghi")

	result := []byte("abcdefghi")
	output := mat.getData()
	assert.Equal(t, output, result)
}

func TestMatrix_SetAndGet(t *testing.T) {
	testBytes := []byte("123456789abc")

	mat := newMatrix(3, 4)
	mat.setData(testBytes)

	output := mat.getData()

	assert.Equal(t, []byte("123456789abc"), output)
}
