package main

import (
	"crypto/rand"
	"fmt"
	// "math"
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
	for i := int64(2); i < 86; i++ {
		if checks.Fermat(big.NewInt(i)) {
			fmt.Printf("%d is a prime number\n", i)
		}
	}
}
