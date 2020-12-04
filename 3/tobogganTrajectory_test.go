package main

import "testing"

func TestGetIndexToRightOrWrap(t *testing.T) {
	t.Run("returns 3 when index is 0, line length is 4 and it increments by 3 to the right", func(t *testing.T) {
		line := "XXXX"
		currentIndex := 0

		want := 3
		got := GetIndexToRightOrWrap(line, currentIndex, 3)

		if got != want {
			t.Errorf("got %d want %d given, %s and %d", got, want, line, currentIndex)
		}
	})

	t.Run("returns 1 when index is 2, line length is 4 and it increments by 3 to the right", func(t *testing.T) {
		line := "XXXX"
		currentIndex := 2

		want := 1
		got := GetIndexToRightOrWrap(line, currentIndex, 3)

		if got != want {
			t.Errorf("got %d want %d given, %s and %d", got, want, line, currentIndex)
		}
	})

	t.Run("returns 0 when index is 1, line length is 4 and it increments by 3 to the right", func(t *testing.T) {
		line := "XXXX"
		currentIndex := 1

		want := 0
		got := GetIndexToRightOrWrap(line, currentIndex, 3)

		if got != want {
			t.Errorf("got %d want %d given, %s and %d", got, want, line, currentIndex)
		}
	})

	t.Run("returns 0 when index is 1, line length is 4 and it increments by 1 to the right", func(t *testing.T) {
		line := "XXXX"
		currentIndex := 1

		want := 2
		got := GetIndexToRightOrWrap(line, currentIndex, 1)

		if got != want {
			t.Errorf("got %d want %d given, %s and %d", got, want, line, currentIndex)
		}
	})
}
