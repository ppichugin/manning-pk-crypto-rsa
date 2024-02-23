package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const numTests = 20

var random = rand.New(rand.NewSource(time.Now().UnixNano())) //nolint:gosec

func main() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	now := time.Now()

	// Test some known primes and composites.
	testKnownValues()

	fmt.Println("elapsed: ", time.Since(now))

	// Generate random primes.
	for {
		// Get the number of digits.
		var numDigits int
		fmt.Printf("\n# Digits: ")
		fmt.Scan(&numDigits)
		if numDigits < 1 {
			break
		}

		// Calculate minimum and maximum values.
		mn := int(math.Pow(10.0, float64(numDigits-1)))
		mx := 10 * mn
		if mn == 1 {
			mn = 2
		} // 1 is not prime.

		// Find a prime.
		fmt.Printf("Prime: %d\n", findPrime(mn, mx, numTests))
	}

}

// Return a pseudo random number in the range [min, max).
func randRange(min int, max int) int {
	return min + random.Intn(max-min)
}

func fastExpMod(num, pow, mod int) int {
	if pow == 0 {
		return 1
	}
	if pow%2 == 0 {
		return fastExpMod((num*num)%mod, pow/2, mod)
	}
	return (num * fastExpMod(num, pow-1, mod)) % mod
}

// Perform tests to see if a number is (probably) prime.
func isProbablyPrime(p int, numTests int) bool {
	if p < 2 {
		return false
	}
	if p != 2 && p%2 == 0 {
		return false
	}

	// Loop for a number of times defined by numTests
	for i := 0; i < numTests; i++ {
		// Generate a random number 'a' in the range of 1 to 'p-1'
		a := randRange(1, p-1)

		// Perform a fast exponentiation modulo operation on 'a', 'p-1', and 'p'
		// If the result is not 1, return false, indicating that 'p' is not a prime number
		if fastExpMod(a, p-1, p) != 1 {
			return false
		}
	}

	// If the loop completes without finding a counterexample, return true, indicating that 'p' is probably a prime number
	return true
}

// Probabilistically find a prime number within the range [min, max).
func findPrime(min, max, numTests int) int {
	for {
		p := randRange(min, max)
		if isProbablyPrime(p, numTests) {
			return p
		}
	}
}

func testKnownValues() {
	primes := []int{
		10009, 11113, 11699, 12809, 14149,
		15643, 17107, 17881, 19301, 19793,
	}
	composites := []int{
		10323, 11397, 12212, 13503, 14599,
		16113, 17547, 17549, 18893, 19999,
	}

	// Calculate and print the probability
	probability := 1 - math.Pow(0.5, float64(numTests))
	fmt.Printf("Probability: %.6f%%\n\n", probability*100)

	// Test the primes
	fmt.Println("Primes:")
	for _, prime := range primes {
		if isProbablyPrime(prime, numTests) {
			fmt.Printf("%d  Prime\n", prime)
		} else {
			fmt.Printf("%d  Composite\n", prime)
		}
	}

	// Test the composites
	fmt.Println("\nComposites:")
	for _, composite := range composites {
		if isProbablyPrime(composite, numTests) {
			fmt.Printf("%d  Prime\n", composite)
		} else {
			fmt.Printf("%d  Composite\n", composite)
		}
	}
}
