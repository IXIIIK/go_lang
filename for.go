package main

import "fmt"

func counter() {

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

func fizzBuzz() {

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fizzBuzz()
	fmt.Println("\n")
	counter()
}
