package matrix

import (
	"errors"
	"fmt"
	"io"
)

type matrix struct {
	data                [][]byte
	length              int
	numberOfRows        int
	numberOfBytesPadded int
}

func NewMatrix(numberOfRows int, length int) matrix {
	var mat matrix
	mat.length = length
	mat.numberOfRows = numberOfRows
	mat.data = make([][]byte, numberOfRows)
	for i := 0; i < numberOfRows; i++ {
		mat.data[i] = make([]byte, length)
	}
	return mat
}

func (m matrix) GetSize() int {
	return m.numberOfRows * m.length
}

func (m matrix) Read(reader io.Reader) error {
	size := m.GetSize()
	if size == 0 {
		return errors.New("size is 0")
	}

	buffer := make([]byte, size)

	numberOfBytesRead, err := reader.Read(buffer)
	if err != nil {
		if err == io.EOF {
			m.numberOfBytesPadded = size - numberOfBytesRead
		}
	}
	fmt.Println(string(buffer[:numberOfBytesRead]), numberOfBytesRead)
	err = m.setData(buffer)
	return err
}
