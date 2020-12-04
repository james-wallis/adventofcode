package main

import (
	"fmt"

	"github.com/james-wallis/adventofcode/utils"
)

type movement struct {
	right int
	down  int
}

const inputFile = "./input.txt"

var part2Movements = []movement{
	movement{
		right: 1,
		down:  1,
	},
	// movement{
	// 	right: 3,
	// 	down:  1,
	// },
	movement{
		right: 5,
		down:  1,
	},
	movement{
		right: 7,
		down:  1,
	},
	movement{
		right: 1,
		down:  2,
	},
}

// GetIndexToRightOrWrap returns the index of the character to the right incremented by a given amount,
// if no such character exists it moves back to the start
func GetIndexToRightOrWrap(line string, currentIndex, numToIncrementBy int) int {
	newIndex := currentIndex + numToIncrementBy
	if newIndex < len(line) {
		return newIndex
	}

	wrappedIndex := newIndex - len(line)
	return wrappedIndex
}

func getStringCharacterAtIndex(line string, index int) string {
	return string(line[index])
}

func isTree(character string) bool {
	return character == "#"
}

func calculateNumberOfTreesHit(lines []string, numIncrementsRight, numIncrementsDown int) int {
	treesHit := 0
	currentIndex := 0

	// Ignore the first line
	for i := numIncrementsDown; i < len(lines); i += numIncrementsDown {
		line := lines[i]
		newIndex := GetIndexToRightOrWrap(line, currentIndex, numIncrementsRight)
		char := getStringCharacterAtIndex(line, newIndex)
		if isTree(char) {
			treesHit++
		}
		currentIndex = newIndex
	}
	return treesHit
}

func main() {
	fmt.Println("***        Day 3        ***")
	fmt.Print("*** Toboggan Trajectory ***\n\n")

	lines, readFileErr := utils.ReadLines(inputFile)
	if readFileErr != nil {
		fmt.Println(readFileErr)
		return
	}

	part1TreesHit := calculateNumberOfTreesHit(lines, 3, 1)

	fmt.Printf("Part 1: Number of trees hit = %d\n", part1TreesHit)

	part2Result := part1TreesHit
	for _, movement := range part2Movements {
		treesHit := calculateNumberOfTreesHit(lines, movement.right, movement.down)
		part2Result *= treesHit
	}

	fmt.Printf("Part 2: Result = %d\n", part2Result)
}
