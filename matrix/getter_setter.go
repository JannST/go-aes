package matrix

func (m *matrix) setData(buffer []byte) {
	if len(buffer) != m.GetSize() {
		panic("numberOfColumns of data is not equal to size")
	}
	copy(m.data, buffer)
}

func (m matrix) getData() []byte {
	return m.data
}
