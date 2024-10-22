package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"primetests/checks"
)

func generateRandomBigInt(digits int64) *big.Int {
	num, err := rand.Int(rand.Reader, new(big.Int).Exp(big.NewInt(10), big.NewInt(digits), nil))

	if err != nil {
		panic(err)
	}

	return num
}

func main() {
	for i := int64(2); i < 1000; i++ {
		n := big.NewInt(i)
		if checks.Fermat(n) != checks.Naive(n) {
			fmt.Printf("Naive and Fermat checks have different outputs for %d\n", i)
		}
	}
}
