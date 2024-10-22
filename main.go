package main

import (
	"fmt"
	"crypto/rand"
	"math/big"
	"math"
)

func generateRandomBigInt(digits int64) big.Int {
	num, err := rand.Int(rand.Reader, new(big.Int).Exp(big.NewInt(10), big.NewInt(digits), nil))

	if err != nil {
		panic(err)
	}

	fmt.Printf("Random number with up to %d digits:\n%s\n", digits, num.String())
	return *num
}

func main() {
	for i:=0;i<5;i++ {
		generateRandomBigInt(int64(math.Pow(float64(10), float64(i))))
	}
}
