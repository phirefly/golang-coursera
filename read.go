// The goal of this assignment is to practice working with the ioutil and os packages in Go.
// Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names. Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

// Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string of size 20 (characters).

// Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file. Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file. After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.

// Submit your source code for the program, “read.go”.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func main() {
	// Create an empty slice of struct pointers
	nameCollection := []*Person{}

	fmt.Println("Welcome!")
	// fmt.Printf("Enter the file you'd like to process: ")
	// fmt.Scan(&input)
	// file, e := ioutil.ReadFile("./names.txt")

	file, e := os.Open("./names.txt")
	if e != nil {
		log.Fatal("*** Something went wrong: ", e)
	}

	scanner := bufio.NewScanner(file)
	newperson := new(Person)

	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		words := strings.Fields(scanner.Text())
		// fmt.Println(words[0], ":", words[1])

		// Create a new struct for this person
		newperson = new(Person)
		newperson.Firstname = words[0]
		newperson.Lastname = words[1]

		nameCollection = append(nameCollection, newperson)
	}

	fmt.Println("***")
	fmt.Println("Total names processed: ", len(nameCollection))
	fmt.Println("***")

	for i, v := range nameCollection {
		fmt.Println(i, ":", v.Firstname, v.Lastname)
	}
}
