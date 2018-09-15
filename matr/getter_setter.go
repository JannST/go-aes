package matr

func (m *matrix) setData(buffer []byte) {
	if len(buffer) != m.Size() {
		panic("length of data is not equal to size")
	}
	copy(m.data, buffer)
}

func (m matrix) getData() []byte {
	return m.data
}
