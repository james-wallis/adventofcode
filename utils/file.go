package utils

import (
	"io/ioutil"
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
