package keyexpand

import (
	"fmt"
	"go-aes/matr"
	"testing"
)

func TestExpandKey(t *testing.T) {
	mat := matr.NewMatrix(4, 4)

	epanded := ExpandKey(mat, 10)
	fmt.Println(epanded)
}
