package main

import "fmt"

func main() {
	var name string

	fmt.Println("enter youre name: ")
	fmt.Scanln(&name)
	fmt.Println("hello, my name is", name)
}