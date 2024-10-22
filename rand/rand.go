package rand

import (
	"crypto/rand"
	"math/big"
)

// return random big.Int with number of digits
func Number(digits int64) *big.Int {
	num, err := rand.Int(rand.Reader, new(big.Int).Exp(big.NewInt(10), big.NewInt(digits), nil))

	if err != nil {
		panic(err)
	}

	return num
}

// return random big.Int < max
func Base(max *big.Int) *big.Int {
	two := big.NewInt(2)
	if max.Cmp(two) == 0 {
		return two
	}
	num, err := rand.Int(rand.Reader, new(big.Int).Sub(max, two))

	if err != nil {
		panic(err)
	}
	return num.Add(num, two)
}
