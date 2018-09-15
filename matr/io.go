package matr

import (
	"github.com/pkg/errors"
	"io"
)

func (m *matrix) ReadFrom(reader io.Reader) error {
	size := m.Size()
	if size == 0 {
		return errors.New("size of matr is 0")
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

func (m matrix) WriteTo(writer io.Writer) error {
	_, err := writer.Write(m.getData())
	return err
}
