package main

import (
	"fmt"
	"math/rand"
	"time"
)

const numTests = 20

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Pick two random primes p and q between 10,000 and 50,000
	p := findPrime(10000, 50000, numTests)
	q := findPrime(10000, 50000, numTests)

	// Calculate the public key modulus n = p * q
	n := p * q

	// Calculate Carmichael’s totient λn for n
	λn := totient(p, q)

	// Pick a random public key exponent e in the range [3, λn) where gcd(e, λn) = 1
	e := randomExponent(λn)

	// Find the inverse of e in the modulus λn
	d := inverseMod(e, λn)

	// Print out all the important values
	fmt.Println("*** Public ***")
	fmt.Println("Public key modulus:   ", n)
	fmt.Println("Public key exponent e:", e)

	fmt.Println("\n*** Private ***")
	fmt.Println("Primes:   ", p, ",", q)
	fmt.Println("λ(n):     ", λn)
	fmt.Println("d:        ", d)

	// Enter a loop where it prompts the user for a message
	for {
		var m int
		fmt.Print("\nMessage: ")
		fmt.Scan(&m)

		// If the user enters a message value less than 1, break out of the loop to end the program
		if m < 1 {
			break
		}

		// Encrypt the message number m
		ciphertext := fastExpMod(m, e, n)
		fmt.Println("Ciphertext:", ciphertext)

		// Decrypt the ciphertext
		plaintext := fastExpMod(ciphertext, d, n)
		fmt.Println("Plaintext: ", plaintext)
	}

}

// gcd returns the greatest common divisor of a and b, using the Euclidean algorithm.
func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	// The algorithm requires a >= b, so we swap them if necessary.
	if a < b {
		a, b = b, a
	}

	// The algorithm requires b != 0.
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// lcm returns the least common multiple of a and b.
func lcm(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	return a * b / gcd(a, b)
}

// Calculate the totient function λ(n)
// where n = p * q and p and q are prime.
func totient(p, q int) int {
	return lcm(p-1, q-1)
}

// Return a pseudo random number in the range [min, max).
func randRange(min int, max int) int {
	return min + random.Intn(max-min)
}

// Pick a random exponent e in the range (2, λn)
// such that gcd(e, λn) = 1.
func randomExponent(λn int) int {
	for {
		e := randRange(2, λn)
		if gcd(e, λn) == 1 {
			return e
		}
	}
}

// Extended Euclidean Algorithm
func extendedGCD(a, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	}
	g, x, y := extendedGCD(b%a, a)
	return g, y - (b/a)*x, x
}

// Function to find the modular inverse of a number
func inverseMod(a, m int) int {
	_, x, _ := extendedGCD(a, m)
	// m is added to handle negative x
	return (x%m + m) % m
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
