package main

import (
	"fmt"
)

func printPrimes(max int) {
	for n := 2; n < max+1; n++ {
		if n == 2 {
			fmt.Printf("%d is prime\n", n)
			continue
		}
		if n%2 == 0 {
			continue
		}
		isPrime := true
		for i := 3; i*i < n+1; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if !isPrime {
			continue
		}
		fmt.Printf("%d is prime\n", n)
	}
}

func main() {
	fmt.Println("=========================================")
	fmt.Println("Primes up to 10:")
	printPrimes(10)
	fmt.Println("=========================================")
	fmt.Println("Primes up to 20:")
	printPrimes(20)
	fmt.Println("=========================================")
	fmt.Println("Primes up to 30:")
	printPrimes(30)
	fmt.Println("=========================================")
}
