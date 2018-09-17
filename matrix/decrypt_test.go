package matrix

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatrix_Decrypt(t *testing.T) {
	key := newMatrix(4, 4)
	key.data, _ = hex.DecodeString("2b7e151628aed2a6abf7158809cf4f3c")
	key_schedule := key.ExpandKey(10)

	mat := newMatrix(4, 4)
	mat.data, _ = hex.DecodeString("3925841d02dc09fbdc118597196a0b32")
	mat.Decrypt(key_schedule, 10)
}

func TestMatrix_ShiftRows(t *testing.T) {
	result, _ := hex.DecodeString("00010203010203000203000103000102")
	mat := newMatrix(4, 4)
	mat.data, _ = hex.DecodeString("00000000010101010202020203030303")
	mat.ShiftRows()
	assert.Equal(t, result, mat.data)
}

func TestMatrix_InvShiftRows(t *testing.T) {
	result, _ := hex.DecodeString("00000000010101010202020203030303")
	mat := newMatrix(4, 4)
	mat.data, _ = hex.DecodeString("00010203010203000203000103000102")
	mat.InvShiftRows()
	assert.Equal(t, result, mat.data)
}

func TestMatrix_InvMixColumns(t *testing.T) {
	result, _ := hex.DecodeString("d4bf5d30e0b452aeb84111f11e2798e5")
	mat := newMatrix(4, 4)
	mat.data, _ = hex.DecodeString("046681e5e0cb199a48f8d37a2806264c")
	mat.InvMixColumns()
	assert.Equal(t, result, mat.data)
}
