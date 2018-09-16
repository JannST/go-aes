package basic

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddRoundKey(t *testing.T) {
	state, _ := hex.DecodeString("04e0482866cbf8068119d326e59a7a4c")
	key, _ := hex.DecodeString("a088232afa54a36cfe2c397617b13905")
	result, _ := hex.DecodeString("a4686b029c9f5b6a7f35ea50f22b4349")

	AddRoundKey(state, key)
	assert.Equal(t, result, state)
}
