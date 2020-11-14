// Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3. During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

package main

import (
	"fmt"
	"sort"
)

func main() {
	var input int
	mySlice := make([]int, 3)
	sum := 0
	for i := 1; i < len(mySlice)+1; i++ {
		fmt.Printf("Enter in an integer: ")
		fmt.Scan(&input)
		// Add the integer to the slice
		// Sort the slice
		fmt.Println(fmt.Sprintf("Sum: %d", sum))
		if input == 555 {
			break
		} else {
			mySlice = append(mySlice, input)
			fmt.Println(fmt.Sprintf("You added: %d", input))
			sort.Ints(mySlice)
			fmt.Println(fmt.Sprintf("Updated sorted list: %d", mySlice))
			sum++
		}
	}

	// begin loop

	// Or exit if the user has typed in 'X'
	// end loop

}
