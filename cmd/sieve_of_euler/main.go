package main

import (
	"fmt"
	"os"
	"time"
)

// TEST
// Max: 1_000_000_000
// Eratosthenes' Sieve. Elapsed: 6.217840 seconds
// Euler's Sieve. Elapsed: 2.161329 seconds
//
// Max: 10_000_000_000
// Eratosthenes' Sieve. Elapsed: 84.752083 seconds
// Euler's Sieve. Elapsed: 27.153868 seconds
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
		sieveEratho := sieveOfEratosthenes(maxN)
		elapsed := time.Since(start)
		fmt.Printf("Eratosthenes' Sieve. Elapsed: %f seconds\n", elapsed.Seconds())

		start = time.Now()
		sieveEuler := eulersSieve(maxN)
		elapsed = time.Since(start)
		fmt.Printf("Euler's Sieve. Elapsed: %f seconds\n", elapsed.Seconds())

		if maxN <= 1000 {
			printSieve(sieveEratho)
			primesEratho := sieveToPrimes(sieveEratho)
			fmt.Println("Eratosthenes' Primes")
			fmt.Println(primesEratho)

			printSieve(sieveEuler)
			primesEuler := sieveToPrimes(sieveEuler)
			fmt.Println("Euler's Primes")
			fmt.Println(primesEuler)
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

func eulersSieve(max int) []bool {
	// Create a boolean array "prime[0..n]" and initialize
	// all entries it as true. A value in prime[i] will
	// finally be false if i is Not a prime, else true.
	prime := make([]bool, max+1)
	for i := 0; i < max+1; i++ {
		prime[i] = true
	}

	// Loop over the odd numbers between 3 and max
	for p := 3; p <= max; p += 2 {
		// If prime[p] is not changed, then it is a prime
		if prime[p] {
			// Calculate the largest odd integer that is less than or equal to max / p
			maxQ := max / p
			if maxQ%2 == 0 {
				maxQ--
			}

			// Loop from maxQ down to p
			for q := maxQ; q >= p; q -= 2 {
				// If q is marked as prime in the table, then cross out the entry for q * p
				if prime[q] {
					prime[q*p] = false
				}
			}
		}
	}

	return prime
}
