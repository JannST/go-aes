package math

func Modulo(num int, base int) int {
	result := num % base

	if result < 0 {
		result = base + result
	}
	return result
}
