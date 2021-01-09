package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	expression := ""
	fmt.Println("Enter the expression")
	fmt.Scan(&expression)
	fmt.Println(ParseCalculate(expression))
}

func ParseCalculate(expression string) int {
	num1 := 0
	re := regexp.MustCompile(`(\d+)([+*/-])(\d+)`)
	out := re.FindAllStringSubmatch(expression, -1)
	b1, _ := strconv.Atoi(out[0][1])
	b2, _ := strconv.Atoi(out[0][3])
	op := out[0][2]

	switch op {
	case "+": num1 = b1 + b2
	case "-": num1 = b1 - b2
	case "/": num1 = b1 / b2
	case "*": num1 = b1 * b2
	}
	return num1
}
