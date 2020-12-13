package main

import (
	"fmt"
	"sort"

	"github.com/james-wallis/adventofcode/utils"
)

// DiffNumbersEqualNumber : returns true if num1 and num2 are different and they equal the result when added together
func DiffNumbersEqualNumber(num1, num2, result int) bool {
	if num1 != num2 && num1+num2 == result {
		return true
	}
	return false
}

// DoTwoNumbersInSliceMakeNumber : returns true if two different numbers in a slice equal the wanted number
func DoTwoNumbersInSliceMakeNumber(input []int, want int) bool {
	for _, num1 := range input {
		for _, num2 := range input {
			valid := DiffNumbersEqualNumber(num1, num2, want)
			if valid {
				return true
			}
		}
	}
	return false
}

// FindNumNotSumOfPreviousNumbers : returns number that isn't made up of any of the prevous X (preamble) numbers in the input slice
// defaults to -1 if not found
func FindNumNotSumOfPreviousNumbers(input []int, preamble int) int {
	for i := preamble; i < len(input); i++ {
		cutSlice := input[i-preamble : i]
		indexIsValidNum := DoTwoNumbersInSliceMakeNumber(cutSlice, input[i])
		if !indexIsValidNum {
			return input[i]
		}
	}
	return -1
}

func sumOfSlice(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// FindNumbersInSliceThatMakeNum : Find a subslice of numbers that make the target number
func FindNumbersInSliceThatMakeNum(input []int, targetNumber int) []int {
	for amountOfMems := 2; amountOfMems < len(input); amountOfMems++ {
		for i := 0; i < len(input); i++ {
			if i+amountOfMems >= len(input) {
				break
			}
			subSlice := input[i : i+amountOfMems]
			sliceTotal := sumOfSlice(subSlice)
			if sliceTotal == targetNumber {
				return subSlice
			}
		}
	}
	return nil
}

func addLowestAndHighestNumsOfSlice(nums []int) int {
	sort.Ints(nums)
	return nums[0] + nums[len(nums)-1]
}

func main() {
	fmt.Println("*** Day 9 ***")
	lines, _ := utils.ReadLinesAndConvertToInt("./input.txt")

	part1 := FindNumNotSumOfPreviousNumbers(lines, 25)
	fmt.Println("Part 1")
	fmt.Printf("First number that doesn't match rules = %d\n\n", part1)

	part2Slice := FindNumbersInSliceThatMakeNum(lines, part1)
	part2 := addLowestAndHighestNumsOfSlice(part2Slice)
	fmt.Println("Part 2")
	fmt.Printf("Smallest and largest added = %d\n\n", part2)

}
