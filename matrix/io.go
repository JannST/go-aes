package matrix

import (
	"fmt"
	"io"
)

func (m *Matrix) ReadFrom(reader io.Reader) error {
	p := make([]byte, Size())
	n, err := reader.Read(p)
	fmt.Println(n, err)
	if err != nil {
		return err
	}

	m.data = p
	return err
}

func (m Matrix) WriteTo(writer io.Writer) error {
	_, err := writer.Write(m.data)
	return err
}
