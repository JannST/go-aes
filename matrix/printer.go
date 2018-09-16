package matrix

import (
	"encoding/hex"
	"fmt"
)

func (m Matrix) Print() {
	for i := 0; i < m.length; i++ {
		fmt.Print(hex.EncodeToString(m.Column(i)), " ")
	}
	fmt.Print("\n")
}
