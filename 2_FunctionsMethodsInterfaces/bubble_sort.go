// Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers. The program should print the integers out on one line, in sorted order, from least to greatest. Use your favorite search tool to find a description of how the bubble sort algorithm works.

// As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted order.

// A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent elements in the slice. You should write a Swap() function which performs this operation. Your Swap() function should take two arguments, a slice of integers and an index value i which indicates a position in the slice. The Swap() function should return nothing, but it should swap the contents of the slice in position i with the contents in position i+1.

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Printf("Enter in a list of up to ten numbers separated by commas: ")
	fmt.Scan(&input)
	inputNumbers := strings.Split(input, ",")

	if len(inputNumbers) > 10 {
		fmt.Println("Sorry, you input more than 10 digits")
		os.Exit(1)
	}
	fmt.Println("*** Starting order: ", inputNumbers)
	mySlice := make([]int, 0)

	for i := 0; i < len(inputNumbers); i++ {
		number, err := strconv.Atoi(inputNumbers[i])
		if err != nil {
			// no op for now
		}
		mySlice = append(mySlice, number)
	}

	bubbleSort(mySlice)
}

func bubbleSort(digits []int) {
	swapped := false

	for i := 0; i < len(digits); i++ {
		if i+1 != len(digits) {
			if digits[i] > digits[i+1] {
				swapped = true

				swap(&digits, i)
			}
		}

		if i+1 == len(digits) && swapped == false {
			fmt.Println("*** Bubble-sorted result: ", digits)
			break
		} else if i+1 == len(digits) {
			bubbleSort(digits)
		}
	}

}

func swap(sliceOfInts *[]int, slicePos int) {
	temp := (*sliceOfInts)[slicePos]
	(*sliceOfInts)[slicePos] = (*sliceOfInts)[slicePos+1]
	(*sliceOfInts)[slicePos+1] = temp
}
