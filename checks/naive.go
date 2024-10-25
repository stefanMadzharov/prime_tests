package checks

import (
	"math/big"
)

func Naive(n *big.Int) bool {
	one := big.NewInt(1)
	two := big.NewInt(2)

	// check if even
	if new(big.Int).Mod(n, two).Cmp(one) == -1 {
		return false
	}
	// check every odd number up to sqrt of if n%i == 0
	sqrtN := new(big.Int).Sqrt(n)
	for i := big.NewInt(3); i.Cmp(sqrtN) < 1; i.Add(i, two) {
		if new(big.Int).Mod(n, i).Cmp(one) == -1 {
			return false
		}
	}
	return true
}
