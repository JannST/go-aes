package algorithm

import (
	"encoding/hex"
	"fmt"
	"go-aes/math"
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
	temp := make([]byte, key.Height())

	for i := 1; i <= rounds; i++ {

		for j := 0; j < key.Length(); j++ {
			if j == 0 {
				copy(temp, result[i-1].Column(key.Length()-1))
				fmt.Println("copy", hex.EncodeToString(temp))
				RotWord(temp, -1)
				fmt.Println("Rot", hex.EncodeToString(temp))
				SubWord(temp)
				fmt.Println("Sub", hex.EncodeToString(temp))
				temp[0] ^= Rcon(i)
				fmt.Println("Rcon", hex.EncodeToString(temp))
			} else {
				copy(temp, result[i].Column(j-1))
			}

			if key.Length() > 6 && j%key.Length() == 4 {
				SubWord(temp)
				fmt.Println("Rot2", hex.EncodeToString(temp))
			}
			fmt.Println("xor", hex.EncodeToString(temp), hex.EncodeToString(result[i-1].Column(j)), hex.EncodeToString(math.Xor(temp, result[i-1].Column(j))))
			result[i].SetColumn(j, math.Xor(temp, result[i-1].Column(j)))
		}
		fmt.Println("=================================================")
		for _, e := range result {
			e.Print()
		}
	}
	return result
}
