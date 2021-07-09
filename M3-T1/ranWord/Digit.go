package ranWord

import "crypto/rand"
import "math/big"

//var Hello = "This is package wordz"

var digitWord = []string{"One", "Two", "Three", "Four", "Five", "Six"}

func Digit() string {
	max := len(cityWord)
	r, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return getCity(r.Int64())
}

func getCity(index int64) string {
	return Prefix + digitWord[index]
}
