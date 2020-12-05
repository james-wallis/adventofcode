package main

import (
	"fmt"
	"testing"
)

func TestDetermineSeatLocation(t *testing.T) {
	var cases = []struct {
		input      string
		wantRow    int64
		wantColumn int64
	}{
		{
			"FBFBBFFRLR",
			44,
			5,
		},
		{
			"BFFFBBFRRR",
			70,
			7,
		},
		{
			"FFFBBBFRRR",
			14,
			7,
		},
		{
			"BBFFBBFRLL",
			102,
			4,
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("returns %d when the input is %s", test.wantRow, test.input), func(t *testing.T) {
			gotRow, gotColumn := DetermineSeatLocation(test.input)
			if gotRow != test.wantRow {
				t.Errorf("got %d want %d", gotRow, test.wantRow)
			}

			if gotColumn != test.wantColumn {
				t.Errorf("got %d want %d", gotColumn, test.wantColumn)
			}
		})
	}
}

func TestCalculateSeatId(t *testing.T) {
	var cases = []struct {
		input string
		want  int64
	}{
		{
			"FBFBBFFRLR",
			357,
		},
		{
			"BFFFBBFRRR",
			567,
		},
		{
			"FFFBBBFRRR",
			119,
		},
		{
			"BBFFBBFRLL",
			820,
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("returns %d when the input is %s", test.want, test.input), func(t *testing.T) {
			got := CalculateSeatID(test.input)
			if got != test.want {
				t.Errorf("got %d want %d", got, test.want)
			}
		})
	}
}
