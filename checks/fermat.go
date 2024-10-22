package checks

import (
	"math/big"
)

// Prime test based on Fermat's little theorem
// 2^(n-1) % n == 1
// Bear in mind that this is a probabilistic algorithm
// https://en.wikipedia.org/wiki/Fermat%27s_little_theorem
func Fermat(n *big.Int) bool {
	one := big.NewInt(1)
	pow := new(big.Int).Exp(big.NewInt(2), new(big.Int).Sub(n, one), n)
	return pow.Cmp(one) == 0
}

// Prime test based on Fermat's little theorem with custom base a
// a^(n-1) % n == 1
// https://en.wikipedia.org/wiki/Fermat%27s_little_theorem
func FermatCustom(n, a *big.Int) bool {
	one := big.NewInt(1)
	pow := new(big.Int).Exp(a, new(big.Int).Sub(n, one), n)
	return pow.Cmp(one) == 0
}
