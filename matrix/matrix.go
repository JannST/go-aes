package matrix

import "encoding/hex"

const Nb = 4

type Matrix struct { //TODO change height to constant value
	data                []byte
	nk                  int
	height              int
	numberOfBytesPadded int
}

func newMatrix(numberOfRows int, length int) Matrix {
	var mat Matrix
	mat.nk = length
	mat.height = numberOfRows
	mat.data = make([]byte, numberOfRows*length)
	return mat
}

func (m Matrix) Size() int {
	return m.height * m.nk
}

func (m Matrix) Height() int {
	return m.height
}

func (m Matrix) Length() int {
	return m.nk
}

func (m Matrix) SameSize(other Matrix) bool {
	return other.Length() == m.nk && other.Height() == m.height
}

func (m Matrix) DebugString() string {
	var ret string
	for i := 0; i < m.nk; i++ {
		ret += hex.EncodeToString(m.Column(i))
		ret += " "
	}
	return ret
}
