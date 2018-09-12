package matrix

func (m *matrix) GetColumn(index int) []byte {
	if index >= m.length || index < 0 {
		panic("out of bounds")
	}
	return m.data[index*m.numberOfRows:index*m.numberOfRows+m.numberOfRows]
}
