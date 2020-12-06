package main

import "fmt"

func fA() func() int {
	//The value for i remains, so it increments
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	fB := fA()
	fmt.Print(fB())
	fmt.Print(fB())

}
