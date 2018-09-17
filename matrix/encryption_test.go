package matrix

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatrix_Encrypt(t *testing.T) {
	key := newMatrix(4, 4)
	key.data, _ = hex.DecodeString("2b7e151628aed2a6abf7158809cf4f3c")
	key_schedule := key.ExpandKey(10)

	mat := newMatrix(4, 4)
	mat.data, _ = hex.DecodeString("3243f6a8885a308d313198a2e0370734")
	mat.Encrypt(key_schedule, 10)
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
	key := newMatrix(4, 4)
	key.data, _ = hex.DecodeString("a088232afa54a36cfe2c397617b13905")

	mat.AddRoundKey(key)
	assert.Equal(t, result, mat.data)
}
