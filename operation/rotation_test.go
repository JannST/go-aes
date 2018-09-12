package operation

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRotateLeft(t *testing.T) {
	w1 := []byte{0x00, 0x01, 0x02, 0x03}
	w2 := []byte{0x01, 0x02, 0x03, 0x00}
	w3 := []byte{0x02, 0x03, 0x00, 0x01}
	w4 := []byte{0x03, 0x00, 0x01, 0x02}

	t1 := []byte{0x00, 0x01, 0x02, 0x03}
	RotateLeft(&t1, 1)
	assert.Equal(t, w2, t1)

	t1 = []byte{0x00, 0x01, 0x02, 0x03}
	RotateLeft(&t1, 2)
	assert.Equal(t, w3, t1)

	t1 = []byte{0x00, 0x01, 0x02, 0x03}
	RotateLeft(&t1, 3)
	assert.Equal(t, w4, t1)

	t1 = []byte{0x00, 0x01, 0x02, 0x03}
	RotateLeft(&t1, 4)
	assert.Equal(t, w1, t1)

	t1 = []byte{0x00, 0x01, 0x02, 0x03}
	RotateLeft(&t1, -2)
	assert.Equal(t, w3, t1)

	t1 = []byte{0x00, 0x01, 0x02, 0x03}
	RotateLeft(&t1, 9)
	assert.Equal(t, w2, t1)
}

func TestRotateRight(t *testing.T) {
	w1 := []byte{0x00, 0x01, 0x02, 0x03}
	w2 := []byte{0x03, 0x00, 0x01, 0x02}
	w3 := []byte{0x02, 0x03, 0x00, 0x01}
	w4 := []byte{0x01, 0x02, 0x03, 0x00}


	t1 := []byte{0x00, 0x01, 0x02, 0x03}
	RotateRight(&t1, 1)
	assert.Equal(t, w2, t1)

	t1 = []byte{0x00, 0x01, 0x02, 0x03}
	RotateRight(&t1, 2)
	assert.Equal(t, w3, t1)

	t1 = []byte{0x00, 0x01, 0x02, 0x03}
	RotateRight(&t1, 3)
	assert.Equal(t, w4, t1)

	t1 = []byte{0x00, 0x01, 0x02, 0x03}
	RotateRight(&t1, 4)
	assert.Equal(t, w1, t1)

	t1 = []byte{0x00, 0x01, 0x02, 0x03}
	RotateRight(&t1, -2)
	assert.Equal(t, w3, t1)

	t1 = []byte{0x00, 0x01, 0x02, 0x03}
	RotateRight(&t1, 9)
	assert.Equal(t, w2, t1)
}