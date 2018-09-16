package basic

func AddRoundKey(state []byte, key []byte) {
	if len(state) != len(key) {
		panic("inputs are not of same size")
	}

	for i := 0; i < len(state); i++ {
		state[i] = state[i] ^ key[i]
	}
}
