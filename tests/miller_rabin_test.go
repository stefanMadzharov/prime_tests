package tests

import (
	"fmt"
	"primetests/prime"
	"testing"
)

func BenchmarkMillerRabinSizes(b *testing.B) {
	sizes := []int{20, 50, 100} // digit length
	bases := []int{10, 15, 20}  // number of bases to try
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
