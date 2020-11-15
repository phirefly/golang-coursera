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
	Fullname string `json:"fullname"`
	Address  string `json:"address"`
}

func main() {
	var input_fullname string
	var input_address string

	newperson := &Person{
		Fullname: "default first",
		Address:  "default address",
	}

	fmt.Printf("Enter your full name: ")
	bufscanner1 := bufio.NewScanner(os.Stdin)
	bufscanner1.Scan()
	input_fullname = bufscanner1.Text()
	newperson.Fullname = input_fullname

	fmt.Printf("Enter your address: ")
	bufscanner2 := bufio.NewScanner(os.Stdin)
	bufscanner2.Scan()
	input_address = bufscanner2.Text()
	newperson.Address = input_address

	barr, error := json.MarshalIndent(newperson, "", " ")
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(string(barr))

}
