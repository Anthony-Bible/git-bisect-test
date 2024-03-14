package main 


import (
  "fmt"
  "os"
  "strconv"
  "strings"

)
func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) int {
	return a / b
}

func Exponent(a, b int) int{
	return a ^ b
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: calc <operation> <number1> <number2>")
		os.Exit(1)
	}

	operation := strings.ToLower(os.Args[1])
	a, err1 := strconv.Atoi(os.Args[2] )
	b, err2 := strconv.Atoi(os.Args[3])
	if err1 != nil || err2 != nil {
		fmt.Println("Error: Please provide two numerical arguments.")
		os.Exit(1)
	}

	switch operation {
	case "add":
		fmt.Println("Result:", Add(a, b))
	case "subtract":
		fmt.Println("Result:", Subtract(a, b))
	case "multiply":
		fmt.Println("Result:", Multiply(a, b))
	case "divide":
		fmt.Println("Result:", Divide(a, b))
	case "exponent":
		fmt.Println("Result:", Exponent(a, b))
	default:
		fmt.Println("Error: Unsupported operation. Use add, subtract, multiply, or divide.")
		os.Exit(1)
	}
}
