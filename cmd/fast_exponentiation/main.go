package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var num, pow, mod int
	for {
		fmt.Print("Enter num: ")
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		checkNumber(num)

		fmt.Print("Enter pow: ")
		_, err = fmt.Scan(&pow)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		checkNumber(num)

		fmt.Print("Enter mod: ")
		_, err = fmt.Scan(&mod)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		checkNumber(num)

		fastExpResult := fastExp(num, pow)
		fmt.Printf("fastExp(%d, %d) = %d\n", num, pow, fastExpResult)

		mathPowResult := int(math.Pow(float64(num), float64(pow)))
		fmt.Printf("fastExp with math.Pow(%d, %d) = %d\n", num, pow, mathPowResult)

		fastExpModResult := fastExpMod(num, pow, mod)
		fmt.Printf("fastExpMod(%d, %d, %d) = %d\n", num, pow, mod, fastExpModResult)

		mathPowModResult := int(math.Pow(float64(num), float64(pow))) % mod
		fmt.Printf("fastExpMod with math.Pow(%d, %d) mod %d = %d\n", num, pow, mod, mathPowModResult)
	}
}

// checkNumber checks if the number is greater than 0. If not, it prints an error message and exits the program.
func checkNumber(n int) {
	if n < 1 {
		fmt.Println("Number must be greater than 0. Exiting...")
		os.Exit(1)
	}
}

func fastExp(num, pow int) int {
	if pow == 0 {
		return 1
	}
	if pow%2 == 0 {
		return fastExp(num*num, pow/2)
	}
	return num * fastExp(num, pow-1)
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
