package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var maxN int
	for {
		fmt.Printf("Max: ")
		_, err := fmt.Scan(&maxN)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		checkNumber(maxN)

		start := time.Now()
		sieve := sieveOfEratosthenes(maxN)
		elapsed := time.Since(start)
		fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())

		if maxN <= 1000 {
			printSieve(sieve)

			primes := sieveToPrimes(sieve)
			fmt.Println(primes)
		}
	}
}

// checkNumber checks if the number is greater than 0. If not, it prints an error message and exits the program.
func checkNumber(n int) {
	if n < 1 {
		fmt.Println("Number must be greater than 0. Exiting...")
		os.Exit(1)
	}
}

func sieveOfEratosthenes(max int) []bool {
	// Create a boolean array "prime[0..n]" and initialize
	// all entries it as true. A value in prime[i] will
	// finally be false if i is Not a prime, else true.
	prime := make([]bool, max+1)
	for i := 0; i < max+1; i++ {
		prime[i] = true
	}

	for p := 2; p*p <= max; p++ {
		// If prime[p] is not changed, then it is a prime
		if prime[p] {
			// Update all multiples of p
			for i := p * p; i <= max; i += p {
				prime[i] = false
			}
		}
	}

	return prime
}

func printSieve(sieve []bool) {
	// Print 2 separately
	if len(sieve) > 2 {
		fmt.Print(2, " ")
	}

	// Loop through only the odd-numbered slice entries
	for p := 3; p < len(sieve); p += 2 {
		if sieve[p] {
			fmt.Print(p, " ")
		}
	}
	fmt.Println()
}

func sieveToPrimes(sieve []bool) []int {
	primes := make([]int, 0)

	// Add 2 separately
	if len(sieve) > 2 {
		primes = append(primes, 2)
	}

	// Loop through only the odd-numbered slice entries
	for p := 3; p < len(sieve); p += 2 {
		if sieve[p] {
			primes = append(primes, p)
		}
	}

	return primes
}
