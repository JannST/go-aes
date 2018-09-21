package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestEncryptFile(t *testing.T) {
	//3925841d02dc09fbdc118597196a0b32
	inbytes, _ := hex.DecodeString("3243f6a8885a308d313198a2e03707343243f6a8885a308d313198a2e0370734")
	in := bytes.NewReader(inbytes)
	var out bytes.Buffer

	key, _ := hex.DecodeString("2b7e151628aed2a6abf7158809cf4f3c")
	EncryptFile(in, &out, key, 128, 10)

	fmt.Println(hex.EncodeToString(out.Bytes()))
}
