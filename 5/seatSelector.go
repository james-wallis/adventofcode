package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/james-wallis/adventofcode/utils"
)

func determineSeatRowOrColumn(str string, lowest, highest int64) int64 {
	midpointLow := int64(math.Floor((float64(highest) + float64(lowest)) / 2))
	midpointHigh := midpointLow + 1
	if len(str) > 1 {
		nextStr := str[1:]

		if strings.HasPrefix(str, "F") || strings.HasPrefix(str, "L") {
			return determineSeatRowOrColumn(nextStr, lowest, midpointLow)
		} else if strings.HasPrefix(str, "B") || strings.HasPrefix(str, "R") {
			return determineSeatRowOrColumn(nextStr, midpointHigh, highest)
		}
	}

	if strings.HasPrefix(str, "B") || strings.HasPrefix(str, "R") {
		return midpointHigh
	}
	return midpointLow
}

// DetermineSeatLocation takes a string (like FBFBBFFRLR) and returns:
// 1. the seat row out of 128 (first 7 characters)
// 2. the seat column out of 8 (final 3 characters)
// where F means "front", B means "back"
func DetermineSeatLocation(str string) (row, column int64) {
	lowest, highestRow, highestColumn := int64(0), int64(127), int64(7)
	row, column = 0, 0
	row = determineSeatRowOrColumn(str[:7], lowest, highestRow)
	column = determineSeatRowOrColumn(str[7:], lowest, highestColumn)
	return
}

// CalculateSeatID returns the seat ID given an input string (like FBFBBFFRLR)
// Every seat also has a unique seat ID: multiply the row by 8, then add the column.
func CalculateSeatID(str string) int64 {
	row, column := DetermineSeatLocation(str)
	return (row * 8) + column
}

func part1(lines []string) (highestNumber int64) {
	highestNumber = 0
	for _, line := range lines {
		seatID := CalculateSeatID(line)
		if seatID > highestNumber {
			highestNumber = seatID
		}
	}
	return highestNumber
}

func part2(lines []string) (mySeatID int64) {
	mySeatID = 0
	var seats []int64
	for _, line := range lines {
		seatID := CalculateSeatID(line)
		seats = append(seats, seatID)
	}

	sort.Slice(seats, func(i, j int) bool { return seats[i] < seats[j] })

	for i, seat := range seats {
		if seat+1 != seats[i+1] {
			return seat + 1
		}
	}
	return 0
}

func main() {
	lines, _ := utils.ReadLines("./input.txt")

	fmt.Println("*** Part 1 ***")
	part1Result := part1(lines)
	fmt.Printf("Highest seat ID = %d\n\n", part1Result)

	fmt.Println("*** Part 2 ***")
	part2Result := part2(lines)
	fmt.Printf("My seat ID = %d\n\n", part2Result)
}
