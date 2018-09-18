package key_schedule

import (
	"go-aes/math"
	t "go-aes/tables"
)

func ExpandKey(key []byte, nb int, nk int, rounds int) []byte {
	keys := make([]byte, nb*(rounds+1)*nb)

	copy(keys[0:len(key)], key)
	temp := make([]byte, nb)

	for i := nk; i < nb*(rounds+1); i++ {
		copy(temp, keys[(i-1)*nb:i*nb])

		if i%nk == 0 {
			RotBytes(temp, -1)
			t.SubWord(temp)
			temp[0] ^= t.Rcon(i / nk)
		} else if nk > 6 && i%nk == 4 {
			t.SubWord(temp)
		}
		r := math.Xor(keys[(i-nk)*nb:(i-nk)*nb+nb], temp)
		copy(keys[i*nb:(i+1)*nb], r)
	}
	return keys
}

func RotBytes(data []byte, steps int) {
	buffer := make([]byte, len(data))
	copy(buffer, data)

	for i := 0; i < len(data); i++ {
		data[math.Modulo(i+steps, len(data))] = buffer[i]
	}
}

/*
func (m *Matrix) RotWord(column int, times int) {
	col := key_schedule[column*m.height : column*m.height+m.height]
	buffer := make([]byte, m.height)
	copy(buffer, col)

	for i, element := range buffer {
		col[math.Modulo(i+times, m.height)] = element
	}
}

*/
