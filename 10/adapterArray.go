package main

import (
	"fmt"
	"sort"

	"github.com/james-wallis/adventofcode/utils"
)

// CountDifferencesOf1Or3InSlice : counts how many differences or 1 and how many of 3 exist in a slice
func CountDifferencesOf1Or3InSlice(input []int) map[int]int {
	differences := map[int]int{
		1: 0,
		3: 0,
	}
	// Ensure that the input slice always starts with a 0
	nums := []int{0}
	nums = append(nums, input...)
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		diff := nums[i+1] - nums[i]
		if diff == 1 || diff == 3 {
			differences[diff]++
		}
	}

	// Add another 3 on as the built-in adapter is always +3 above the previous adapter
	differences[3]++

	return differences
}

// SplitSliceIntoSubMapsOn3Difference : Splits a slice into subslices when a difference of 3 is found
func SplitSliceIntoSubMapsOn3Difference(input []int) []map[int]bool {
	sort.Ints(input)
	split := []map[int]bool{
		make(map[int]bool),
	}
	currentSplitIndex := 0

	// Add 0 prefix
	split[0][0] = true

	for i := 0; i < len(input)-1; i++ {
		diff := input[i+1] - input[i]
		split[currentSplitIndex][input[i]] = true
		if diff == 3 {
			split = append(split, make(map[int]bool))
			currentSplitIndex++
		}
	}

	// Add final number into final slice
	split[currentSplitIndex][input[len(input)-1]] = true
	return split
}

// CountAllArrangementsOfMap : returns the tribonacci value for the length of the map (equates to potential combinations)
func CountAllArrangementsOfMap(input map[int]bool) int {
	// Seems to online increase up to 13
	tribon := []int{1, 1, 2, 4, 7, 13}
	return tribon[len(input)-1]
}

// CountAllArrangementsOfSlice : returns the total different arrangements the jolt slice can be in
func CountAllArrangementsOfSlice(input []int) int {
	total := 1
	maps := SplitSliceIntoSubMapsOn3Difference(input)
	for _, subMap := range maps {
		arrangementsOfSubMap := CountAllArrangementsOfMap(subMap)
		total *= arrangementsOfSubMap
	}
	return total
}

func main() {
	fmt.Println("*** Day 10 ***")
	lines, _ := utils.ReadLinesAndConvertToInt("./input.txt")

	part1Map := CountDifferencesOf1Or3InSlice(lines)
	part1 := part1Map[1] * part1Map[3]
	fmt.Println("Part 1")
	fmt.Printf("1 jolt multiplied by 3 jolt differences = %d\n\n", part1)

	part2 := CountAllArrangementsOfSlice(lines)
	fmt.Println("Part 2")
	fmt.Printf("total arrangements = %d\n\n", part2)
}
