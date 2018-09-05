package word

import (
	"encoding/hex"
)

type Word struct {
	data [4]byte
}

func (w Word) toString() string {
	return hex.EncodeToString(w.data[:])
}