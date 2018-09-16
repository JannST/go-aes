package matrix

import (
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatrix_ShiftRows(t *testing.T) {
	result, _ := hex.DecodeString("00010203010203000203000103000102")
	fmt.Println(result)
	mat := newMatrix(4, 4)
	mat.data, _ = hex.DecodeString("00000000010101010202020203030303")
	mat.ShiftRows()
	assert.Equal(t, result, mat.data)
}

func TestRotateVertical(t *testing.T) {
	mat := newMatrix(4, 4)
	mat.setData([]byte("1234abcd5678efgh"))
	mat.RotWord(0, 1)
	mat.RotWord(3, -2)
	mat.RotWord(1, 4)
	mat.RotWord(2, -1)

	result := []byte("4123abcd6785ghef")
	assert.Equal(t, result, mat.data)
}
