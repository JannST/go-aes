package matrix

func (m Matrix) Column(index int) []byte {
	return m.data[index*m.height : index*m.height+m.height]
}

func (m *Matrix) SetColumn(index int, data []byte) {
	for i := index * m.height; i < index*m.height+m.height; i++ {
		m.data[i] = data[i-index*m.height]
	}
}
