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

type Bird struct {
	food string
	locomotion string
	sound string
}

type Snake struct {
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

func (b Bird) Speak() {
	fmt.Println("speaking: ", b.sound)
}
func (b Bird) Move() {
	fmt.Println("moving: ", b.locomotion)
}
func (b Bird) Eat() {
	fmt.Println("eating: ", b.food)
}

func (s Snake) Speak() {
	fmt.Println("speaking: ", s.sound)
}
func (s Snake) Move() {
	fmt.Println("moving: ", s.locomotion)
}
func (s Snake) Eat() {
	fmt.Println("eating: ", s.food)
}


func main() {
	var myMap = make(map[string]Animal)
	myMap["cow"] = Cow{"grass", "walk", "moo"}
	myMap["bird"] = Bird{"worms", "fly", "peep"}
	myMap["snake"] = Snake{"mice", "slither", "hsss"}

	var aliasMap = map[string][]string{
		"cow": {"testcow"},
		"bird": {"testbird"},
		"snake": {"testsnake"},
	}
	//append(aliasMap["cow"], "blahcow")
	//append(aliasMap["bird"], "blahbird")
	//append(aliasMap["snake"], "blahsnake")
	//
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
		chosenName := parsedInput[2]

		switch command {
		case "newanimal":
			fmt.Println("--- new animal was chosen")

			var a1 Animal
			var myAnimal = myMap[chosenAnimal]
			a1 = myAnimal //Now the concrete type that a1 Animal is assigned to is myAnimal
			_ = a1
			myAnimal.Speak()

			//	CURRENT: Record the name in a map of aliases
			//  Getting... ./week4_animals.go:106:10: first argument to append must be slice; have struct { types.elem types.Type }
			res := append(aliasMap[chosenAnimal], chosenName)
			fmt.Println("**** Updated aliases")
			fmt.Println("res: ", res)

			aliasMap[chosenAnimal] = res
			fmt.Println("==> updated aliasMap: ", aliasMap[chosenAnimal])
			//fmt.Println(aliasMap[chosenAnimal])


		case "query":
			fmt.Println("--- query was chosen")
		default:
			fmt.Println("--- action not known")
		}

		//fmt.Println("*******")
		//fmt.Println("command: ", command)
		//fmt.Println("chosenAnimal: ", chosenAnimal)
		//fmt.Println("chosenAction: ", chosenAction)
		//
		//fmt.Println("--- Created it!---")
		//fmt.Println("")

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