package main

import (
	"fmt"
	"testing"
)

func TestCountLettersInString(t *testing.T) {
	var cases = []struct {
		input string
		want  int
	}{
		{

			"abc",
			3,
		},
		{

			"a\nb\nc",
			3,
		},
		{

			"ab\nac",
			3,
		},
		{

			"a\na\na\na",
			1,
		},
		{
			"b",
			1,
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("returns %d when the input is %s", test.want, test.input), func(t *testing.T) {
			got := CountLettersInString(test.input)
			if got != test.want {
				t.Errorf("got %d want %d", got, test.want)
			}
		})
	}
}

func TestCountQuestionsAnswered(t *testing.T) {
	var cases = []struct {
		input []string
		want  int
	}{
		{

			[]string{
				"abc",
				"a\nb\nc",
				"ab\nac",
				"a\na\na\na",
				"b",
			},
			11,
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("returns %d when the input is %s", test.want, test.input), func(t *testing.T) {
			got := CountQuestionsAnswered(test.input)
			if got != test.want {
				t.Errorf("got %d want %d", got, test.want)
			}
		})
	}
}

func TestCountLettersInEachLine(t *testing.T) {
	var cases = []struct {
		input string
		want  int
	}{
		{

			"abc",
			3,
		},
		{

			"a\nb\nc",
			0,
		},
		{

			"ab\nac",
			1,
		},
		{

			"a\na\na\na",
			1,
		},
		{
			"b",
			1,
		},
		{

			"a\na\na\nb",
			0,
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("returns %d when the input is %s", test.want, test.input), func(t *testing.T) {
			got := CountLettersInEachLine(test.input)
			if got != test.want {
				t.Errorf("got %d want %d", got, test.want)
			}
		})
	}
}

func TestCountQuestionsAnsweredByEveryone(t *testing.T) {
	var cases = []struct {
		input []string
		want  int
	}{
		{

			[]string{
				"abc",
				"a\nb\nc",
				"ab\nac",
				"a\na\na\na",
				"b",
			},
			6,
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("returns %d when the input is %s", test.want, test.input), func(t *testing.T) {
			got := CountQuestionsAnsweredByEveryone(test.input)
			if got != test.want {
				t.Errorf("got %d want %d", got, test.want)
			}
		})
	}
}
