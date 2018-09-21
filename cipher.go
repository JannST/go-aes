package main

import (
	"github.com/pkg/errors"
	"go-aes/key_schedule"
	"go-aes/matrix"
	"io"
)

type KeyLength int

const (
	K128 KeyLength = 128
	K192 KeyLength = 192
	K256 KeyLength = 256
)

func EncryptFile(input io.Reader, output io.Writer, key []byte, klength int, rounds int) {
	if rounds < 0 {
		panic(errors.New("rounds can not be < 0"))
	}
	if klength%32 != 0 {
		panic(errors.New("key length is not dividable by 32"))
	}
	if len(key) != klength/8 {
		panic(errors.New("key doest not have size of (keyLength/8)"))
	}

	var (
		mat         = matrix.Matrix{}
		err         error
		close       bool
		keySchedule = key_schedule.ExpandKey(key, 4, klength/32, rounds)
	)

	for !close {
		err = mat.ReadFrom(input)
		if err != nil && err != io.EOF {
			panic(err)
		} else if err == io.EOF {
			close = true
		}

		mat.Encrypt(keySchedule, rounds)

		err = mat.WriteTo(output)
		if err != nil {
			panic(err)
		}
	}
}
