package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/james-wallis/adventofcode/utils"
)

// FindDirectHoldersOfBag : returns a slice containing the names of bags that are direct holders of the target bag
func FindDirectHoldersOfBag(targetBag string, rules []string) []string {
	outerBagsThatContainBag := []string{}
	for _, rule := range rules {
		splitStr := strings.Split(rule, "contain")
		outerBag, innerBags := splitStr[0], splitStr[1]
		if strings.Contains(innerBags, targetBag) {
			outerBagRe := regexp.MustCompile("[^\\s]+\\s+[^\\s]+")
			sanitisedOuterBag := outerBagRe.FindString(outerBag)
			outerBagsThatContainBag = append(outerBagsThatContainBag, sanitisedOuterBag)
		}
	}
	return outerBagsThatContainBag
}

func recusivelyGetAllOutermostBags(outerBags []string, rules []string) []string {
	eventualBags := []string{}
	for _, outerBag := range outerBags {
		outerOuterBags := FindDirectHoldersOfBag(outerBag, rules)
		outerOuterOuterBags := recusivelyGetAllOutermostBags(outerOuterBags, rules)
		// Add the current outerBags into the array of eventual bag holders
		eventualBags = append(eventualBags, outerBag)
		// Add the eventual bag holders that were found through recursion
		eventualBags = append(eventualBags, outerOuterOuterBags...)
	}
	return eventualBags
}

// GetNumEventualBagHolders : returns the number of bags that will eventually hold the targetBag
func GetNumEventualBagHolders(targetBag string, rules []string) int {
	allBags := make(map[string]bool)
	outerBags := FindDirectHoldersOfBag(targetBag, rules)

	allOuterOuterBags := recusivelyGetAllOutermostBags(outerBags, rules)
	for _, bag := range allOuterOuterBags {
		allBags[bag] = true
	}

	return len(allBags)
}

func findBagsInsideTargetBag(targetBag string, rules []string) string {
	for _, rule := range rules {
		splitStr := strings.Split(rule, "contain ")
		outerBag, innerBagsSentence := splitStr[0], splitStr[1]
		if strings.Contains(outerBag, targetBag) {
			// If target bag doesn't have any bags inside, return empty string
			if innerBagsSentence == "no other bags." {
				return ""
			}
			return innerBagsSentence
		}
	}
	return ""
}

func getNameAndAmountOfInnerBag(bag string) (string, int64) {
	amountAndName := strings.Split(bag, " ")
	amount, _ := strconv.ParseInt(amountAndName[0], 10, 64)
	name := strings.Join([]string{amountAndName[1], amountAndName[2]}, " ")
	return name, amount
}

func findOuterBagInRules(targetBag string, rules []string) string {
	for _, rule := range rules {
		if strings.HasPrefix(rule, targetBag) {
			return rule
		}
	}
	return ""
}

func parseInnerBags(rule string) map[string]int64 {
	innerBags := make(map[string]int64)
	splitStr := strings.Split(rule, "contain ")
	_, innerBagsSentence := splitStr[0], splitStr[1]
	for _, bag := range strings.Split(innerBagsSentence, ", ") {
		name, amount := getNameAndAmountOfInnerBag(bag)
		innerBags[name] = amount
	}
	return innerBags
}

// CountBagsInside : returns the total number of bags that can be inside a given bag
func CountBagsInside(targetBag string, rules []string) int64 {
	targetBagsRule := findOuterBagInRules(targetBag, rules)
	if strings.HasSuffix(targetBagsRule, "no other bags.") {
		return 0
	}

	innerBags := parseInnerBags(targetBagsRule)
	total := int64(0)
	for bag, noBags := range innerBags {
		// Add 1 for the inital bag
		total += (CountBagsInside(bag, rules) + 1) * noBags
	}
	return total
}

func main() {
	fmt.Println("*** Day 7 ***")
	lines, _ := utils.ReadLines("./input.txt")

	part1 := GetNumEventualBagHolders("shiny gold", lines)
	fmt.Println("Part 1")
	fmt.Printf("Number of bags that will eventually hold a shiny gold bag = %d\n\n", part1)

	part2 := CountBagsInside("shiny gold", lines)
	fmt.Println("Part 2")
	fmt.Printf("Number of bags inside the shiny gold one = %d\n\n", part2)
}
