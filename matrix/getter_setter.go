package matrix

func (m *Matrix) setData(buffer []byte) {
	if len(buffer) != m.Size() {
		panic("nk of data is not equal to size")
	}
	copy(m.data, buffer)
}

func (m Matrix) Data() []byte {
	return m.data
}
