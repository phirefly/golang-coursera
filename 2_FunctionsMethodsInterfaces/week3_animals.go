package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food, locomotion, noise string
}

func main() {
	// Make a type called Animal which is a struct containing three fields:food, locomotion, and noise, all of which are strings
	// Make three methods called Eat(), Move(), and Speak(). The receiver type of all of your methods should be
	// your Animal type. The Eat() method should print the animal’s food, the Move()  should print the
	// animal’s locomotion, and the Speak()  should print the animal’s spoken sound.

	var input string
	fmt.Println("*** Welcome! Type in 'cow|bird|snake' and 'eat|move|speak'")
	fmt.Printf("> ")
	bufscanner := bufio.NewScanner(os.Stdin)
	bufscanner.Scan()
	input = bufscanner.Text()

	parsedInput := strings.Split(input, " ")
	fmt.Println("animal: ", parsedInput[0])
	fmt.Println("action: ", parsedInput[1])
}

func Eat() {
	fmt.Println("*** Eat")
}

func Move() {
	fmt.Println("*** Move")
}

func Speak() {
	fmt.Println("*** Speak")
}
