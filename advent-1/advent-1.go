package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Advent1 is the main function for the first advent of code challenge
func main() {
	// read the input file
	input, _ := os.ReadFile("input.txt")
	
	calc := func(r *strings.Replacer) (result int) {
		for _, line := range strings.Fields(string(input)) {
			
			line = r.Replace(line)
			line = r.Replace(line) // need to make second pass incase replacement takes the start/end of another word (e.g. "one" -> "o1e")
			
			firstIndex := strings.IndexAny(line, "0123456789")
			lastIndex := strings.LastIndexAny(line, "0123456789")

			vals := [2]string{
				line[firstIndex : firstIndex+1],
				line[lastIndex : lastIndex+1],
			}

			val, _ := strconv.Atoi(strings.Join(vals[:], ""))

			result += val
		}

		return
	}

	fmt.Println(calc(strings.NewReplacer()))
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