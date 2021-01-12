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
type Mover interface { Move() }
type Eater interface { Eat() }

type Cow struct {
	food string
	locomotion string
	sound string
}

func (c Cow) Speak() {
	fmt.Println(c.sound)
}
func (c Cow) Move() {
	fmt.Println(c.locomotion)
}
func (c Cow) Eat() {
	fmt.Println(c.food)
}

func main() {
	//var s1 Speaker
	//var m1 Mover
	//var e1 Eater

	var cow1 = Cow{"grass", "walk", "moo"}
	cow1.Speak()


	//animal := make(map[string]Animal)
	//animal["cow"] = Animal{"grass", "walk", "moo"}
	//animal["bird"] = Animal{"worms", "fly", "peep"}
	//animal["snake"] = Animal{"mice", "slither", "hsss"}

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

		fmt.Println("*******")
		fmt.Println("command: ", command)
		fmt.Println("chosenAnimal: ", chosenAnimal)
		fmt.Println("chosenAction: ", chosenAction)

		fmt.Println("--- Created it!", chosenAction)

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