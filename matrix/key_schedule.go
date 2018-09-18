package matrix

import (
	"go-aes/math"
	"go-aes/tables"
)

func (m *Matrix) ExpandKey(rounds int) []Matrix {
	result := make([]Matrix, rounds+1)
	var nm Matrix
	for i := 1; i < len(result); i++ {
		nm = NewMatrix(m.Height(), m.Length())
		result[i] = nm
	}

	result[0] = *m
	temp := make([]byte, m.Height())

	for i := 1; i <= rounds; i++ {

		for j := 0; j < m.Length(); j++ {
			if j == 0 {
				copy(temp, result[i-1].Column(m.Length()-1))
				RotBytes(temp, -1)
				tables.SubWord(temp)
				temp[0] ^= tables.Rcon(i)
			} else {
				copy(temp, result[i].Column(j-1))
			}

			if m.Length() > 6 && j%m.Length() == 4 {
				tables.SubWord(temp)
			}
			result[i].SetColumn(j, math.Xor(temp, result[i-1].Column(j)))
		}
	}
	return result
}

func RotBytes(data []byte, steps int) {
	buffer := make([]byte, len(data))
	copy(buffer, data)

	for i := 0; i < len(data); i++ {
		data[math.Modulo(i+steps, len(data))] = buffer[i]
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
