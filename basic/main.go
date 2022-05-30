package main

import (
	"fmt"
	"os"
)

func main() {

	// Get total exam count
	var examCount uint
	fmt.Printf("Please type your exam count: ")
	_, err := fmt.Scanf("%v", &examCount)

	// Declare exam list with total point
	// Attention: If you use dynamic array range, probably get "Invalid array bound 'examCount', must be a constant expression" error.
	// Etc. var examPoints []float32
	// You can handle it with the definition below.
	var examPoints []uint = make([]uint, uint(examCount))
	var examsTotal uint = 0

	// Get each exam point
	var i uint = 0
	for i < uint(examCount) {
		// Get exam 1
		fmt.Printf("\nPlease type your exam %v: ", i+1)
		_, err = fmt.Scanf("%v", &examPoints[i])

		/*
			if err != nil {
				fmt.Printf("\nUnaccountable type, your datum must be int or float type")
				os.Exit(0)
			} else if examPoints[i] < 0 {
				fmt.Printf("\nUnaccountable type, your datum must be int or float type")
				os.Exit(0)
			}
		*/

		switch {
		case err != nil:
		case examPoints[i] < 0:
			fmt.Printf("\nUnaccountable type, your datum must be int or float type")
			os.Exit(0)
		}
		examsTotal += examPoints[i]
		i++
	}

	// Get exam average
	var examAverage float32 = float32(examsTotal / examCount)
	fmt.Printf("Your numeric score: %.2f\n", examAverage)

	// Detect average letter score and print on console line
	switch {
	case examAverage <= 29:
		fmt.Println("Your score: FF")
	case examAverage <= 39:
		fmt.Println("Your score: FD")
	case examAverage <= 49:
		fmt.Println("Your score: DD")
	case examAverage <= 59:
		fmt.Println("Your score: DC")
	case examAverage <= 69:
		fmt.Println("Your score: CC")
	case examAverage <= 79:
		fmt.Println("Your score: BB")
	case examAverage <= 89:
		fmt.Println("Your score: BA")
	case examAverage <= 100:
		fmt.Println("Your score: AA")
	default:
		fmt.Println("Unaccountable result")
	}

	// Range using with an example.
	println("----------------------------")
	println("Your step by step calculations:")

	var stepTotal float32 = 0
	for i, v := range examPoints {
		stepTotal += float32(v)
		fmt.Printf("%v. Step total: %v\n", i+1, stepTotal)
	}
	fmt.Printf("Your step by step calculations total: %v\n", stepTotal)
}
