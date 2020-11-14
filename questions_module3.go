// Module 3 Questions

package main

import "fmt"


func main() {

	// Question #1
  x := []int {4, 8, 5}
  y := -1
  for _, elt := range x {
  	fmt.Println("elt: ", elt)
    if elt > y {
      y = elt
    }
  }
  fmt.Print(y)
  // Answer: 8


//   Question #2
  x := [...]int {4, 8, 5}
  y := x[0:2] //4,8 => 1,8 
  z := x[1:3] //8,5 => 8,3
  y[0] = 1
  z[1] = 3
  fmt.Print(x)
}

// Answer: 1,8,3

// Question #3
x := [...]int {1, 2, 3, 4, 5}
y := x[0:2] //1,2
z := x[1:4] //2,3,4
fmt.Print(len(y), cap(y), len(z), cap(z)) 
// Answer: 2, 5, 3, 4


// Question #4
x := map[string]int {
	"ian": 1, "harris": 2}
for i, j := range x {
	if i == "harris" {
	fmt.Print(i, j)
	}
}

// Answer: harris, 2



// Question #5
type P struct {
    x string
	y int
} 
func main() {
  b := P{"x", -1} // b.x => "x", b.y => -1
  a := [...]P{P{"a", 10}, 
        P{"b", 2},
        P{"c", 3}}
  for _, z := range a {
    if z.y > b.y {
      b = z
    }
  }
  fmt.Println(b.x)
}

// Answer: a



// Question #6
s := make([]int, 0, 3)
s = append(s, 100)
fmt.Println(len(s), cap(s))
// Answer: 1 3
// It creates a slice of length 0 and a capacity of 3
// It then adds the integer 100, which is why is returns a length of 1

}
