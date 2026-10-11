package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a float64
	for {
		var astring string
		fmt.Print("Input left operand ")
		fmt.Scan(&astring)

		val, err := strconv.ParseFloat(astring, 64)
		if err == nil {
			a = val
			break
		}
		fmt.Println("Invalid input")
	}

	var operation string
	for {
		fmt.Print("Input operation ")
		fmt.Scan(&operation)

		if operation == "+" || operation == "-" || operation == "*" || operation == "/" {
			break
		}
		fmt.Println("Invalid input")
	}

	var b float64
	for {
		var bstring string
		fmt.Print("Input right operand ")
		fmt.Scan(&bstring)

		val, err := strconv.ParseFloat(bstring, 64)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		if operation == "/" && val == 0 {
			fmt.Println("You can't divide by 0")
			continue
		}
		b = val
		break
	}
	var res float64
	switch operation {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	}

	if operation == "/" {
		fmt.Printf("%.3f", res)
	} else {
		fmt.Printf("%g", res)
	}
}
