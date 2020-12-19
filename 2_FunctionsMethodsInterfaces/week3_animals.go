package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
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
	cow := Animal{"grass", "walk", "moo"}
	_ = cow
	bird := Animal{"worms", "fly", "peep"}
	_ = bird
	snake := Animal{"mice", "slither", "hsss"}
	_ = snake

	var input string
	fmt.Println("*** Welcome! Type in 'cow|bird|snake' and 'eat|move|speak'")
	fmt.Printf("> ")
	bufscanner := bufio.NewScanner(os.Stdin)
	bufscanner.Scan()
	input = bufscanner.Text()

	parsedInput := strings.Split(input, " ")
	fmt.Println("animal: ", parsedInput[0])
	fmt.Println("action: ", parsedInput[1])
	chosenAction := parsedInput[1]
	_ = chosenAction
	// TODO: Translate animal to the variable. Capitalize the action verb and pass it in
	meth := reflect.ValueOf(cow).MethodByName(strings.Title(chosenAction))
	meth.Call(nil)
}

func (a Animal) Eat() {
	fmt.Println("*** eats:", a.food)
}

func (a Animal) Move() {
	fmt.Println("*** moves:", a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println("*** speaks:", a.noise)
}
