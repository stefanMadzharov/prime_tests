package checks

import (
	"math/big"
)

// check for n % d == 0 for every odd number from 3 to roundDown(sqrt(n))
func Naive(n *big.Int) bool {
	zero := big.NewInt(0)
	one := big.NewInt(1)
	two := big.NewInt(2)

	if n.Cmp(two) == 0 {
		return true
	}
	if n.Cmp(two) == -1 || new(big.Int).Rem(n, two).Cmp(zero) == 0 {
		return false
	}

	sqrtN := new(big.Int).Sqrt(n)
	for i := big.NewInt(3); i.Cmp(sqrtN) < 1; i.Add(i, two) {
		if new(big.Int).Mod(n, i).Cmp(one) == -1 {
			return false
		}
	}
	return true
}
