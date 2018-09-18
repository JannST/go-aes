package matrix

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestMatrix_Read(t *testing.T) {
	result := []byte("1111222233334444")
	reader := strings.NewReader("1111222233334444")

	mat := NewMatrix()
	err := mat.ReadFrom(reader)
	assert.Equal(t, nil, err)
	assert.Equal(t, result, mat.data)

	reader = strings.NewReader("")
	mat = NewMatrix()
	err = mat.ReadFrom(reader)
	assert.NotEqual(t, nil, err)
}

func TestMatrix_WriteTo(t *testing.T) {
	var writer bytes.Buffer
	result := "1111222233334444"
	mat := NewMatrix()
	mat.data = []byte("1111222233334444")
	mat.WriteTo(&writer)
	assert.Equal(t, result, writer.String())
}

func TestMatrix_WriteRead(t *testing.T) {
	inp := []byte("1111222233334")
	result := []byte("1111222233334\x00\x00\x00")
	reader := bytes.NewReader(inp)

	mat := NewMatrix()
	mat.ReadFrom(reader)

	var writer bytes.Buffer
	mat.WriteTo(&writer)
	assert.Equal(t, result, writer.Bytes())
}
