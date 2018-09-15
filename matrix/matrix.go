package matrix

type matrix struct {
	data                []byte
	numberOfColumns     int
	numberOfRows        int
	numberOfBytesPadded int
}

func NewMatrix(numberOfRows int, length int) matrix {
	var mat matrix
	mat.numberOfColumns = length
	mat.numberOfRows = numberOfRows
	mat.data = make([]byte, numberOfRows*length)
	return mat
}

func (m matrix) GetSize() int {
	return m.numberOfRows * m.numberOfColumns
}
