package ranWord

import "crypto/rand"
import "math/big"

var Hello = "This is package wordz"
var Prefix = "Random word: "
var cityWord = []string{"Moscow", "Tyumen", "Praga", "Berlin", "New York"}

func City() string {
	max := len(cityWord)
	r, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return get(r.Int64())
}

func get(index int64) string {
	return Prefix + cityWord[index]
}
