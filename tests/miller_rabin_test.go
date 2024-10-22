package tests

import (
	"fmt"
	"primetests/prime"
	"testing"
)

var sizes [3]int = [3]int{200, 300, 400} //digit length
var bases [3]int = [3]int{20, 30, 40}    //number of bases to try

func BenchmarkMillerRabin(b *testing.B) {
	for _, d := range sizes {
		for _, base := range bases {
			b.Run(fmt.Sprintf("%d-digits-%d-base", d, base), func(b *testing.B) {
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					prime.Generate(prime.MillerRabin, d, base)
				}
			})
		}
	}
}

func BenchmarkMillerRabinPar(b *testing.B) {
	for _, d := range sizes {
		for _, base := range bases {
			b.Run(fmt.Sprintf("%d-digits-%d-base", d, base), func(b *testing.B) {
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					prime.Generate(prime.MillerRabinPar, d, base)
				}
			})
		}
	}
}
