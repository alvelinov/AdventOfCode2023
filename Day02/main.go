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
	formattedInput := strings.Split(string(rawInput), "\r\n")

	part, _ := strconv.Atoi(os.Args[2])

	var result int
	if part == 1 {
		result = Part1(formattedInput)
	} else if part == 2 {
		result = Part2(formattedInput)
	} else {
		log.Fatal("Missing or invalid second argument - problem part (1/2)")
	}

	fmt.Print(result)
}
