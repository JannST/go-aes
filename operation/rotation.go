package operation

func RotateLeft(input *[]byte, times int) {
	if times < 0 {
		times *= -1
	}
	length := len(*input)
	times = times % length
	buffer := make([]byte, length)
	copy(buffer, *input)

	for i, element := range buffer {
		(*input)[modulo(i-times, length)] = element
	}
}

func RotateRight(input *[]byte, times int) {
	if times < 0 {
		times *= -1
	}
	length := len(*input)
	times = times % length
	buffer := make([]byte, length)
	copy(buffer, *input)

	for i, element := range buffer {
		(*input)[modulo(i+times, length)] = element
	}
}

func modulo(num int, base int) int {
	result := num % base

	if result < 0 {
		result = base + result
	}

	return result
}
