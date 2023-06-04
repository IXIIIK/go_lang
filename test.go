package main

import "fmt"

func str(n string) {

	for i := 0; i < len(n); i++ {
		fmt.Printf("%c\n", n[i])
	}

}

func main() {
	str("shalom")
}
