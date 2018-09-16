package algorithm

import "go-aes/math"

func RotWord(data []byte, steps int) {
	buffer := make([]byte, len(data))
	copy(buffer, data)

	for i := 0; i < len(data); i++ {
		data[math.Modulo(i+steps, len(data))] = buffer[i]
	}
}
