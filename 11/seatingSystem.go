package main

import (
	"fmt"
	"reflect"

	"github.com/james-wallis/adventofcode/utils"
)

func isOccupied(char string) bool {
	return char == "#"
}

func numColOccupied(line []string, x int, isMiddleRow bool) int {
	leftCol, midCol, rightCol := x-1, x, x+1
	numOccupied := 0

	if leftCol >= 0 && isOccupied(line[leftCol]) {
		numOccupied++
	}

	if rightCol < len(line) && isOccupied(line[rightCol]) {
		numOccupied++
	}

	if !isMiddleRow && isOccupied(line[midCol]) {
		numOccupied++
	}

	return numOccupied
}

func NumOccupiedAdjacentSeats(board [][]string, x, y int) int {
	topRow, midRow, bottomRow := y-1, y, y+1
	numOccupied := 0

	if topRow >= 0 {
		numOccupied += numColOccupied(board[topRow], x, false)
	}

	if bottomRow < len(board) {
		numOccupied += numColOccupied(board[bottomRow], x, false)
	}

	numOccupied += numColOccupied(board[midRow], x, true)

	return numOccupied
}

func IterateOverBoard(board [][]string) [][]string {
	newBoard := [][]string{}

	for y, line := range board {
		newLine := []string{}
		for x, char := range line {
			characterToAppend := board[y][x]
			if char == "." {
				characterToAppend = "."
			} else {
				num := NumOccupiedAdjacentSeats(board, x, y)
				if char == "#" && num >= 4 {
					characterToAppend = "L"
				} else if char == "L" && num == 0 {
					characterToAppend = "#"
				}
			}
			newLine = append(newLine, characterToAppend)
		}
		newBoard = append(newBoard, newLine)
	}
	return newBoard
}

func countOccupiedSeats(board [][]string) int {
	occupiedSeats := 0
	for _, line := range board {
		for _, char := range line {
			if char == "#" {
				occupiedSeats++
			}
		}
	}
	return occupiedSeats
}

func IterateBoardUntilNoSeatsChange(board [][]string) int {
	currentState := board
	nextState := [][]string{}
	totalIterations := 0
	statesAreTheSame := false
	for !statesAreTheSame {
		totalIterations++
		nextState = IterateOverBoard(currentState)
		statesAreTheSame = reflect.DeepEqual(currentState, nextState)
		currentState = nextState
		nextState = nil
	}

	return countOccupiedSeats(currentState)
}

func main() {
	fmt.Println("*** Day 11 ***")
	input, _ := utils.ReadLinesAndSplitCharacters("./input.txt")

	part1 := IterateBoardUntilNoSeatsChange(input)
	fmt.Println("Part 1")
	fmt.Printf("Number of occupied seats when board has stopped changing = %d\n\n", part1)

}
