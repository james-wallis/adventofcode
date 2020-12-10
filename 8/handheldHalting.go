package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/james-wallis/adventofcode/utils"
)

// ParseOperation : take a string operation and returns it in a more helpful format
func ParseOperation(line string) (lineIncrement, accIncrement int, op string) {
	lineIncrement, accIncrement = 0, 0
	strs := strings.Split(line, " ")
	op = strs[0]
	val, _ := strconv.Atoi(strs[1])
	switch op {
	case "acc":
		lineIncrement = 1
		accIncrement = val
		break
	case "jmp":
		lineIncrement = val
		break
	default:
		lineIncrement = 1
	}
	return
}

// DetermineAccWhenInstructionExecutedSecondTime : given a series of instructions that eventually loop, returns the accumulator value on the first loop
func DetermineAccWhenInstructionExecutedSecondTime(lines []string) int {
	linesHit := make(map[int]bool)
	accumulator := 0

	lineNumber := 0
	for lineNumber < len(lines) {
		_, lineAlreadyHit := linesHit[lineNumber]
		if lineAlreadyHit {
			break
		}

		linesHit[lineNumber] = true
		line := lines[lineNumber]
		lineIncrement, accIncrement, _ := ParseOperation(line)
		lineNumber += lineIncrement
		accumulator += accIncrement
	}
	return accumulator
}

type state struct {
	lineNumber  int
	accumulator int
	op          string
}

func flipOperationJmpNop(line string) string {
	strs := strings.Split(line, " ")

	if strs[0] == "jmp" {
		strs[0] = "nop"
	} else {
		strs[0] = "jmp "
	}
	return strings.Join(strs, " ")
}

// ChangeSingleNopOrJmpToFixInfiniteLoop : runs a series of commands and when an infinite loop is detected, attempts to resolve it and finish the program
// returns the value of the accumulator when the program exists
func ChangeSingleNopOrJmpToFixInfiniteLoop(lines []string) int {
	linesCopy := lines
	linesHit := make(map[int]bool)
	executionOrder := []state{}
	accumulator := 0
	lineNumber := 0
	attemptedToFixOnce := false
	alreadyTriedLineChanges := make(map[int]bool)

	for lineNumber < len(lines) {
		lineAlreadyHit, lineEntryExists := linesHit[lineNumber]

		// If the line has already been hit we have hit infinite loop stage, attempt a fix here
		if lineEntryExists && lineAlreadyHit {
			i := len(executionOrder) - 1
			for i > 0 {
				linesHit[executionOrder[i].lineNumber] = false

				// Kept track of previous attempts to fix the infinite loop, they didn't work so skip
				_, alreadyTriedToChangeThisLine := alreadyTriedLineChanges[executionOrder[i].lineNumber]

				if executionOrder[i].op != "acc" && i <= len(executionOrder)-2 && !alreadyTriedToChangeThisLine {
					lineNumber = executionOrder[i].lineNumber
					accumulator = executionOrder[i].accumulator

					// Keep track of this attempted fix to the program so that we can ignore it if it fails
					alreadyTriedLineChanges[lineNumber] = true
					break
				}
				i--
			}

			// reset lines to remove previous modifications
			lines = linesCopy
			lines[lineNumber] = flipOperationJmpNop(lines[lineNumber])

			attemptedToFixOnce = true
		} else {
			linesHit[lineNumber] = true
			line := lines[lineNumber]
			lineIncrement, accIncrement, op := ParseOperation(line)
			if !attemptedToFixOnce {
				// While we haven't attempted to fix the order, keep track of the execution order of statements so we can revert
				// After we've attempted a fix we don't care as the final fix will never be further than the current attempted fix
				executionOrder = append(executionOrder, state{lineNumber, accumulator, op})
			}
			lineNumber += lineIncrement
			accumulator += accIncrement
		}
	}
	return accumulator
}

func main() {
	fmt.Println("*** Day 7 ***")
	lines, _ := utils.ReadLines("./input.txt")

	part1 := DetermineAccWhenInstructionExecutedSecondTime(lines)
	fmt.Println("Part 1")
	fmt.Printf("Value of accumulator before instruction is repeated = %d\n\n", part1)

	part2 := ChangeSingleNopOrJmpToFixInfiniteLoop(lines)
	fmt.Println("Part 2")
	fmt.Printf("Value of accumulator after instruction is fixed = %d\n\n", part2)
}
