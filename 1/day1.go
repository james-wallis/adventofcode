package main

import (
	"sort"
	"strconv"

	"github.com/james-wallis/adventofcode/utils"
)

// ReadLinesAndConvertToInts reads a file, splits it into lines and converts the lines into ints (assumes given file only contains ints)
func ReadLinesAndConvertToInts(path string) ([]int64, error) {
	lines, readFileErr := utils.ReadLines(path)
	if readFileErr != nil {
		return nil, readFileErr
	}

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

// CalculateWhichTwoNumbersMake2020 given an input array, will return two numbers which when added together make 2020 O(n^2)
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

// CalculateWhichThreeNumbersMake2020 given an input array, will return three numbers which when added together make 2020 O(n^3)
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

// OptimisedTwoNumbersMake2020 like CalculateWhichTwoNumbersMake2020 but O(N)
func OptimisedTwoNumbersMake2020(numbers []int64) (int64, int64) {
	sort.Slice(numbers, func(i, j int) bool { return numbers[i] < numbers[j] })
	start := 0
	end := len(numbers) - 1
	for i := 0; i < len(numbers); i++ {
		low := numbers[start]
		high := numbers[end]
		sol := low + high

		if sol == 2020 {
			return low, high
		} else if sol > 2020 {
			end--
		} else if sol < 2020 {
			start++
		}
	}
	return 0, 0
}
