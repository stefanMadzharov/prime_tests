# PrimeTests

This is a small Go project I did a few months ago when I got interested in prime number generation and primality testing. Just putting it out there in case someone finds it useful.

This isn't meant to be a full-blown library or production tool — just an experiment in writing prime generators using different algorithms and messing around with concurrent processing.

## What It Does

PrimeGen lets you generate large prime numbers with a few different primality testing algorithms:

* Naive — straightforward trial division (slow, only for small numbers).
* Fermat — Fermat primality test (quick, not very reliable).
* Miller-Rabin — well-known probabilistic test.
* Miller-Rabin Parallel — same as above, but runs tests in parallel using goroutines.

## How It Works

You call the `Generate` function with:

```go
prime, err := Generate(algo, params...)
```

Where:

* `algo` is the algorithm you want to use (`Naive`, `Fermat`, `MillerRabin`, or `MillerRabinPar`).
* `params[0]` is the number of digits for the prime you're looking for.
* `params[1]` (only for Miller-Rabin algorithms) is the number of bases to test against.

Example:

```go
p, err := Generate(MillerRabinPar, 100, 10) // 100-digit prime, tested with 10 bases
```

It returns a `*big.Int` that should be prime (with high probability depending on the algorithm used), or an error if something goes wrong.

## Benchmark

To compare the performance of the regular Miller-Rabin test versus the parallel version, you can run the benchmarks with:

```bash
go test ./tests -bench=. -benchmem
```

## Notes

* The parallel version (`MillerRabinPar`) spins up goroutines to test different candidates and cancels them once a prime is found.
* There's some basic parameter validation in place to prevent crashing or wasting time.
* This is mostly about learning and fun — don’t rely on it for anything security-critical.

## Why?

Honestly just got curious about how primality tests work and wanted to see how far I could go with Go’s concurrency model. It's been a fun way to learn more about number theory and `math/big`, and to play with goroutines in a useful context.

## License

MIT, do what you want with it.
