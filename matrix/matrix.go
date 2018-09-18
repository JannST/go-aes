package matrix

const Nb = 4

type Matrix struct {
	data []byte
}

func NewMatrix() Matrix {
	var mat Matrix
	mat.data = make([]byte, Size())
	return mat
}

func Size() int {
	return Nb * Nb
}

/*
func (m Matrix) SameSize(other Matrix) bool {
	return other.Nk() == m.nk && other.Height() == m.height
}


func (m Matrix) DebugString() string {
	var ret string
	for i := 0; i < Nb; i++ {
		ret += hex.EncodeToString(m.Column(i))
		ret += " "
	}
	return ret
}
*/
