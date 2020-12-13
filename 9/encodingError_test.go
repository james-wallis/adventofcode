package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDiffNumbersEqualNumber(t *testing.T) {
	var cases = []struct {
		num1   int
		num2   int
		result int
		want   bool
	}{
		{

			1,
			2,
			3,
			true,
		},
		{

			2,
			3,
			1,
			false,
		},
		{

			2,
			2,
			4,
			false,
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("returns %t when the input is %d, %d and %d", test.want, test.num1, test.num2, test.result), func(t *testing.T) {
			got := DiffNumbersEqualNumber(test.num1, test.num2, test.result)
			if got != test.want {
				t.Errorf("lineIncrement: got %t want %t", got, test.want)
			}
		})
	}
}

func TestDoTwoNumbersInSliceMakeNumber(t *testing.T) {
	var cases = []struct {
		input []int
		num   int
		want  bool
	}{
		{

			[]int{
				1,
				2,
			},
			3,
			true,
		},
		{

			[]int{
				1,
				2,
				3,
				4,
				5,
			},
			9,
			true,
		},
		{

			[]int{
				1,
				2,
				3,
				4,
				5,
			},
			60,
			false,
		},
		{

			[]int{
				1,
				2,
				3,
				4,
				5,
			},
			3,
			true,
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("returns %t when the input is %v and the given num is %d", test.want, test.input, test.num), func(t *testing.T) {
			got := DoTwoNumbersInSliceMakeNumber(test.input, test.num)
			if got != test.want {
				t.Errorf("lineIncrement: got %t want %t", got, test.want)
			}
		})
	}
}

func TestFindNumNotSumOfPreviousNumbers(t *testing.T) {
	var cases = []struct {
		input    []int
		preamble int
		want     int
	}{
		{
			[]int{
				35,
				20,
				15,
				25,
				47,
				40,
				62,
				55,
				65,
				95,
				102,
				117,
				150,
				182,
				127,
				219,
				299,
				277,
				309,
				576,
			},
			5,
			127,
		},
		{
			[]int{
				35,
				20,
				15,
				55,
			},
			3,
			-1,
		},
		{
			[]int{
				1,
				2,
				3,
				5,
				7,
				8,
			},
			2,
			7,
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("returns %d when the input is %v and the preamble is %d", test.want, test.input, test.preamble), func(t *testing.T) {
			got := FindNumNotSumOfPreviousNumbers(test.input, test.preamble)
			if got != test.want {
				t.Errorf("lineIncrement: got %d want %d", got, test.want)
			}
		})
	}
}

func TestFindNumbersInSliceThatMakeNum(t *testing.T) {
	var cases = []struct {
		input  []int
		target int
		want   []int
	}{
		{
			[]int{
				35,
				20,
				15,
				25,
				47,
				40,
				62,
				55,
				65,
				95,
				102,
				117,
				150,
				182,
				127,
				219,
				299,
				277,
				309,
				576,
			},
			127,
			[]int{
				15,
				25,
				47,
				40,
			},
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("returns %d when the input is %v and the target number is %d", test.want, test.input, test.target), func(t *testing.T) {
			got := FindNumbersInSliceThatMakeNum(test.input, test.target)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("lineIncrement: got %v want %v", got, test.want)
			}
		})
	}
}
