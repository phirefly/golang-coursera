package main

import "fmt"

var f func(string) int

func test(x string) int {

	return len(x)

}

func main() {
	f = test
	fmt.Println(f("hello there"))
}
