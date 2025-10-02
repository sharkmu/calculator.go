package main

import "fmt"

func main() {
	var operation string
	var num1 float64
	var num2 float64
	var opName string

	fmt.Print("Choose an operation ( +, -, *, : ) ")
	fmt.Scan(&operation)
	switch operation {
	case "+":
		opName = "add to"
	case "-":
		opName = "subtract from"
	case "*":
		opName = "multiply"
	case ":":
		opName = "divide"
	}

	fmt.Print("Type the number that you want to ", opName, ": ")
	fmt.Scan(&num1)
	fmt.Print("The other number: ")
	fmt.Scan(&num2)

	if (operation == ":" || operation == "/") && num2 == 0 {
		fmt.Println("You can't divide with zero (0).")
	} else {
		fmt.Println("After thorough calculations, the result is: ", calc(operation, num1, num2))
	}
}

func calc(op string, n1 float64, n2 float64) (result float64) {
	switch op {
	case "+":
		result = n1 + n2
	case "-":
		result = n1 - n2
	case "*":
		result = n1 * n2
	case ":", "/":
		if n2 != 0 {
			result = n1 / n2
		}
	}

	return
}
