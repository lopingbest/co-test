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
