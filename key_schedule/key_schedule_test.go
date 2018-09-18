package key_schedule

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
func TestMatrix_RotWord(t *testing.T) {
	mat := newMatrix(4, 4)
	mat.setData([]byte("1234abcd5678efgh"))
	mat.RotWord(0, 1)
	mat.RotWord(3, -2)
	mat.RotWord(1, 4)
	mat.RotWord(2, -1)

	result := []byte("4123abcd6785ghef")
	assert.Equal(t, result, mat.data)
}
*/

func TestRotWord(t *testing.T) {
	test := []byte("12345")
	RotBytes(test, -1)
	assert.Equal(t, []byte("23451"), test)

	test = []byte("12345")
	RotBytes(test, 5)
	assert.Equal(t, []byte("12345"), test)

	test = []byte("12345")
	RotBytes(test, 2)
	assert.Equal(t, []byte("45123"), test)

	test = []byte("12345")
	RotBytes(test, 4)
	assert.Equal(t, []byte("23451"), test)
}

func TestExpand128BitKey(t *testing.T) {
	result, _ := hex.DecodeString("00000000000000000000000000000000626363636263636362636363626363639b9898c9f9fbfbaa9b9898c9f9fbfbaa90973450696ccffaf2f457330b0fac99ee06da7b876a1581759e42b27e91ee2b7f2e2b88f8443e098dda7cbbf34b9290ec614b851425758c99ff09376ab49ba7217517873550620bacaf6b3cc61bf09b0ef903333ba9613897060a04511dfa9fb1d4d8e28a7db9da1d7bb3de4c664941b4ef5bcb3e92e21123e951cf6f8f188e")
	key, _ := hex.DecodeString("00000000000000000000000000000000")
	keys := ExpandKey(key, 4, 4, 10)
	assert.Equal(t, result, keys)
}

func TestExpand192BitKey(t *testing.T) {
	result, _ := hex.DecodeString("000102030405060708090a0b0c0d0e0f10111213141516175846f2f95c43f4fe544afef55847f0fa4856e2e95c43f4fe40f949b31cbabd4d48f043b810b7b34258e151ab04a2a5557effb5416245080c2ab54bb43a02f8f662e3a95d66410c08f501857297448d7ebdf1c6ca87f33e3ce510976183519b6934157c9ea351f1e01ea0372a995309167c439e77ff12051edd7e0e887e2fff68608fc842f9dcc154859f5f237a8d5a3dc0c02952beefd63ade601e7827bcdf2ca223800fd8aeda32a4970a331a78dc09c418c271e3a41d5d")
	key, _ := hex.DecodeString("000102030405060708090a0b0c0d0e0f1011121314151617")
	keys := ExpandKey(key, 4, 6, 12)
	assert.Equal(t, result, keys)
}

func TestExpand256BitKey(t *testing.T) {
	result, _ := hex.DecodeString("000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1fa573c29fa176c498a97fce93a572c09c1651a8cd0244beda1a5da4c10640badeae87dff00ff11b68a68ed5fb03fc15676de1f1486fa54f9275f8eb5373b8518dc656827fc9a799176f294cec6cd5598b3de23a75524775e727bf9eb45407cf390bdc905fc27b0948ad5245a4c1871c2f45f5a66017b2d387300d4d33640a820a7ccff71cbeb4fe5413e6bbf0d261a7dff01afafee7a82979d7a5644ab3afe6402541fe719bf500258813bbd55a721c0a4e5a6699a9f24fe07e572baacdf8cdea24fc79ccbf0979e9371ac23c6d68de36")
	key, _ := hex.DecodeString("000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f")
	keys := ExpandKey(key, 4, 8, 14)
	assert.Equal(t, result, keys)
}
