package main

import (
	"fmt"
)

func mostRepeated(data []string) string {
	counts := make(map[string]int)

	// Count occurrences of each element in the array
	for _, item := range data {
		counts[item]++
	}

	var mostRepeatedData string
	maxCount := 0

	// Find the most repeated data
	for item, count := range counts {
		if count > maxCount {
			maxCount = count
			mostRepeatedData = item
		} else if count == maxCount {
			// If there is equality, append to the result string
			mostRepeatedData += ", " + item
		}
	}

	return mostRepeatedData
}

func main() {
	input := []string{"apple", "pie", "apple", "red", "red", "red", "apple"}
	output := mostRepeated(input)

	fmt.Println("Input:")
	fmt.Println(input)
	fmt.Println("Output:")
	fmt.Printf("\"%s\"\n", output)
}
