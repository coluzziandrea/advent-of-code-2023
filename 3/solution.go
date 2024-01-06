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

func extractNextNumberWithIndex(input string, index int) (int, int) {
	var number string
	var lastIndex int 
	for i := index; i < len(input); i++ {
		if(unicode.IsDigit(rune(input[i]))) {
			number += string(input[i])
			lastIndex = i
		} else {
			break
		}
	}
	marks, err := strconv.Atoi(number)
	check(err)
	return marks, lastIndex
}


func checkForSymbols(input string, start int, end int) bool {
	if(start < 0) {
		start = 0
	}
	if(end >= len(input)) {
		end = len(input) - 1
	}
	for i := start; i <= end; i++ {
		if(!unicode.IsDigit(rune(input[i])) && rune(input[i]) != '.') {
			return true
		}
	}
	return false
}

func main() {
	dat, err := os.ReadFile("input.txt")
	check(err)

	input := string(dat)

	inputLines := strings.Split(input, "\n")

	var sum = 0

	for lineIndex, line := range inputLines {
		fmt.Println("line: ", line)

		var charIndex = 0

		for charIndex < len(line) {
			// 3 cases:
			// 1. digit
			// 2. space (dot)
			// 3. symbol
			char := rune(line[charIndex])

			if(unicode.IsDigit(char)) {
				number, lastIndex := extractNextNumberWithIndex(line, charIndex)
				fmt.Println("number: ", number)
				fmt.Println("lastIndex: ", lastIndex)

				// check for symbols in this line, previous line and next line
				// for range (charIndex - 1, lastIndex + 1) 

				var hasAdjacentSymbols = false

				if(lineIndex > 0) {
					hasAdjacentSymbols = hasAdjacentSymbols || checkForSymbols(inputLines[lineIndex - 1], charIndex - 1, lastIndex + 1)
				}

				if(lineIndex < len(inputLines) - 1) {
					hasAdjacentSymbols = hasAdjacentSymbols || checkForSymbols(inputLines[lineIndex + 1], charIndex - 1, lastIndex + 1)
				}
				
				hasAdjacentSymbols = hasAdjacentSymbols || checkForSymbols(line, charIndex - 1, lastIndex + 1)

				if(hasAdjacentSymbols) {
					fmt.Println("hasAdjacentSymbols: ", number)
					sum += number
				}

				charIndex = lastIndex + 1 
			} else {
				charIndex++
			}

		}		
	}

	fmt.Println("sum: ", sum)
}
