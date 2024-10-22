package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"primetests/checks"
)

func generateRandomNumber(digits int64) *big.Int {
	num, err := rand.Int(rand.Reader, new(big.Int).Exp(big.NewInt(10), big.NewInt(digits), nil))

	if err != nil {
		panic(err)
	}

	return num
}

func main() {
	d := int64(200)
	for i := 0; i < 10000; i++ {
		n := generateRandomNumber(d)
		if checks.MillerRabin(n, 10) {
			fmt.Println(n)
			break
		}
	}
}
