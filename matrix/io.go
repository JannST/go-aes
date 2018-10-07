package matrix

import (
	"io"
)

func (m *Matrix) ReadFrom(reader io.Reader) error {
	p := make([]byte, Size())
	n, err := reader.Read(p)
	if err == io.EOF && n == 0 {
		return err
	} else if err != nil && err != io.EOF {
		return err
	}

	m.data = p
	return err
}

func (m Matrix) WriteTo(writer io.Writer) error {
	_, err := writer.Write(m.data)
	return err
}
