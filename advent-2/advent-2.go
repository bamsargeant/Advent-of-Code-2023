package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Advent of Code 2023, Day 2, Part 1 & 2
// https://adventofcode.com/2023/day/2
func main() {
	// read the input file
	input, _ := os.ReadFile("input.txt")

	redLimit := 12
	greenLimit := 13
	blueLimit := 14

	// find each number and colour name
	regexSearch := regexp.MustCompile(`(\d+) (\w+)`)

	calc := func() (result int, result2 int) {
		// loop through each line in the input file
		for index, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {

			// insert the coloured blocks into a map
			colourMap := make(map[string]int, 3)

			// get the number and colour name of each block
			for _, block := range regexSearch.FindAllStringSubmatch(line, -1) {
				// convert the number to an int
				num, _ := strconv.Atoi(block[1])

				// add the number to the colour map
				// if the colour already exists, use the max value
				if val, ok := colourMap[block[2]]; ok {
					colourMap[block[2]] = int(math.Max(float64(val), float64(num)))
				} else {
					colourMap[block[2]] = num
				}
			}

			// calculate the power of the colours
			power := colourMap["red"] * colourMap["green"] * colourMap["blue"]
			
			result2 += power

			// check if the colour map is not over the limit
			if colourMap["red"] > redLimit ||
				colourMap["green"] > greenLimit ||
				colourMap["blue"] > blueLimit {
				continue
			}

			result += index + 1
		}

		return
	}

	result, result2 := calc()

	// part 1
	fmt.Println(result)

	// part 2
	fmt.Println(result2)
}
