package main

import "fmt"

func celcius() {
	fmt.Println("Enter a fahrenheit: ")
	var input int
	fmt.Scan(&input)

	output := (input - 32) * 5 / 9

	fmt.Println("number of degrees Celsius:", output, "\n")
}

func feet_to_meters() {
	fmt.Println("Enter a feet:")
	var input float64
	fmt.Scan(&input)
	output := input * 0.3048
	fmt.Println("number of meters:", output)
}

func main() {
	celcius()
	feet_to_meters()
}
