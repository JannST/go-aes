package matr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatrix_Column(t *testing.T) {
	mat := newMatrix(4, 6)
	mat.setData([]byte("111122223333444455556666"))

	col := mat.Column(0)
	assert.Equal(t, []byte("1111"), col)

	col = mat.Column(5)
	assert.Equal(t, []byte("6666"), col)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	mat.Column(6)
}

func TestMatrix_ColumnIsChangable(t *testing.T) {
	mat := newMatrix(4, 6)
	mat.setData([]byte("111122223333444455556666"))

	col := mat.Column(0)
	col[0] = '7'
	assert.NotEqual(t, []byte("711122223333444455556666"), mat.getData())
}

func TestMatrix_SetColumn(t *testing.T) {
	mat := newMatrix(4, 4)
	mat.setData([]byte("1111222233334444"))
	mat.SetColumn(1, []byte("ffff"))
	assert.Equal(t, []byte("1111ffff33334444"), mat.data)
}
