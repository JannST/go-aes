package matrix

import (
	"fmt"
	"go-aes/math"
	s "go-aes/tables"
)

func (m *Matrix) Decrypt(keys []Matrix, rounds int) {
	if len(keys) != rounds+1 {
		panic("len(keys) != rounds+1")
	}

	fmt.Println("==========start keys===========")
	for _, key := range keys {
		key.Print()
	}
	fmt.Println("==========End keys===========")

	fmt.Println("Initial State:", m.Tos())

	m.AddRoundKey(keys[rounds])
	fmt.Println("AddRoundKey", m.Tos())

	for round := rounds - 1; round >= 1; round-- {
		m.InvShiftRows()
		fmt.Println(round, "ShiftRows", m.Tos())
		m.InvSubBytes()
		fmt.Println(round, "SubBytes", m.Tos())
		m.AddRoundKey(keys[round])
		fmt.Println(round, "Adding key", keys[round].Tos())
		fmt.Println(round, "AddRoundKey", m.Tos())
		m.InvMixColumns()
		fmt.Println(round, "MixColumns", m.Tos())
	}

	m.InvShiftRows()
	fmt.Println("ShiftRows", m.Tos())
	m.InvSubBytes()
	fmt.Println("SubBytes", m.Tos())
	m.AddRoundKey(keys[0])
	fmt.Println("AddRoundKey", m.Tos())

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
