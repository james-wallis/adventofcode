package main

import "fmt"

const inputFile = "./input.txt"

func main() {
	numbers, readLinesErr := ReadLines(inputFile)
	if readLinesErr != nil {
		fmt.Println("Error reading file ", readLinesErr)
	}

	x, y := CalculateWhichTwoNumbersMake2020(numbers)
	if x == -1 && y == -1 {
		fmt.Println("Unable to find two numbers that add up to make 2020")
	} else {
		fmt.Println("Success!")
		fmt.Printf("%d + %d = 2020\n", x, y)
		fmt.Printf("%d * %d = %d\n", x, y, x*y)
		fmt.Printf("answer should be %d\n\n", 913824)
	}

	x, y, z := CalculateWhichThreeNumbersMake2020(numbers)
	if x == -1 && y == -1 && z == -1 {
		fmt.Println("Unable to find two numbers that add up to make 2020")
	} else {
		fmt.Println("Success!")
		fmt.Printf("%d + %d + %d = 2020\n", x, y, z)
		fmt.Printf("%d * %d * %d = %d\n", x, y, z, x*y*z)
		fmt.Printf("answer should be %d\n", 240889536)
	}
}
