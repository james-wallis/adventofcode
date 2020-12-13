package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// ReadLines reads a file, splits it into lines returns an array of strings
func ReadLines(path string) ([]string, error) {
	content, readFileErr := ioutil.ReadFile(path)
	if readFileErr != nil {
		return nil, readFileErr
	}
	lines := strings.Split(string(content), "\n")

	return lines, nil
}

// ReadLinesWithSeparator reads a file, splits it into lines using a custom separator returns an array of strings
func ReadLinesWithSeparator(path string, sep string) ([]string, error) {
	content, readFileErr := ioutil.ReadFile(path)
	if readFileErr != nil {
		return nil, readFileErr
	}
	lines := strings.Split(string(content), sep)

	return lines, nil
}

// ReadLinesAndConvertToInt : reads a file splitting using the newline character and converts the contents to ints
func ReadLinesAndConvertToInt(path string) ([]int, error) {
	lines, readLinesErr := ReadLines(path)
	if readLinesErr != nil {
		return nil, readLinesErr
	}

	linesAsInts := []int{}
	for _, line := range lines {
		convertedNum, conversionErr := strconv.Atoi(line)
		if conversionErr != nil {
			return nil, conversionErr
		}
		linesAsInts = append(linesAsInts, convertedNum)
	}

	return linesAsInts, nil
}
