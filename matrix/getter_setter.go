package matrix

import (
	"github.com/pkg/errors"
)

func (m matrix) setData(buffer []byte) error {
	if len(buffer) != m.GetSize() {
		return errors.New("length of data is not equal to size")
	}
	bufferIndex := 0

	for i := 0; i < m.length; i++ {
		for j := 0; j < m.numberOfRows; j++ {
			m.data[j][i] = buffer[bufferIndex]
			bufferIndex++
		}
	}
	return nil
}

func (m matrix) getData() []byte {
	buffer := make([]byte, m.GetSize())
	bufferIndex := 0

	for i := 0; i < m.length; i++ {
		for j := 0; j < m.numberOfRows; j++ {
			buffer[bufferIndex] = m.data[j][i]
			bufferIndex++
		}
	}
	return buffer
}
