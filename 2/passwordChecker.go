package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const inputFile = "./input.txt"

// ReadLines reads a file, splits it into lines returns an array of strings
func ReadLines(path string) ([]string, error) {
	content, readFileErr := ioutil.ReadFile(path)
	if readFileErr != nil {
		return nil, readFileErr
	}
	lines := strings.Split(string(content), "\n")

	return lines, nil
}

func getRangeFromString(str string) (low, high int64, err error) {
	splitStr := strings.Split(str, "-")
	low, high, err = 0, 0, nil
	if len(splitStr) != 2 {
		err = errors.New("Unexpected number of fields returned")
		return
	}

	low, err = strconv.ParseInt(splitStr[0], 10, 64)
	if err != nil {
		return
	}

	high, err = strconv.ParseInt(splitStr[1], 10, 64)
	if err != nil {
		return
	}

	return
}

func removeColonFromEndOfString(str string) string {
	return strings.TrimRight(str, ":")
}

func formatLine(line string) (low, high int64, characterWithoutColon, password string, err error) {
	splitLine := strings.Fields(line)
	low, high, characterWithoutColon, password, err = 0, 0, "", "", nil
	if len(splitLine) != 3 {
		err = errors.New("Unexpected number of fields returned")
		return
	}

	low, high, err = getRangeFromString(splitLine[0])
	if err != nil {
		return
	}

	characterWithoutColon = removeColonFromEndOfString(splitLine[1])

	password = splitLine[2]
	return
}

func method1(low, high int64, character, password string) bool {
	// For method one, the password matches if there exist a number of the given character in the password
	// Which is less or equal to the given low number and higher or equal to the given high number
	// For example 1-3 a: abcde as the character a is at postion 1
	characterRegex := regexp.MustCompile(character)
	matches := characterRegex.FindAllStringIndex(password, -1)
	numberOfMatches := int64(len(matches))
	if numberOfMatches >= low && numberOfMatches <= high {
		return true
	}
	return false
}

func method2(low, high int64, character, password string) bool {
	// For method two, the password matches if the given character exists at one of the given locations but not both
	// For example:
	// 		1-3 a: abcde is valid: position 1 contains a and position 3 does not.
	// 		1-3 b: cdefg is invalid: neither position 1 nor position 3 contains b.
	//		2-9 c: ccccccccc is invalid: both position 2 and position 9 contain c.
	lowChar := string(password[low-1])
	highChar := string(password[high-1])
	if (lowChar == character && highChar != character) || (lowChar != character && highChar == character) {
		return true
	}
	return false
}

func calculateValidPasswords(lines []string, method int) (int, error) {
	validPasswords := 0
	for _, line := range lines {
		low, high, character, password, formatErr := formatLine(line)
		if formatErr != nil {
			return 0, formatErr
		}

		if method == 1 && method1(low, high, character, password) {
			validPasswords++
		} else if method == 2 && method2(low, high, character, password) {
			validPasswords++
		}
	}
	return validPasswords, nil
}

func main() {
	lines, readLinesErr := ReadLines(inputFile)
	if readLinesErr != nil {
		fmt.Println("Error reading file ", readLinesErr)
		return
	}

	validPasswordsMethod1, calculateValidPasswordErr := calculateValidPasswords(lines, 1)
	if calculateValidPasswordErr != nil {
		fmt.Println("Error calculating valid passwords ", readLinesErr)
		return
	}
	fmt.Printf("%d valid passwords using method 1\n", validPasswordsMethod1)

	validPasswordsMethod2, calculateValidPasswordErr := calculateValidPasswords(lines, 2)
	if calculateValidPasswordErr != nil {
		fmt.Println("Error calculating valid passwords ", readLinesErr)
		return
	}
	fmt.Printf("%d valid passwords using method 2\n", validPasswordsMethod2)
}
