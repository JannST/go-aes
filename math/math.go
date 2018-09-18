package math

func Modulo(num int, base int) int {
	result := num % base

	if result < 0 {
		result = base + result
	}
	return result
}

//Xors tow byte slices bitwise. Returns new slice and will not modify input
func Xor(data1 []byte, data2 []byte) []byte {
	if len(data1) != len(data2) {
		panic("different length of input")
	}

	result := make([]byte, len(data1))

	for i := 0; i < len(data1); i++ {
		result[i] = data1[i] ^ data2[i]
	}
	return result
}
