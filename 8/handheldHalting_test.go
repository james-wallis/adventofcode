package main

import (
	"fmt"
	"testing"
)

var input = []string{
	"nop +0",
	"acc +1",
	"jmp +4",
	"acc +3",
	"jmp -3",
	"acc -99",
	"acc +1",
	"jmp -4",
	"acc +6",
}

func TestParseOperation(t *testing.T) {
	var cases = []struct {
		operation     string
		lineIncrement int
		accIncrement  int
		op            string
	}{
		{

			"nop +0",
			1,
			0,
			"nop",
		},
		{

			"acc +1",
			1,
			1,
			"acc",
		},
		{

			"jmp +4",
			4,
			0,
			"jmp",
		},
		{

			"acc +3",
			1,
			3,
			"acc",
		},
		{
			"jmp -3",
			-3,
			0,
			"jmp",
		},
		{
			"acc -99",
			1,
			-99,
			"acc",
		},
		{
			"jmp -4",
			-4,
			0,
			"jmp",
		},
		{
			"acc +6",
			1,
			6,
			"acc",
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("returns %d and %d when the input is %s", test.lineIncrement, test.accIncrement, test.operation), func(t *testing.T) {
			gotLineIncrement, gotAccIncrement, gotOp := ParseOperation(test.operation)
			if gotLineIncrement != test.lineIncrement {
				t.Errorf("lineIncrement: got %d want %d", gotLineIncrement, test.lineIncrement)
			}

			if gotAccIncrement != test.accIncrement {
				t.Errorf("accIncrement: got %d want %d", gotAccIncrement, test.accIncrement)
			}

			if gotOp != test.op {
				t.Errorf("op: got %s want %s", gotOp, test.op)
			}
		})
	}
}

func TestDetermineAccWhenInstructionExecutedSecondTime(t *testing.T) {
	t.Run("should return 5 when the test input is used", func(t *testing.T) {
		want := 5
		got := DetermineAccWhenInstructionExecutedSecondTime(input)

		if got != want {
			t.Errorf("accIncrement: got %d want %d", got, want)
		}
	})
}

func TestChangeSingleNopOrJmpToFixInfiniteLoop(t *testing.T) {
	t.Run("should return 8 when the test input is used", func(t *testing.T) {
		want := 8
		got := ChangeSingleNopOrJmpToFixInfiniteLoop(input)

		if got != want {
			t.Errorf("accIncrement: got %d want %d", got, want)
		}
	})
}
