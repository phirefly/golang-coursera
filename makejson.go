// Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

// Submit your source code for the program, “makejson.go”.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Address   string `json:"address"`
}

func main() {
	var input_first string
	var input_last string
	var input_address string

	newperson := &Person{
		Firstname: "default first",
		Lastname:  "default last",
		Address:   "default address",
	}

	fmt.Printf("Enter your first name: ")
	fmt.Scan(&input_first)
	newperson.Firstname = input_first

	fmt.Printf("Enter your last name: ")
	fmt.Scan(&input_last)
	newperson.Lastname = input_last

	fmt.Printf("Enter your address: ")
	bufscanner := bufio.NewScanner(os.Stdin)
	bufscanner.Scan()
	input_address = bufscanner.Text()
	newperson.Address = input_address

	barr, error := json.MarshalIndent(newperson, "", " ")
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(string(barr))

}
