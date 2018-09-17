package matrix

import (
	"fmt"
	"go-aes/math"
	s "go-aes/tables"
)

func (m *Matrix) Encrypt(keys []Matrix, rounds int) {
	if len(keys) != rounds+1 {
		panic("len(keys) != rounds+1")
	}

	fmt.Println("==========start keys===========")
	for _, key := range keys {
		key.Print()
	}
	fmt.Println("==========End keys===========")

	m.AddRoundKey(keys[0])
	fmt.Println("aika", m.Tos())
	for round := 1; round < rounds; round++ {
		m.SubBytes()
		fmt.Println(round, "SubBytes", m.Tos())
		m.ShiftRows()
		fmt.Println(round, "ShiftRows", m.Tos())
		m.MixColumns()
		fmt.Println(round, "MixColumns", m.Tos())
		m.AddRoundKey(keys[round])
		fmt.Println(round, "Adding key -->", keys[round].Tos(), "\n")
		fmt.Println(round, "AddRoundKey", m.Tos())
	}
	fmt.Println("final round :o")
	m.SubBytes()
	fmt.Println("SubBytes", m.Tos())
	m.ShiftRows()
	fmt.Println("ShiftRows", m.Tos())
	m.AddRoundKey(keys[rounds])
	fmt.Println("AddRoundKey", m.Tos())
}

func (m *Matrix) SubBytes() {
	s.SubWord(m.data)
}

func (m *Matrix) ShiftRows() {
	buf := make([]byte, len(m.data))
	var dst int
	copy(buf, m.data)

	for row := 1; row < m.height; row++ {
		for col := 0; col < m.length; col++ {
			dst = math.Modulo(col-row, m.length)
			m.data[dst*m.height+row] = buf[col*m.height+row]
		}
	}
}

func (m *Matrix) MixColumns() {
	buf := [4]byte{}
	for i := 0; i < m.Size(); i += 4 {
		buf[0] = s.Mul2(m.data[i]) ^ s.Mul3(m.data[i+1]) ^ m.data[i+2] ^ m.data[i+3]
		buf[1] = m.data[i] ^ s.Mul2(m.data[i+1]) ^ s.Mul3(m.data[i+2]) ^ m.data[i+3]
		buf[2] = m.data[i] ^ m.data[i+1] ^ s.Mul2(m.data[i+2]) ^ s.Mul3(m.data[i+3])
		buf[3] = s.Mul3(m.data[i]) ^ m.data[i+1] ^ m.data[i+2] ^ s.Mul2(m.data[i+3])
		m.data[i] = buf[0]
		m.data[i+1] = buf[1]
		m.data[i+2] = buf[2]
		m.data[i+3] = buf[3]
	}
}

func (m *Matrix) AddRoundKey(key Matrix) {
	if m.length != key.length || m.length != key.length {
		panic("input matrix is not of same size")
	}

	for i := 0; i < len(m.data); i++ {
		m.data[i] ^= key.data[i]
	}
}
