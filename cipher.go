package main

import (
	"github.com/pkg/errors"
	"go-aes/key_schedule"
	"go-aes/matrix"
	"io"
)

func Encrypt(input io.Reader, output io.Writer, key []byte, klength int, rounds int) {
	if rounds < 0 {
		panic(errors.New("rounds can not be < 0"))
	}
	if klength%32 != 0 {
		panic(errors.New("key length is not dividable by 32"))
	}
	if len(key) != klength/8 {
		panic(errors.New("key does not have size of (keyLength/8)"))
	}

	var (
		mat         = matrix.Matrix{}
		err         error
		keySchedule = key_schedule.ExpandKey(key, 4, klength/32, rounds)
	)

	for {
		err = mat.ReadFrom(input)

		if err != nil && err != io.EOF {
			panic(err)
		} else if err == io.EOF {
			break
		}
		mat.Encrypt(keySchedule, rounds)

		err = mat.WriteTo(output)
		if err != nil {
			panic(err)
		}
	}
}

func Decrypt(input io.Reader, output io.Writer, key []byte, klength int, rounds int) {
	if rounds < 0 {
		panic(errors.New("rounds can not be < 0"))
	}
	if klength%32 != 0 {
		panic(errors.New("key length is not dividable by 32"))
	}
	if len(key) != klength/8 {
		panic(errors.New("key does not have size of (keyLength/8)"))
	}

	var (
		mat         = matrix.Matrix{}
		err         error
		keySchedule = key_schedule.ExpandKey(key, 4, klength/32, rounds)
	)

	for {
		err = mat.ReadFrom(input)

		if err != nil && err != io.EOF {
			panic(err)
		} else if err == io.EOF {
			break
		}
		mat.Decrypt(keySchedule, rounds)

		err = mat.WriteTo(output)
		if err != nil {
			panic(err)
		}
	}
}
