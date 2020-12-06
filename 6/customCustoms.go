package main

import (
	"fmt"
	"strings"

	"github.com/james-wallis/adventofcode/utils"
)

// CountLettersInString : Counts the individual letters in a string, ignores new lines
func CountLettersInString(line string) int {
	m := make(map[string]bool)
	for _, l := range line {
		letter := string(l)
		_, ok := m[letter]
		if !ok && letter != "\n" {
			m[letter] = true
		}
	}
	return len(m)
}

// CountQuestionsAnswered : Counts the number of questions answered given an slice of strings
func CountQuestionsAnswered(lines []string) int {
	questionsAnswered := 0
	for _, line := range lines {
		questionsAnswered += CountLettersInString(line)
	}
	return questionsAnswered
}

func existsInAllOtherAnswers(otherAnswers []string, letter string) bool {
	for _, line := range otherAnswers {
		if !strings.Contains(line, letter) {
			return false
		}
	}
	return true
}

// CountLettersInEachLine : given a group of answers, returns the amount of answers that exist in each group member's answers
func CountLettersInEachLine(groupOfAnswers string) int {
	lines := strings.Split(groupOfAnswers, "\n")
	if len(lines) <= 1 {
		return CountLettersInString(lines[0])
	}

	// All valid answers must be in each group
	// If the answer does not exist in the first line AND the rest, then it is not a valid answer
	firstLotOfAnswers := strings.Split(lines[0], "")
	otherAnswers := lines[1:]
	questionsAnsweredByAll := 0

	for _, l := range firstLotOfAnswers {
		if existsInAllOtherAnswers(otherAnswers, string(l)) {
			questionsAnsweredByAll++
		}

	}
	return questionsAnsweredByAll
}

// CountQuestionsAnsweredByEveryone : returns the number of questions answered by everyone in a group
// I.e. For each group, count the number of questions to which everyone answered "yes". What is the sum of those counts?
func CountQuestionsAnsweredByEveryone(groups []string) int {
	questionsAnswered := 0
	for _, group := range groups {
		questionsAnswered += CountLettersInEachLine(group)
	}
	return questionsAnswered
}

func main() {
	fmt.Println("*** Day 5 ***")
	lines, _ := utils.ReadLinesWithSeparator("./input.txt", "\n\n")

	part1 := CountQuestionsAnswered(lines)
	fmt.Println("Part 1")
	fmt.Printf("Number of questions answered = %d\n\n", part1)

	part2 := CountQuestionsAnsweredByEveryone(lines)
	fmt.Println("Part 2")
	fmt.Printf("Number of questions answered by everyone = %d\n\n", part2)
}
