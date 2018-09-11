package matrix

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatrix_setData(t *testing.T) {
	testBytes := []byte("abcdefghi")
	resultData := [][]byte{{'a', 'd', 'g'}, {'b', 'e', 'h'}, {'c', 'f', 'i'}}

	mat := NewMatrix(3, 3)
	mat.setData(testBytes)
	assert.Equal(t, mat.data, resultData)
}

func TestMatrix_GetData(t *testing.T) {
	mat := NewMatrix(3, 3)
	mat.data = [][]byte{{'a', 'd', 'g'}, {'b', 'e', 'h'}, {'c', 'f', 'i'}}

	result := []byte("abcdefghi")
	output := mat.getData()
	assert.Equal(t, output, result)
}

func TestMatrix_SetAndGet(t *testing.T) {
	testBytes := []byte("123456789abc")

	mat := NewMatrix(3, 4)
	mat.setData(testBytes)
	fmt.Println(mat.data)

	output := mat.getData()

	assert.Equal(t, testBytes, output)
}
