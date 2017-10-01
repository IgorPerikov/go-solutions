package main

import (
	"strconv"
	"fmt"
)

func main() {
	fmt.Print(fizzBuzz(15))
}

func fizzBuzz(n int) []string {
	result := make([]string, n)
	for i := 0; i < n; i++ {
		number := i + 1
		switch {
		case number%3 == 0 && number%5 == 0:
			result[i] = "FizzBuzz"
		case number%3 == 0:
			result[i] = "Fizz"
		case number%5 == 0:
			result[i] = "Buzz"
		default:
			result[i] = strconv.Itoa(number)
		}
	}
	return result
}
