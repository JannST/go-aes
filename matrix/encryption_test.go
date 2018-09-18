package matrix

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	k "go-aes/key_schedule"
	"testing"
)

func TestMatrix_Encrypt128(t *testing.T) {
	result, _ := hex.DecodeString("3925841d02dc09fbdc118597196a0b32")
	key, _ := hex.DecodeString("2b7e151628aed2a6abf7158809cf4f3c")
	keySchedule := k.ExpandKey(key, Nb, 4, 10)

	mat := newMatrix(4, 4)
	mat.data, _ = hex.DecodeString("3243f6a8885a308d313198a2e0370734")
	mat.Encrypt(keySchedule, 10)
	assert.Equal(t, result, mat.data)
}

func TestMatrix_Encrypt192(t *testing.T) {
	result, _ := hex.DecodeString("dda97ca4864cdfe06eaf70a0ec0d7191")
	key, _ := hex.DecodeString("000102030405060708090a0b0c0d0e0f1011121314151617")
	keySchedule := k.ExpandKey(key, Nb, 6, 12)

	mat := newMatrix(4, 4)
	mat.data, _ = hex.DecodeString("00112233445566778899aabbccddeeff")
	mat.Encrypt(keySchedule, 12)
	assert.Equal(t, result, mat.data)
}

func TestMatrix_Encrypt256(t *testing.T) {
	result, _ := hex.DecodeString("8ea2b7ca516745bfeafc49904b496089")
	key, _ := hex.DecodeString("000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f")
	keySchedule := k.ExpandKey(key, Nb, 8, 14)

	mat := newMatrix(4, 4)
	mat.data, _ = hex.DecodeString("00112233445566778899aabbccddeeff")
	mat.Encrypt(keySchedule, 14)
	assert.Equal(t, result, mat.data)
}

func TestMatrix_MixColumns(t *testing.T) {
	result, _ := hex.DecodeString("046681e5e0cb199a48f8d37a2806264c")
	mat := newMatrix(4, 4)
	mat.data, _ = hex.DecodeString("d4bf5d30e0b452aeb84111f11e2798e5")
	mat.MixColumns()
	assert.Equal(t, result, mat.data)
}

func TestAddRoundKey(t *testing.T) {
	result, _ := hex.DecodeString("a4686b029c9f5b6a7f35ea50f22b4349")
	mat := newMatrix(4, 4)
	mat.data, _ = hex.DecodeString("04e0482866cbf8068119d326e59a7a4c")

	key, _ := hex.DecodeString("a088232afa54a36cfe2c397617b13905")

	mat.AddRoundKey(key)
	assert.Equal(t, result, mat.data)
}
