package main

import "fmt"

func main() {
	fmt.Println("Hello, CoPilot!")
}

//fizzbuzz
func fizzbuzz(n int) string {
	if n%15 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	} else {
		return fmt.Sprintf("%d", n)
	}
}

// fibonacci
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

//factorial
func factorial(n int) int {
	if n < 2 {
		return n
	}
	return n * factorial(n-1)
}

// isPrime
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
