package main

import (
	"fmt"
    "os"
	"strings"
	"unicode"
	"strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func extractDigits(input string) int {
	var firstDigit rune
	var lastDigit rune
	for _, char := range input {
		if(unicode.IsDigit(char)) {
			if(firstDigit == 0) {
				firstDigit = char
			} 
			lastDigit = char
		}
	}
	var firstAndLastDigits = string(firstDigit) + string(lastDigit)
	marks, err := strconv.Atoi(firstAndLastDigits)
	check(err)
	return marks
}

func main() {
	dat, err := os.ReadFile("input.txt")
	check(err)
    fmt.Println("hello world")

	// is a variable
	input := string(dat)

	/*
		for each line we should 
		extract the first digit and last digit
		convert them to ints

		we should make an array out of it
		and finally sum all the values in the array
	*/
	inputLines := strings.Split(input, "\n")

	var sum = 0

	for _, line := range inputLines {
		digitFromLine := extractDigits(line)

		sum += digitFromLine
	}

	fmt.Println(sum)
	// fmt.Println(inputLines[0])
}