package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/james-wallis/adventofcode/utils"
)

func TestNumOccupiedAdjacentSeats(t *testing.T) {
	var cases = []struct {
		inputFile string
		x         int
		y         int
		want      int
	}{
		{
			"./testInitial.txt",
			0,
			0,
			0,
		},
		{
			"./testFirstRound.txt",
			0,
			0,
			2,
		},
		{
			"./testFirstRound.txt",
			4,
			8,
			8,
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("returns %v", test.want), func(t *testing.T) {
			input, _ := utils.ReadLinesAndSplitCharacters(test.inputFile)
			got := NumOccupiedAdjacentSeats(input, test.x, test.y)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v want %v", got, test.want)
			}
		})
	}
}

func TestIterateOverBoard(t *testing.T) {
	var cases = []struct {
		inputFile string
		wantFile  string
	}{
		{
			"./testInitial.txt",
			"./testFirstRound.txt",
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("returns the contens of %s", test.wantFile), func(t *testing.T) {
			input, _ := utils.ReadLinesAndSplitCharacters(test.inputFile)
			want, _ := utils.ReadLinesAndSplitCharacters(test.wantFile)
			got := IterateOverBoard(input)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
