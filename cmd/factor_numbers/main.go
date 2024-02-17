package main

import (
	"fmt"
	"os"
	"time"
)

// Some Tests:
// Number to factor: 8876044532898802067
// findFactors:       0.488225 seconds
// 8876044532898802067
// [1500450271 5915587277]
//
// findFactorsSieve: 0.073120 seconds
// 8876044532898802067
// [1500450271 5915587277]

var primes []int

func main() {
	// Build an Euler's sieve holding numbers up to 1.6 billion.
	primes = sieveToPrimes(eulersSieve(1_600_000_000))

	var number int
	for {
		fmt.Printf("Number to factor: ")
		_, err := fmt.Scan(&number)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		checkNumber(number)

		// Find the factors the slow way.
		start := time.Now()
		factors := findFactors(number)
		elapsed := time.Since(start)
		fmt.Printf("findFactors:       %f seconds\n", elapsed.Seconds())
		fmt.Println(multiplySlice(factors))
		fmt.Println(factors)
		fmt.Println()

		// Use the Euler's sieve to find the factors.
		start = time.Now()
		factors = findFactorsSieve(number)
		elapsed = time.Since(start)
		fmt.Printf("findFactorsSieve: %f seconds\n", elapsed.Seconds())
		fmt.Println(multiplySlice(factors))
		fmt.Println(factors)
		fmt.Println()
	}
}

// checkNumber checks if the number is greater than 0. If not, it prints an error message and exits the program.
func checkNumber(n int) {
	if n < 1 {
		fmt.Println("Number must be greater than 0. Exiting...")
		os.Exit(1)
	}
}

func findFactors(num int) []int {
	var factors []int

	// While num is divisible by 2, add 2 to factors and divide num by 2
	for num%2 == 0 {
		factors = append(factors, 2)
		num /= 2
	}

	factor := 3

	// While factor * factor <= num
	for factor*factor <= num {
		// If num is divisible by factor, add factor to factors and divide num by factor
		if num%factor == 0 {
			factors = append(factors, factor)
			num /= factor
		} else {
			// If num is not divisible by factor, add 2 to factor
			factor += 2
		}
	}

	// If factor is not 1, add factor to factors
	if num > 1 {
		factors = append(factors, num)
	}

	return factors
}

func multiplySlice(slice []int) int {
	product := 1
	for _, num := range slice {
		product *= num
	}
	return product
}

func eulersSieve(max int) []bool {
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

func findFactorsSieve(num int) []int {
	var factors []int
	if num < 0 {
		factors = append(factors, -1)
		num = -num
	}

	// Pull out prime factors.
	for _, factor := range primes {
		if factor*factor > num {
			if num != 1 {
				factors = append(factors, num)
			}
			break
		}
		for num%factor == 0 {
			factors = append(factors, factor)
			num /= factor
		}
	}

	return factors
}
