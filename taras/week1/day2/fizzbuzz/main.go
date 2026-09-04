package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) string {

	switch {
	case n%3 == 0 && n%5 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(n)
	}

}

func main() {

	for i := 1; i <= 100; i++ {
		fmt.Printf("Число: %d, результат: %s \n", i, fizzBuzz(i))
	}
}
