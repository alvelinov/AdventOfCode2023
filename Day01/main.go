package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var filename string = os.Args[1]
	rawInput, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Missing input file!")
	}

	part, _ := strconv.Atoi(os.Args[2])

	var sum int = 0
	formattedInput := strings.Split(string(rawInput), winNewLine)
	if part == 2 {
		for _, line := range formattedInput {
			sum += 10*FindFirstDigit_2(line) + FindSecondDigit_2(line)
		}
	} else if part == 1 {
		for _, line := range formattedInput {
			sum += 10*FindFirstDigit_1(line) + FindSecondDigit_1(line)
		}
	} else {
		log.Fatal("Missing or invalid second argument - problem part (1/2)")
	}

	fmt.Println(sum)
}
