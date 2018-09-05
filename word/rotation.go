package word

func (w *Word) RotateLeft(times int) {
	length := len(w.data)
	var result [4]byte

	for i, element := range w.data {
		result[modulo(i-times, length)] = element
	}
	w.data = result
}

func (w *Word) RotateRight(times int) {
	length := len(w.data)
	var result [4]byte

	for i, element := range w.data {
		result[modulo(i+times, length)] = element
	}
	w.data = result
}

func modulo(num int, base int) int {
	result := num % base

	if result < 0 {
		result = base + result
	}

	return result
}
