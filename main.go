package main

import "fmt"

func main() {
	var operation string
	var num1 int
	var num2 int
	var op_name string

	fmt.Print("Choose an operation ( +, -, *, : ) ")
	fmt.Scan(&operation)
	switch operation {
	case "+":
		op_name = "add to"
	case "-":
		op_name = "subtract from"
	case "*":
		op_name = "multiply"
	case ":":
		op_name = "divide"
	}

	fmt.Print("Type the number that you want to ", op_name, ": ")
	fmt.Scan(&num1)
	fmt.Print("The other number: ")
	fmt.Scan(&num2)
}
