package matrix

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"github.com/JannST/go-aes/key_schedule"
	"testing"
)

func TestMatrix_Decrypt128(t *testing.T) {
	result, _ := hex.DecodeString("3243f6a8885a308d313198a2e0370734")
	key, _ := hex.DecodeString("2b7e151628aed2a6abf7158809cf4f3c")
	keySchedule := key_schedule.ExpandKey(key, Nb, 4, 10)

	mat := NewMatrix()
	mat.data, _ = hex.DecodeString("3925841d02dc09fbdc118597196a0b32")
	mat.Decrypt(keySchedule, 10)
	assert.Equal(t, result, mat.data)
}

func TestMatrix_Decrypt192(t *testing.T) {
	result, _ := hex.DecodeString("00112233445566778899aabbccddeeff")
	key, _ := hex.DecodeString("000102030405060708090a0b0c0d0e0f1011121314151617")
	keySchedule := key_schedule.ExpandKey(key, Nb, 6, 12)

	mat := NewMatrix()
	mat.data, _ = hex.DecodeString("dda97ca4864cdfe06eaf70a0ec0d7191")
	mat.Decrypt(keySchedule, 12)
	assert.Equal(t, result, mat.data)
}

func TestMatrix_Decrypt256(t *testing.T) {
	result, _ := hex.DecodeString("00112233445566778899aabbccddeeff")
	key, _ := hex.DecodeString("000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f")
	keySchedule := key_schedule.ExpandKey(key, Nb, 8, 14)

	mat := NewMatrix()
	mat.data, _ = hex.DecodeString("8ea2b7ca516745bfeafc49904b496089")
	mat.Decrypt(keySchedule, 14)
	assert.Equal(t, result, mat.data)
}

func TestMatrix_ShiftRows(t *testing.T) {
	result, _ := hex.DecodeString("00010203010203000203000103000102")
	mat := NewMatrix()
	mat.data, _ = hex.DecodeString("00000000010101010202020203030303")
	mat.ShiftRows()
	assert.Equal(t, result, mat.data)
}

func TestMatrix_InvShiftRows(t *testing.T) {
	result, _ := hex.DecodeString("00000000010101010202020203030303")
	mat := NewMatrix()
	mat.data, _ = hex.DecodeString("00010203010203000203000103000102")
	mat.InvShiftRows()
	assert.Equal(t, result, mat.data)
}

func TestMatrix_InvMixColumns(t *testing.T) {
	result, _ := hex.DecodeString("d4bf5d30e0b452aeb84111f11e2798e5")
	mat := NewMatrix()
	mat.data, _ = hex.DecodeString("046681e5e0cb199a48f8d37a2806264c")
	mat.InvMixColumns()
	assert.Equal(t, result, mat.data)
}
