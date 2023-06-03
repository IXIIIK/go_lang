package main

import (
	"fmt"
	"math/rand"
)

func calculator() {
	var user_answer int
	choice_operator := []string{
		"+",
		"-",
		"*",
		"/",
	}
	operator := choice_operator[rand.Intn(len(choice_operator))]
	num1 := rand.Intn(100)
	num2 := rand.Intn(50)
	println(num1, operator, num2, "= ?")
	fmt.Scan(&user_answer)

	if operator == "+" && num1+num2 == user_answer {
		fmt.Println("Correct!", "\n")
	} else if operator == "-" && num1-num2 == user_answer {
		fmt.Println("Correct!", "\n")
	} else if operator == "*" && num1*num2 == user_answer {
		fmt.Println("Correct!", "\n")
	} else if operator == "/" && num1/num2 == user_answer {
		fmt.Println("Correct!", "\n")
	} else {
		fmt.Println("Try again :(", "\n")
	}
}

func main() {
	for {
		calculator()
	}
}
