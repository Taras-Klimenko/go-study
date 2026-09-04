package main

import (
	"fmt"
	"strconv"
)

// fizzBuzz возвращает "Fizz" для кратных 3, "Buzz" для кратных 5,
// "FizzBuzz" для кратных 15 и само число строкой в остальных случаях.
func fizzBuzz(n int) string {
	switch {
	case n%15 == 0:
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
		fmt.Println(fizzBuzz(i))
	}
}
