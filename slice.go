// Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3. During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var input string
	mySlice := make([]int, 3)
	sum := 0
	for i := 1; i < len(mySlice)+2; i++ { // This is pretty ugly
		fmt.Printf("Enter in an integer: ")
		fmt.Scan(&input)
		// Add the integer to the slice
		// Sort the slice
		fmt.Println(fmt.Sprintf("Sum: %d", sum))
		if input == "X" {
			break
		} else {
			if intinput, err := strconv.Atoi(input); err == nil {
				mySlice = append(mySlice, intinput)
				if i <= 3 { //This will break the loop if we delete 3. keeping at 2 so you can see it work.
					mySlice = mySlice[1:4]
				}

				fmt.Println(fmt.Sprintf("You added: %d", intinput))
				sort.Ints(mySlice)
				sum++
			} else {
				fmt.Println(intinput, " is not an integer")
			}
		}
		fmt.Println(fmt.Sprintf("Updated sorted list: %d", mySlice))
	}

}
