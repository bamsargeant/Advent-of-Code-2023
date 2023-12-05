package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type CoordinateKey struct {
	X, Y int
}

type CoordinateBounds struct {
	value, startX, endX, startY, endY int
}

type Gear struct {
	coordinate CoordinateKey
	numbers    []int
}

var (
	// the map of symbols and their coordinates
	symbolMap map[CoordinateKey]string

	// the map of gears and their coordinates
	gearMap map[CoordinateKey]Gear

	// slice of numbers and their coordinate bounds
	numberSlice []CoordinateBounds
)

func populateCharMap() {
	// read the input file
	input, _ := os.ReadFile("test.txt")
	symbolMap = make(map[CoordinateKey]string)

	// loop through each line in the input file
	for lineIndex, line := range strings.Fields(string(input)) {
		currentCharSlice := []string{}
		currentStartX := -1
		currentEndX := -1

		// loop through each character in the line
		for charIndex, char := range strings.TrimSpace(line) {

			if unicode.IsDigit(char) {
				currentCharSlice = append(currentCharSlice, string(char))

				if currentStartX == -1 {
					currentStartX = charIndex
				}

				// finish the number if it is at the end of the line
				if charIndex != len(line)-1 {
					continue
				}
			}

			if !unicode.IsDigit(char) && char != '.' {
				symbolMap[CoordinateKey{charIndex, lineIndex}] = string(char)
			}

			// if the number is finished, add it to the number slice
			if len(currentCharSlice) > 0 {
				currentEndX = charIndex - 1
				currentNumber, _ := strconv.Atoi(strings.Join(currentCharSlice, ""))

				numberSlice = append(
					numberSlice,
					CoordinateBounds{
						value:  currentNumber,
						startX: currentStartX - 1,
						endX:   currentEndX + 1,
						startY: lineIndex - 1,
						endY:   lineIndex + 1,
					})

				// reset the current char slice
				currentCharSlice = []string{}
				currentStartX = -1
				currentEndX = -1
			}
		}
	}
}

// Advent of Code 2023, Day 3, Part 1 & 2
// https://adventofcode.com/2023/day/3
func main() {
	populateCharMap()
	gearMap := make(map[CoordinateKey]Gear)

	calc := func() (result int, result2 int) {
		// loop through each number in the number slice
		for _, number := range numberSlice {

			numberCounted := false

			// loop through each coordinate in the symbol map
			for coordinate, symbol := range symbolMap {

				// check if the coordinate is within the bounds of the number
				if coordinate.X >= number.startX &&
					coordinate.X <= number.endX &&
					coordinate.Y >= number.startY &&
					coordinate.Y <= number.endY {

					// part 1 - only count the number once
					if !numberCounted {
						result += number.value
						numberCounted = true
					}

					// part 2 - set the symbol coordinate to find the gears
					if symbol == "*" {
						// if the coordinate is already in the gear map, append the number
						if gear, ok := gearMap[coordinate]; ok {
							gear.numbers = append(gear.numbers, number.value)
							gearMap[coordinate] = gear

						// if the coordinate is not in the gear map, add it
						} else {
							gearMap[coordinate] = Gear{
								coordinate: coordinate,
								numbers:    []int{number.value},
							}
						}
					}
				}
			}
		}

		// part 2 - if exacty 2 numbers share the same gear, multiply them
		for _, gear := range gearMap {
			if len(gear.numbers) == 2 {
				result2 += gear.numbers[0] * gear.numbers[1]
			}
		}

		return
	}

	res1, res2 := calc()

	fmt.Println(res1)
	fmt.Println(res2)
}
