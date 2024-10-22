package checks

import (
	"crypto/rand"
	"math/big"
)

// n is the primality tested number
// b is the number of bases the number is going to try
func MillerRabin(n *big.Int, b int) bool {
	zero := big.NewInt(0)
	one := big.NewInt(1)
	two := big.NewInt(2)
	if n.Cmp(two) == 0 {
		return true
	}
	if n.Cmp(two) == -1 || new(big.Int).Rem(n, two).Cmp(zero) == 0 {
		return false
	}

	for i := 0; i < b; i++ {
		a := generateRandomBase(new(big.Int).Sub(n, one))
		if !checkBase(n, a) {
			return false
		}
	}
	return true
}

func checkBase(n, a *big.Int) bool {
	remainders := make([]*big.Int, 0)
	zero := big.NewInt(0)
	one := big.NewInt(1)

	// halve n-1 until its odd
	pow := new(big.Int).Sub(n, one)
	for new(big.Int).Rem(pow, big.NewInt(2)).Cmp(zero) == 0 {
		remainder := new(big.Int).Exp(a, pow, n) // a^(n-1) % n
		remainders = append(remainders, remainder)
		pow.Div(pow, big.NewInt(2))
	}
	return checkCondition1(remainders) && checkCondition2(n, remainders)
}

func checkCondition1(remainders []*big.Int) bool {
	return remainders[0].Cmp(big.NewInt(1)) == 0
}

func checkCondition2(n *big.Int, remainders []*big.Int) bool {
	one := big.NewInt(1)
	nMinusOne := new(big.Int).Sub(n, one)
	for i := 0; i < len(remainders)-1; i++ {
		remainder := remainders[i]
		if remainder.Cmp(one) == 0 {
			next := remainders[i+1]
			if next.Cmp(one) != 0 && next.Cmp(nMinusOne) != 0 {
				return false
			}
		}
	}
	return true
}

func generateRandomBase(max *big.Int) *big.Int {
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
