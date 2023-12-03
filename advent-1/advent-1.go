package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Advent of Code 2023, Day 1, Part 1 & 2
func main() {
	// read the input file
	input, _ := os.ReadFile("input.txt")

	// function to calculate the sum of all the numbers in the input file
	calc := func(r *strings.Replacer) (result int) {
		
		// loop through each line in the input file
		for _, line := range strings.Fields(string(input)) {

			// replace all the words with numbers (e.g. "one" -> "o1e")
			line = r.Replace(line)
			// need to make second pass incase words share a letter (e.g. "eightwo" -> "e8two" -> "e8t2o")
			line = r.Replace(line)

			// find the first and last index of a number in the string
			firstIndex := strings.IndexAny(line, "0123456789")
			lastIndex := strings.LastIndexAny(line, "0123456789")

			// join the values into a string
			vals := line[firstIndex : firstIndex+1] + line[lastIndex : lastIndex+1]

			// convert the string array to an int
			val, _ := strconv.Atoi(vals)

			// add the value to the result
			result += val
		}

		return
	}

	// part 1
	fmt.Println(calc(strings.NewReplacer()))

	// part 2
	// replace all the words with numbers
	// keep the first and last letter incase they are shared between words (e.g. "eightwo" -> "e8two" -> "e8t2o")
	fmt.Println(calc(strings.NewReplacer(
		"one", "o1e",
		"two", "t2o",
		"three", "t3e",
		"four", "f4r",
		"five", "f5e",
		"six", "s6x",
		"seven", "s7n",
		"eight", "e8t",
		"nine", "n9n",
		"zero", "z0o")))
}
