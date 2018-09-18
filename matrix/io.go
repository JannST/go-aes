package matrix

import (
	"github.com/pkg/errors"
	"io"
)

func (m *Matrix) ReadFrom(reader io.Reader) error {
	p := make([]byte, Size())
	n, err := reader.Read(p)
	if err != nil {
		return err
	}

	if n == 0 {
		return errors.New("null bytes read")
	}

	m.data = p
	return err
}

func (m Matrix) WriteTo(writer io.Writer) error {
	_, err := writer.Write(m.data)
	return err
}
