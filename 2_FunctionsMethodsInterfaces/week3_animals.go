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
	animal := make(map[string]Animal)
	animal["cow"] = Animal{"grass", "walk", "moo"}
	animal["bird"] = Animal{"worms", "fly", "peep"}
	animal["snake"] = Animal{"mice", "slither", "hsss"}

	var input string
	fmt.Println("*** Welcome! Type in 'cow|bird|snake' and 'eat|move|speak'")
	fmt.Printf("> ")
	bufscanner := bufio.NewScanner(os.Stdin)
	bufscanner.Scan()
	input = bufscanner.Text()

	parsedInput := strings.Split(input, " ")
	chosenAnimal := parsedInput[0]
	chosenAction := parsedInput[1]

	meth := reflect.ValueOf(animal[chosenAnimal]).MethodByName(strings.Title(chosenAction))
	meth.Call(nil)
}

func (a Animal) Eat() {
	fmt.Println("=>", a.food)
}

func (a Animal) Move() {
	fmt.Println("=>", a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println("=>", a.noise)
}
