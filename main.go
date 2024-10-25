package main

import (
	"fmt"
	"primetests/prime"
)

func main() {
	n, err := prime.Generate(prime.MillerRabinPar, 1000, 20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Found prime number: ", n)
}
