package matrix

import (
	"github.com/pkg/errors"
	"io"
)

func (m *Matrix) ReadFrom(reader io.Reader) error {
	size := m.Size()
	if size == 0 {
		return errors.New("size of matrix is 0")
	}

	p := make([]byte, size)
	n, err := reader.Read(p)
	if err != nil {
		return err
	}

	if n == 0 {
		return errors.New("null bytes read")
	}

	m.numberOfBytesPadded = size - n
	m.setData(p)
	return err
}

func (m Matrix) WriteTo(writer io.Writer) error {
	_, err := writer.Write(m.Data())
	return err
}
