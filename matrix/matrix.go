package matrix

import "encoding/hex"

type Matrix struct { //TODO change height to constant value
	data                []byte
	length              int
	height              int
	numberOfBytesPadded int
}

func newMatrix(numberOfRows int, length int) Matrix {
	var mat Matrix
	mat.length = length
	mat.height = numberOfRows
	mat.data = make([]byte, numberOfRows*length)
	return mat
}

func NewMatrix(height int, length int) Matrix {
	var mat Matrix
	mat.length = length
	mat.height = height
	mat.data = make([]byte, height*length)
	return mat
}

func (m Matrix) Size() int {
	return m.height * m.length
}

func (m Matrix) Height() int {
	return m.height
}

func (m Matrix) Length() int {
	return m.length
}

func (m Matrix) SameSize(other Matrix) bool {
	return other.Length() == m.length && other.Height() == m.height
}

func (m Matrix) DebugString() string {
	var ret string
	for i := 0; i < m.length; i++ {
		ret += hex.EncodeToString(m.Column(i))
		ret += " "
	}
	return ret
}
