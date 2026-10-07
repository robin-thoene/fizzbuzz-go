package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := 3
	result := Exec(input)
	fmt.Printf("%s", result)
}

func Exec(in int) []string {
	if in < 1 || in > 100000 {
		return []string{}
	}

	result := []string{}
	for i := 1; i <= in; i++ {
		switch {
		case i%3 == 0 && i%5 != 0:
			result = append(result, "Fizz")
		case i%3 != 0 && i%5 == 0:
			result = append(result, "Buzz")
		case i%3 == 0 && i%5 == 0:
			result = append(result, "FizzBuzz")
		default:
			result = append(result, strconv.Itoa(i))
		}
	}

	return result
}
