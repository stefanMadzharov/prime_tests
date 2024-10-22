package checks

import (
	"context"
	"math/big"
	"primetests/rand"
)

// n is the primality tested number
// b is the number of bases the number is going to try
// https://en.wikipedia.org/wiki/Miller%E2%80%93Rabin_primality_test
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

	for range b {
		a := rand.Base(new(big.Int).Sub(n, one))
		if !checkBase(n, a) {
			return false
		}
	}
	return true
}

// MillerRabinPar performs a parallel Miller–Rabin primality test.
// ctx is used to cancel other goroutines when a prime is found.
// n is the number to test.
// b is the number of random bases used.
// cancel should be called when a prime is confirmed.
// notPrimeChan is used to signal that n is not prime.
// resultChan is used to return n if it's probably prime.
// https://en.wikipedia.org/wiki/Miller%E2%80%93Rabin_primality_test
func MillerRabinPar(ctx context.Context, n *big.Int, b int, cancel context.CancelFunc, notPrimeChan chan<- struct{}, resultChan chan<- *big.Int) {
	zero := big.NewInt(0)
	one := big.NewInt(1)
	two := big.NewInt(2)
	if n.Cmp(two) == 0 {
		notPrimeChan <- struct{}{}
		return
	}
	if n.Cmp(two) == -1 || new(big.Int).Rem(n, two).Cmp(zero) == 0 {
		notPrimeChan <- struct{}{}
		return
	}

	for range b {
		select {
		case <-ctx.Done():
			return
		default:
			a := rand.Base(new(big.Int).Sub(n, one))
			if !checkBase(n, a) {
				notPrimeChan <- struct{}{}
				return
			}
		}
	}
	cancel()
	resultChan <- n
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
