package prime

import (
	"fmt"
	"math/big"
	"primetests/checks"
	"primetests/rand"
)

type PrimeAlgorithm int

const (
	Naive = iota
	Fermat
	MillerRabin
)

func Generate(algo PrimeAlgorithm, params ...int) (*big.Int, error) {
	if err := warnings(algo, params...); err != nil {
		return nil, err
	}
	digits := int64(params[0])
	for {
		n := rand.Number(digits)
		switch algo {
		case Naive:
			if checks.Naive(n) {
				return n, nil
			}
		case Fermat:
			if checks.Fermat(n) {
				return n, nil
			}
		case MillerRabin:
			bases := params[1]
			if checks.MillerRabin(n, bases) {
				return n, nil
			}
		}
	}
}

func warnings(algo PrimeAlgorithm, params ...int) error {
	if len(params) < 1 {
		return fmt.Errorf("missing parameters")
	}
	switch algo {
	case Naive:
		if len(params) > 1 {
			return fmt.Errorf("too much parameters")
		}
		if params[0] > 7 {
			fmt.Println("There are too much digits for the 'naive' algorithm. This is going to be slow and can even crash the program")
		}
	case Fermat:
		if len(params) > 1 {
			return fmt.Errorf("too much parameters")
		}
		if params[0] > 100000 {
			fmt.Println("There are too much digits for Fermat's algorithm. This is going to be slow and can even crash the program")
		}
	case MillerRabin:
		if len(params) < 2 {
			return fmt.Errorf("too few parameters")
		}
		if len(params) > 2 {
			return fmt.Errorf("too much parameters")
		}
		if params[0] > 100000 {
			fmt.Println("There are too much digits for the Miller-Rabins' algorithm. This is going to be slow and can even crash the program")
		}
	}
	return nil
}
