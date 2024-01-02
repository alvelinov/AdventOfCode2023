package main

import (
	"strconv"
	"strings"
)

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

func Part1(formattedInput []string) int {
	var sum int = 0

	for i, game := range formattedInput {
		gameRounds := strings.Split(strings.Trim(strings.Split(game, ":")[1], " "), "; ")
		var invalidGame bool = false
		for _, round := range gameRounds {
			roundArr := strings.Split(round, ", ")
			for _, roundParts := range roundArr {
				roundPartsApart := strings.Split(roundParts, " ")
				num, _ := strconv.Atoi(roundPartsApart[0])
				switch roundPartsApart[1][0] {
				case 'r':
					if num > MAX_RED {
						invalidGame = true
					}
				case 'g':
					if num > MAX_GREEN {
						invalidGame = true
					}
				case 'b':
					if num > MAX_BLUE {
						invalidGame = true
					}
				}
				if invalidGame {
					break
				}
			}
			if invalidGame {
				break
			}
		}
		if !invalidGame {
			sum += i + 1
		}
	}

	return sum
}

func Part2(formattedInput []string) int {
	var sum int = 0

	for _, game := range formattedInput {
		var maxRed int = 0
		var maxGreen int = 0
		var maxBlue int = 0

		gameRounds := strings.Split(strings.Trim(strings.Split(game, ":")[1], " "), "; ")
		for _, round := range gameRounds {
			roundArr := strings.Split(round, ", ")
			for _, roundParts := range roundArr {
				roundPartsApart := strings.Split(roundParts, " ")
				num, _ := strconv.Atoi(roundPartsApart[0])
				switch roundPartsApart[1][0] {
				case 'r':
					if num > maxRed {
						maxRed = num
					}
				case 'g':
					if num > maxGreen {
						maxGreen = num
					}
				case 'b':
					if num > maxBlue {
						maxBlue = num
					}
				}
			}
		}
		sum += maxRed * maxGreen * maxBlue
	}

	return sum
}
