package matrix

func (m *matrix) RotateVertical(column int, times int) {
	col := m.data[column*m.numberOfRows : column*m.numberOfRows+m.numberOfRows]
	buffer := make([]byte, m.numberOfRows)
	copy(buffer, col)

	for i, element := range buffer {
		col[modulo(i+times, m.numberOfRows)] = element
	}
}

func (m *matrix) RotateHorizontal(row int, times int) {
	buffer := make([]byte, m.numberOfColumns)

	for i := 0; i < m.numberOfColumns; i++ {
		buffer[i] = m.data[row+m.numberOfRows*i]
	}

	for i := 0; i < m.numberOfColumns; i++ {
		m.data[row+m.numberOfRows*i] = buffer[modulo(i-times, m.numberOfRows)]
	}
}

func modulo(num int, base int) int {
	result := num % base

	if result < 0 {
		result = base + result
	}
	return result
}
