package word

import "testing"

func TestRotateLeft(t *testing.T) {
	w1 := Word{data: [4]byte{0x00, 0x01, 0x02, 0x03}}
	w2 := Word{data: [4]byte{0x01, 0x02, 0x03, 0x00}}
	w3 := Word{data: [4]byte{0x02, 0x03, 0x00, 0x01}}
	w4 := Word{data: [4]byte{0x03, 0x00, 0x01, 0x02}}

	t1 := Word{data: [4]byte{0x00, 0x01, 0x02, 0x03}}
	t1.RotateLeft(0)
	if t1 != w1 {
		t.Error(t1, "should be equal to", w1, "but it is not")
	}

	t1 = Word{data: [4]byte{0x00, 0x01, 0x02, 0x03}}
	t1.RotateLeft(1)
	if t1 != w2 {
		t.Error(t1, "should be equal to", w2, "but it is not")
	}

	t1 = Word{data: [4]byte{0x00, 0x01, 0x02, 0x03}}
	t1.RotateLeft(2)
	if t1 != w3 {
		t.Error(t1, "should be equal to", w3, "but it is not")
	}

	t1 = Word{data: [4]byte{0x00, 0x01, 0x02, 0x03}}
	t1.RotateLeft(3)
	if t1 != w4 {
		t.Error(t1, "should be equal to", w4, "but it is not")
	}

	t1 = Word{data: [4]byte{0x00, 0x01, 0x02, 0x03}}
	t1.RotateLeft(4)
	if t1 != w1 {
		t.Error(t1, "should be equal to", w1, "but it is not")
	}
}

func TestRotateRight(t *testing.T) {
	w1 := Word{data: [4]byte{0x00, 0x01, 0x02, 0x03}}
	w4 := Word{data: [4]byte{0x01, 0x02, 0x03, 0x00}}
	w3 := Word{data: [4]byte{0x02, 0x03, 0x00, 0x01}}
	w2 := Word{data: [4]byte{0x03, 0x00, 0x01, 0x02}}

	t1 := Word{data: [4]byte{0x00, 0x01, 0x02, 0x03}}
	t1.RotateRight(0)
	if t1 != w1 {
		t.Error(t1, "should be equal to", w1, "but it is not")
	}

	t1 = Word{data: [4]byte{0x00, 0x01, 0x02, 0x03}}
	t1.RotateRight(1)
	if t1 != w2 {
		t.Error(t1, "should be equal to", w2, "but it is not")
	}

	t1 = Word{data: [4]byte{0x00, 0x01, 0x02, 0x03}}
	t1.RotateRight(2)
	if t1 != w3 {
		t.Error(t1, "should be equal to", w3, "but it is not")
	}

	t1 = Word{data: [4]byte{0x00, 0x01, 0x02, 0x03}}
	t1.RotateRight(3)
	if t1 != w4 {
		t.Error(t1, "should be equal to", w4, "but it is not")
	}

	t1 = Word{data: [4]byte{0x00, 0x01, 0x02, 0x03}}
	t1.RotateRight(4)
	if t1 != w1 {
		t.Error(t1, "should be equal to", w1, "but it is not")
	}
}
