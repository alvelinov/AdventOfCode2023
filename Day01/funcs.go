package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func FindFirstDigit_1(line string) int {
	re, _ := regexp.Compile(`\d`)
	digit := re.Find([]byte(line))
	intVal, _ := strconv.Atoi(string(digit))

	return intVal
}

func FindSecondDigit_1(line string) int {
	re, _ := regexp.Compile(`\d`)
	var digit []byte = nil

	length := len(line)
	i := 1
	for digit == nil {
		digit = re.Find([]byte(line[(length - i):]))
		i++
	}

	intVal, _ := strconv.Atoi(string(digit))
	return intVal
}

func FindFirstDigit_2(line string) int {
	re, _ := regexp.Compile(`one|two|three|four|five|six|seven|eight|nine|\d`)
	digit := re.Find([]byte(line))

	var intVal int
	if len(digit) > 1 {
		stringDigitMap := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}
		intVal, _ = strconv.Atoi(stringDigitMap[string(digit)])
	} else {
		intVal, _ = strconv.Atoi(string(digit))
	}
	fmt.Printf("%d ", intVal)
	return intVal
}

func FindSecondDigit_2(line string) int {
	re, _ := regexp.Compile(`one|two|three|four|five|six|seven|eight|nine|\d`)
	var digit []byte = nil
	length := len(line)
	i := 1
	for digit == nil {
		digit = re.Find([]byte(line[(length - i):]))
		i++
	}

	var intVal int
	if len(digit) > 1 {
		stringDigitMap := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}
		intVal, _ = strconv.Atoi(stringDigitMap[string(digit)])
	} else {
		intVal, _ = strconv.Atoi(string(digit))
	}
	fmt.Println(intVal)
	return intVal
}
