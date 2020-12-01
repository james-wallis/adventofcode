package main

import "testing"

func TestCalculateWhichTwoNumbersMake2020(t *testing.T) {
	t.Run("collection of 10 numbers that will make 2020", func(t *testing.T) {
		numbers := []int64{1009, 1, 2, 4, 5, 6, 7, 8, 9, 1011}

		want1 := 1009
		want2 := 1011
		num1, num2 := CalculateWhichTwoNumbersMake2020(numbers)

		if num1 != 1009 && num2 != 1011 {
			// %v is the default format, useful for arrays
			t.Errorf("got %d and %d want %d and %d  given, %v", num1, num2, want1, want2, numbers)
		}
	})

	t.Run("collection of 10 numbers that will not make 2020", func(t *testing.T) {
		numbers := []int64{10, 1, 2, 4, 5, 6, 7, 8, 9, 100}

		want := -1
		num1, num2 := CalculateWhichTwoNumbersMake2020(numbers)

		if num1 != -1 && num2 != -1 {
			// %v is the default format, useful for arrays
			t.Errorf("got %d and %d want %d and %d  given, %v", num1, num2, want, want, numbers)
		}
	})
}
