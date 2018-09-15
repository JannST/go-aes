package matr

func (m matrix) XorWithColumn(index int, data []byte) {
	if m.height != len(data) {
		panic("eeek! different size")
	}

	for i := index * m.height; i < index*m.height+m.height; i++ {
		data[i-index*m.height] ^= m.data[i]
	}
}

func (m matrix) Column(index int) []byte {
	return m.data[index*m.height : index*m.height+m.height]
}

func (m *matrix) SetColumn(index int, data []byte) {
	for i := index * m.height; i < index*m.height+m.height; i++ {
		m.data[i] = data[i-index*m.height]
	}
}
