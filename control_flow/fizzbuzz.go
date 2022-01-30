package main

import "fmt"

func fizzbuzz(number int) {
	switch {
	case number%15 == 0:
		fmt.Println("FizzBuzz")
	case number%3 == 0:
		fmt.Println("Fizz")
	case number%5 == 0:
		fmt.Println("Buzz")
	default:
		fmt.Println(number)
	}
}

func main() {
	for i := 1; i <= 100; i++ {
		fizzbuzz(i)
	}
}
