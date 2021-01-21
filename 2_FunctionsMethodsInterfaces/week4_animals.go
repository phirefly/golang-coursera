package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func contains(arr[]string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}


func main() {
	var myMap = make(map[string]Animal)
	myMap["cow"] = Cow{"grass", "walk", "moo"}
	myMap["bird"] = Bird{"worms", "fly", "peep"}
	myMap["snake"] = Snake{"mice", "slither", "hsss"}

	var aliasMap = map[string][]string{
		"cow": {"defaultcow"},
		"bird": {"defaultbird"},
		"snake": {"defaultsnake"},
	}

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
			res := append(aliasMap[chosenAnimal], chosenName)
			aliasMap[chosenAnimal] = res
			fmt.Println("==> updated aliasMap: ", aliasMap[chosenAnimal])

		case "query":
			if (contains(aliasMap["cow"], chosenAnimal)) { // CURRENT: Look inside the alias group for this instead of myMap and iterate through the types
				var myAnimal = myMap["cow"]
				var a1 Animal
				a1 = myAnimal //Now the concrete type that a1 Animal is assigned to is myAnimal
				_ = a1

				switch chosenName {
				case "eat":
					myAnimal.Eat()
				case "move":
					myAnimal.Move()
				case "speak":
					myAnimal.Speak()
				default:
					fmt.Println("--- wrong action!")
				}
			} else {
				fmt.Println("--- Your animal wasn't found. Try again. ")
			}


		default:
			fmt.Println("--- action not known")
		}
	}
}

