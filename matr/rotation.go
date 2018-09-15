package matr

import "go-aes/math"

func (m *matrix) RotateVertical(column int, times int) {
	col := m.data[column*m.height : column*m.height+m.height]
	buffer := make([]byte, m.height)
	copy(buffer, col)

	for i, element := range buffer {
		col[math.Modulo(i+times, m.height)] = element
	}
}

func (m *matrix) RotateHorizontal(row int, times int) {
	buffer := make([]byte, m.length)

	for i := 0; i < m.length; i++ {
		buffer[i] = m.data[row+m.height*i]
	}

	for i := 0; i < m.length; i++ {
		m.data[row+m.height*i] = buffer[math.Modulo(i-times, m.height)]
	}
}
