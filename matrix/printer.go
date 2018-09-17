package matrix

import (
	"encoding/hex"
	"fmt"
)

func (m Matrix) Tos() string {
	var ret string
	for i := 0; i < m.length; i++ {
		ret += hex.EncodeToString(m.Column(i))
		ret += " "
	}
	return ret
}

func (m Matrix) Print() {
	fmt.Println(m.Tos())
}
