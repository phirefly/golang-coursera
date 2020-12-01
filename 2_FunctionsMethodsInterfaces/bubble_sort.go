// Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers. The program should print the integers out on one line, in sorted order, from least to greatest. Use your favorite search tool to find a description of how the bubble sort algorithm works.

// As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted order.

// A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent elements in the slice. You should write a Swap() function which performs this operation. Your Swap() function should take two arguments, a slice of integers and an index value i which indicates a position in the slice. The Swap() function should return nothing, but it should swap the contents of the slice in position i with the contents in position i+1.

package main

import "fmt"

func main() {
	mySlice := make([]int, 0)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)
	mySlice = append(mySlice, 20)
	mySlice = append(mySlice, 15)
	mySlice = append(mySlice, 33)
	mySlice = append(mySlice, 12)
	bubbleSort(mySlice)
}

func bubbleSort(digits []int) {
	fmt.Println("digits", digits)
	swapped := false

	for i := 0; i < len(digits); i++ {
		// fmt.Println("length: ", len(digits))
		if i+1 != len(digits) {
			if digits[i] > digits[i+1] {
				swapped = true

				swap()
				temp := digits[i]
				fmt.Println("temp", temp)
				digits[i] = digits[i+1]
				fmt.Println("digits[", i, "]:", digits[i])
				digits[i+1] = temp
				fmt.Println("digits[", i+1, "]:", digits[i+1])
				fmt.Println(digits)
			}
		}

		if i+1 == len(digits) && swapped == false {
			fmt.Println("*** nothing swapped! ")
			fmt.Println("*** Final: ", digits)
			break
		} else if i+1 == len(digits) {
			bubbleSort(digits)
		}
	}

}

func swap() {
	fmt.Println("*** swap ***")
}
