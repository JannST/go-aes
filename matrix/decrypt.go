package matrix

import (
	"go-aes/math"
	s "go-aes/tables"
)

func (m *Matrix) Decrypt(keys []Matrix, rounds int) {
	if len(keys) != rounds+1 {
		panic("len(keys) != rounds+1")
	}

	m.AddRoundKey(keys[rounds])

	for round := rounds - 1; round >= 1; round-- {
		m.InvShiftRows()
		m.InvSubBytes()
		m.AddRoundKey(keys[round])
		m.InvMixColumns()
	}

	m.InvShiftRows()
	m.InvSubBytes()
	m.AddRoundKey(keys[0])
}

func (m *Matrix) InvSubBytes() {
	s.InvSubWord(m.data)
}

func (m *Matrix) InvShiftRows() {
	buf := make([]byte, len(m.data))
	var dst int
	copy(buf, m.data)

	for row := 1; row < m.height; row++ {
		for col := 0; col < m.length; col++ {
			dst = math.Modulo(col+row, m.length)
			m.data[dst*m.height+row] = buf[col*m.height+row]
		}
	}
}

func (m *Matrix) InvMixColumns() {
	buf := [4]byte{}
	for i := 0; i < m.Size(); i += 4 {
		buf[0] = s.Mul14(m.data[i]) ^ s.Mul11(m.data[i+1]) ^ s.Mul13(m.data[i+2]) ^ s.Mul9(m.data[i+3])
		buf[1] = s.Mul9(m.data[i]) ^ s.Mul14(m.data[i+1]) ^ s.Mul11(m.data[i+2]) ^ s.Mul13(m.data[i+3])
		buf[2] = s.Mul13(m.data[i]) ^ s.Mul9(m.data[i+1]) ^ s.Mul14(m.data[i+2]) ^ s.Mul11(m.data[i+3])
		buf[3] = s.Mul11(m.data[i]) ^ s.Mul13(m.data[i+1]) ^ s.Mul9(m.data[i+2]) ^ s.Mul14(m.data[i+3])
		m.data[i] = buf[0]
		m.data[i+1] = buf[1]
		m.data[i+2] = buf[2]
		m.data[i+3] = buf[3]
	}
}
