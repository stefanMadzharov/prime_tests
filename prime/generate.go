package prime

import (
	"context"
	"fmt"
	"math/big"
	"primetests/checks"
	"primetests/rand"
	"runtime"
)

type PrimeAlgorithm int

const (
	Naive = iota
	Fermat
	MillerRabin
	MillerRabinPar // parallel Miller-Rabin
)

var ctx context.Context
var cancel context.CancelFunc

// Generate generates a probable prime number using the selected primality test algorithm.
// algo specifies the primality testing algorithm to use (Naive, Fermat, MillerRabin, or MillerRabinPar).
// params is a variadic parameter:
// - params[0] is the number of digits of the desired prime.
// - params[1] is required for MillerRabin and MillerRabinPar and defines the number of bases used in the test.
// For parallel Miller-Rabin (MillerRabinPar), goroutines are used for concurrent primality checking.
func Generate(algo PrimeAlgorithm, params ...int) (*big.Int, error) {
	if err := warnings(algo, params...); err != nil {
		return nil, err
	}
	digits := int64(params[0])
	goroutines := runtime.NumCPU()

	ctx, cancel = context.WithCancel(context.Background())
	resultChan := make(chan *big.Int)
	notPrimeChan := make(chan struct{})
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
		case MillerRabinPar:
			bases := params[1]

			if goroutines > 0 {
				go checks.MillerRabinPar(ctx, n, bases, cancel, notPrimeChan, resultChan)
				goroutines--
				continue
			}
			select {
			case result := <-resultChan:
				return result, nil
			case <-notPrimeChan:
				goroutines++
			}
		}
	}
}

// warnings checks if the input parameters are valid for the given primality test algorithm.
// algo specifies the chosen algorithm.
// params contains:
// - params[0] = number of digits
// - params[1] (optional) = number of bases (for MillerRabin and MillerRabinPar)
// Returns an error if parameters are missing, too many, or inappropriate for the algorithm.
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
