package matrix

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMatrix_GetColumn(t *testing.T) {
	mat := NewMatrix(4, 6)
	mat.setData([]byte("111122223333444455556666"))

	col := mat.GetColumn(0)
	assert.Equal(t, []byte("1111"), col)

	col = mat.GetColumn(5)
	assert.Equal(t, []byte("6666"), col)


	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	mat.GetColumn(6)
}

func TestMatrix_GetColumnIsChangable(t *testing.T) {
	mat := NewMatrix(4, 6)
	mat.setData([]byte("111122223333444455556666"))

	col := mat.GetColumn(0)
	col[0] = '7'
	assert.Equal(t, []byte("711122223333444455556666"), mat.getData())
}