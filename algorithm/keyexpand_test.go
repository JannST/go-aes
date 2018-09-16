package algorithm

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"go-aes/matr"
	"testing"
)

func TestExpandKey(t *testing.T) {
	mat := matr.NewMatrix(4, 4)
	key, _ := hex.DecodeString("2b7e151628aed2a6abf7158809cf4f3c")
	mat.ReadFrom(bytes.NewReader(key))
	epanded := ExpandKey(mat, 10)
	fmt.Println(epanded)
}
