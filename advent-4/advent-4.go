package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"
)

type GameMatches struct {
	totalMatches, totalWinnings int
	matchingCount               map[string]int
	copiesCount                 int
}

var (
	gameMatchesMap map[int]GameMatches
)

func sum(arr []int) int {
	sum := 0

	for index := range arr {
		sum += arr[index]
	}

	return sum
}

func calculateCardWinnings(input []byte) (result int, gameMatchesMap map[int]GameMatches) {
	gameMatchesMap = make(map[int]GameMatches)

	// loop through each line in the input file
	for gameIndex, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {

		// separate into lottery numbers and winning numbers
		cardNumbers := strings.Split(strings.Split(line, ":")[1], "|")
		lotteryNumbers := strings.Split(cardNumbers[0], " ")
		winningNumbers := strings.Split(cardNumbers[1], " ")

		currentGame := GameMatches{
			totalWinnings: 0,
			totalMatches:  0,
			matchingCount: make(map[string]int),
			copiesCount:   1,
		}

		// count the number of matches between the lottery numbers and the winning numbers into a map
		for _, lotteryNumber := range lotteryNumbers {
			lotteryNumber = strings.TrimSpace(lotteryNumber)

			if len(lotteryNumber) == 0 || !unicode.IsDigit(rune(lotteryNumber[0])) {
				continue
			}

			for _, winningNumber := range winningNumbers {
				winningNumber = strings.TrimSpace(winningNumber)

				if len(winningNumber) == 0 || !unicode.IsDigit(rune(winningNumber[0])) {
					continue
				}

				if lotteryNumber == winningNumber {
					currentGame.matchingCount[lotteryNumber]++
					currentGame.totalMatches++
				}
			}
		}

		winningsCount := 0
		for _, matchCount := range currentGame.matchingCount {
			winningsCount += matchCount
		}

		// part 1
		// multiply the winnings by 2 for each match
		// e.g. 2^0 = 1, 2^1 = 2, 2^2 = 4, 2^3 = 8, etc.
		if winningsCount > 0 {
			currentGame.totalWinnings = int(math.Exp2(float64(winningsCount - 1)))
		}

		gameMatchesMap[gameIndex] = currentGame
	}

	// part 1
	result = 0
	for _, gameMatches := range gameMatchesMap {
		result += gameMatches.totalWinnings
	}

	return
}

func calcCardCopies(gameMatchesMap map[int]GameMatches) (result int) {
	cardCount := make([]int, len(gameMatchesMap))

	// fill with 1 to start
	for i := range cardCount {
		cardCount[i] = 1
	}

	newCards := sum(cardCount)
	result = newCards

	// while loop until there are no more new cards
	for {
		if newCards == 0 {
			break
		}

		nextCardCount := getNextCards(cardCount, gameMatchesMap)
		newCards = sum(nextCardCount)
		result += newCards
		cardCount = nextCardCount
	}

	return
}

func getNextCards(cardCount []int, gameMatchesMap map[int]GameMatches) (nextCardCount []int) {
	nextCardCount = make([]int, len(cardCount))
	
	for index := range gameMatchesMap {		
		currentCard := gameMatchesMap[index]

		if cardCount[index] > 0 {			
			matchesCount := currentCard.totalMatches

			for matchIndex := index + 1; matchIndex < index+1+matchesCount; matchIndex++ {
				nextCardCount[matchIndex] += cardCount[index]
			}
		}
	}

	return
}

// Advent of Code 2023, Day 4, Part 1 & 2
// https://adventofcode.com/2023/day/4
func main() {

	// read the input file
	input, _ := os.ReadFile("input.txt")

	result, gameMatchesMap := calculateCardWinnings(input)
	result2 := calcCardCopies(gameMatchesMap)

	// part 1
	fmt.Println(result)

	// part 2
	fmt.Println(result2)
}
