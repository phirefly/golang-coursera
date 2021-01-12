package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//type Animal struct {
//	food, locomotion, noise string
//}
//
type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food string
	locomotion string
	sound string
}

func (c Cow) Speak() {
	fmt.Println("speaking: ", c.sound)
}
func (c Cow) Move() {
	fmt.Println("moving: ", c.locomotion)
}
func (c Cow) Eat() {
	fmt.Println("eating: ", c.food)
}

func main() {
	var cow1 = Cow{"grass", "walk", "moo"}
	cow1.Speak()

	for {
		var input string
		fmt.Println("*** Type in 'newanimal' and '<new animal name>' and 'cow|bird|snake'")
		fmt.Println("*** OR 'query' and '<animal name>' and 'eat|move|speak'")
		fmt.Printf("> ")
		bufscanner := bufio.NewScanner(os.Stdin)
		bufscanner.Scan()
		input = bufscanner.Text()

		parsedInput := strings.Split(input, " ")
		command := parsedInput[0]
		chosenAnimal := parsedInput[1]
		chosenAction := parsedInput[2]

		switch command {
		case "newanimal":
			fmt.Println("--- new animal was chosen")
		case "query":
			fmt.Println("--- query was chosen")
		default:
			fmt.Println("--- action not known")
		}

		fmt.Println("*******")
		fmt.Println("command: ", command)
		fmt.Println("chosenAnimal: ", chosenAnimal)
		fmt.Println("chosenAction: ", chosenAction)

		fmt.Println("--- Created it!---")
		fmt.Println("")

		//meth := reflect.ValueOf(cow1[chosenAnimal]).MethodByName(strings.Title(chosenAction))
		//meth.Call(nil)
	}
}

//func (a Animal) Eat() {
//	fmt.Println("=>", a.food)
//}

//func (a Animal) Move() {
//	fmt.Println("=>", a.locomotion)
//}
//
//func (a Animal) Speak() {
//	fmt.Println("=>", a.noise)
//}
//
//func AnimalEat (a Animal) bool {
//	switch := sh := a. (type) {
//		case Cow:
//			fmt.Printf("Cow")
//		case Bird:
//			fmt.Printf("Bird")
//		case Snake:
//			fmt.Printf("Snake")
//		}
//}