package main

import (
	"fmt"
   	"os"
	"strings"
	"strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}



func extractGame(input string) int {
	gameX := strings.Split(input, ":")[0] // Game 1, Game 2, Game 3, etc.
	gameX = strings.Split(gameX, " ")[1] // 1, 2, 3, etc.
	number, err := strconv.Atoi(gameX)
	check(err)
	return number 
}


func extractCubeOfColor(input string, color string) int {
	var sum = 0

	cubesPerColor := strings.Split(input, ",") // 14 red  2 blue  1 green

	for _, cubePerColor := range cubesPerColor {
		cubeAndColorSplitted := strings.Split(cubePerColor, " ")
		if(strings.Contains(cubeAndColorSplitted[2], color)) {
			marks, err := strconv.Atoi(cubeAndColorSplitted[1])
			check(err)
			sum += marks
		}
	}
		
	
	return sum
}


func setIsOverloaded(setInput string) bool {
	maxRedCubes := 12
	maxGreenCubes := 13
	maxBlueCubes := 14

	redCubes := extractCubeOfColor(setInput, "red")
	greenCubes := extractCubeOfColor(setInput, "green")
	blueCubes := extractCubeOfColor(setInput, "blue")


	fmt.Println("red cubes: ", redCubes)
	fmt.Println("greenCubes: ", greenCubes)
	fmt.Println("blueCubes: ", blueCubes)

	return redCubes > maxRedCubes || greenCubes > maxGreenCubes || blueCubes > maxBlueCubes	
}

func main() {
	dat, err := os.ReadFile("input.txt")
	check(err)

	input := string(dat)

	inputLines := strings.Split(input, "\n")


	var sum = 0

	for _, line := range inputLines {
		fmt.Println(line)

		gameX := extractGame(line)

		fmt.Println("game: ", gameX)

		sets := strings.Split(line, ":")[1] 
		splittedSets := strings.Split(sets, ";")

		var gameIsOk = true

		for _, splittedSet := range splittedSets {
			fmt.Println("splittedSet: ", splittedSet)
			if(setIsOverloaded(splittedSet)) {
				gameIsOk = false
			}			
		}
		
		if(gameIsOk) {
			sum += gameX
		}
		
	}

	fmt.Println(sum)
}