package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCountDifferencesOf1Or3InSlice(t *testing.T) {
	var cases = []struct {
		input []int
		want  map[int]int
	}{
		{

			[]int{
				2, 3, 4, 1,
			},
			map[int]int{
				1: 4,
				3: 1,
			},
		},
		{

			[]int{
				10, 8, 7, 6, 4, 1, 12,
			},
			map[int]int{
				1: 3,
				3: 2,
			},
		},
		{

			[]int{
				16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4,
			},
			map[int]int{
				1: 7,
				3: 5,
			},
		},
		{

			[]int{
				28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3,
			},
			map[int]int{
				1: 22,
				3: 10,
			},
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("returns %v when the input is %v", test.want, test.input), func(t *testing.T) {
			got := CountDifferencesOf1Or3InSlice(test.input)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v want %v", got, test.want)
			}
		})
	}
}

func TestSplitSliceIntoSubMapsOn3Difference(t *testing.T) {
	var cases = []struct {
		input []int
		want  []map[int]bool
	}{
		{

			[]int{
				16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4,
			},
			[]map[int]bool{
				map[int]bool{
					1: true,
				},
				map[int]bool{
					4: true,
					5: true,
					6: true,
					7: true,
				},
				map[int]bool{
					10: true,
					11: true,
					12: true,
				},
				map[int]bool{
					15: true,
					16: true,
				},
				map[int]bool{
					19: true,
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("returns %v when the input is %v", test.want, test.input), func(t *testing.T) {
			got := SplitSliceIntoSubMapsOn3Difference(test.input)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v want %v", got, test.want)
			}
		})
	}
}

func TestCountAllArrangementsOfMap(t *testing.T) {
	var cases = []struct {
		input map[int]bool
		want  int
	}{
		{
			map[int]bool{
				1: true,
			},
			1,
		},
		{
			map[int]bool{
				4: true,
				5: true,
				6: true,
				7: true,
			},
			4,
		},
		{
			map[int]bool{
				10: true,
				11: true,
				12: true,
			},
			2,
		},
		{
			map[int]bool{
				15: true,
				16: true,
			},
			1,
		},
		{
			map[int]bool{
				19: true,
			},
			1,
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("returns %v when the input is %v", test.want, test.input), func(t *testing.T) {
			got := CountAllArrangementsOfMap(test.input)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v want %v", got, test.want)
			}
		})
	}
}

func TestCountAllArrangementsOfSlice(t *testing.T) {
	var cases = []struct {
		input []int
		want  int
	}{
		{
			[]int{
				16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4,
			},
			8,
		},
		{
			[]int{
				28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3,
			},
			19208,
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("returns %v when the input is %v", test.want, test.input), func(t *testing.T) {
			got := CountAllArrangementsOfSlice(test.input)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v want %v", got, test.want)
			}
		})
	}
}
