package matr

type Matrix interface {
	Size() int
	Height() int
	Length() int
	Data() []byte
	Print()
	RotateHorizontal(row int, times int)
	RotateVertical(column int, times int)
	SubColumn(index int)
	InvSubColumn(index int)
	XorColumnRcon(index int, rcon int)
	XorWithColumn(index int, data []byte)
	Column(index int) []byte
	SetColumn(index int, data []byte)
}

type matrix struct {
	data                []byte
	length              int
	height              int
	numberOfBytesPadded int
}

func newMatrix(numberOfRows int, length int) matrix {
	var mat matrix
	mat.length = length
	mat.height = numberOfRows
	mat.data = make([]byte, numberOfRows*length)
	return mat
}

func NewMatrix(numberOfRows int, length int) Matrix {
	var mat matrix
	mat.length = length
	mat.height = numberOfRows
	mat.data = make([]byte, numberOfRows*length)
	return &mat
}

func (m matrix) Size() int {
	return m.height * m.length
}

func (m matrix) Height() int {
	return m.height
}

func (m matrix) Length() int {
	return m.length
}

func (m matrix) Data() []byte {
	return m.data
}
