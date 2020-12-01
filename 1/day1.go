package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// ReadLines reads a file, splits it into lines and converts the lines into ints (assumes given file only contains ints)
func ReadLines(path string) ([]int64, error) {
	content, readFileErr := ioutil.ReadFile(path)
	if readFileErr != nil {
		return nil, readFileErr
	}
	lines := strings.Split(string(content), "\n")
	var numbers []int64

	for i := 0; i < len(lines); i++ {
		convertedLine, conversionErr := strconv.ParseInt(lines[i], 10, 64)
		if conversionErr != nil {
			return nil, conversionErr
		}
		numbers = append(numbers, convertedLine)
	}

	return numbers, nil
}

func addInts(x int64, y int64) int64 {
	return x + y
}

func multiplyInts(x int64, y int64) int64 {
	return x * y
}

func intIs2020(x int64) bool {
	return x == 2020
}

func findNumberToMake2020(numToAddTo int64, numbers []int64) (result int64) {
	result = -1
	for i := 0; i < len(numbers); i++ {
		if intIs2020(addInts(numToAddTo, numbers[i])) {
			return numbers[i]
		}
	}
	return
}

// CalculateWhichTwoNumbersMake2020 given an input array, will return two numbers which when added together make 2020
func CalculateWhichTwoNumbersMake2020(numbers []int64) (int64, int64) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if intIs2020(addInts(numbers[i], numbers[j])) {
				return numbers[i], numbers[j]
			}
		}
	}
	return -1, -1
}

// CalculateWhichThreeNumbersMake2020 given an input array, will return three numbers which when added together make 2020
func CalculateWhichThreeNumbersMake2020(numbers []int64) (int64, int64, int64) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			for k := 0; k < len(numbers); k++ {
				if intIs2020(addInts(numbers[k], addInts(numbers[i], numbers[j]))) {
					return numbers[i], numbers[j], numbers[k]
				}
			}
		}
	}
	return -1, -1, -1
}
