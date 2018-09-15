package matr

import (
	"encoding/hex"
	"fmt"
)

func (m matrix) Print() {
	for i := 0; i < m.length; i++ {
		fmt.Print(hex.EncodeToString(m.Column(i)), " ")
	}
	fmt.Print("\n")
}
