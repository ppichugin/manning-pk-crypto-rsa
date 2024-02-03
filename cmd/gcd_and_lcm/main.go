package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b int
	for {
		fmt.Print("Enter value for A: ")
		_, err := fmt.Scan(&a)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		checkNumber(a)

		fmt.Print("Enter value for B: ")
		_, err = fmt.Scan(&b)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		checkNumber(b)

		g := gcd(a, b)
		l := lcm(a, b)

		fmt.Printf("GCD(%d, %d) = %d\n", a, b, g)
		fmt.Printf("LCM(%d, %d) = %d\n", a, b, l)
	}
}

// checkNumber checks if the number is greater than 0. If not, it prints an error message and exits the program.
func checkNumber(n int) {
	if n < 1 {
		fmt.Println("Number must be greater than 0. Exiting...")
		os.Exit(1)
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
