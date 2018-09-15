package keyexpand

import (
	"fmt"
	"go-aes/matr"
)

func ExpandKey(key matr.Matrix, rounds int) []matr.Matrix {
	result := make([]matr.Matrix, rounds+1)
	var nm matr.Matrix
	for i := 1; i < len(result); i++ {
		nm = matr.NewMatrix(key.Height(), key.Length())
		result[i] = nm
	}

	result[0] = key
	var temp matr.Matrix

	for i := 1; i <= rounds; i++ {

		fmt.Println("=================================================")
		for _, e := range result {
			e.Print()
		}

		copy(temp, result[i-1])
		var col []byte

		temp.RotateVertical(temp.Length()-1, -1)
		temp.SubColumn(temp.Length() - 1)
		temp.XorColumnRcon(temp.Length()-1, i/key.Length())

		for j := 0; j < temp.Length(); j++ {
			if key.Length() > 6 && j%key.Length() == 4 {
				temp.SubColumn(j)
			}

			if j == 0 {
				col = temp.Column(temp.Length() - 1)
			} else {
				col = result[i].Column(j - 1)
			}
			fmt.Println(j, j, j, j)
			temp.XorWithColumn(j, col)
			result[i].SetColumn(j, col)
		}

	}
	return result
}
