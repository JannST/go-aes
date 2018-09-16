package matrix

import (
	"go-aes/algorithm/basic"
	"go-aes/math"
)

func (m *Matrix) ShiftRows() {
	buf := make([]byte, len(m.data))
	var dst int
	copy(buf, m.data)

	for row := 1; row < m.height; row++ {
		for col := 0; col < m.length; col++ {
			dst = math.Modulo(col-row, m.length)
			m.data[dst*m.height+row] = buf[col*m.height+row]
		}
	}
}

func (m *Matrix) RotWord(column int, times int) {
	col := m.data[column*m.height : column*m.height+m.height]
	buffer := make([]byte, m.height)
	copy(buffer, col)

	for i, element := range buffer {
		col[math.Modulo(i+times, m.height)] = element
	}
}

func (m *Matrix) SubBytes() {
	basic.SubWord(m.data)
}
