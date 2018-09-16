package matrix

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
