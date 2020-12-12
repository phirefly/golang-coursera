package main

import "fmt"

func main() {
	var inputAcceleration float64
	var inputInitialVelocity float64
	var inputInitialDisplacement float64
	var inputTime float64

	fmt.Print("Enter in a value for 'acceleration':")
	fmt.Scan(&inputAcceleration)
	fmt.Print("Enter in a value for 'initial velocity':")
	fmt.Scan(&inputInitialVelocity)
	fmt.Print("Enter in a value for 'initial displacement':")
	fmt.Scan(&inputInitialDisplacement)
	fmt.Println("This is what you input...")
	fmt.Println("acceleration:", inputAcceleration)
	fmt.Println("initial velocity:", inputInitialVelocity)
	fmt.Println("initial displacement:", inputInitialDisplacement)

	myFunc := GenDisplaceFn(inputAcceleration, inputInitialVelocity, inputInitialDisplacement)

	fmt.Println("Your function has been created! Enter in a value for 'time' to use this function:")
	fmt.Scan(&inputTime)
	fmt.Println("Your result: ", myFunc(inputTime))
}

func GenDisplaceFn(a float64, vo float64, so float64) func(t float64) float64 {
	fn:= func(t float64) float64 {
		return .5 * a * t * t +
			vo * t +
			so
	}
	return fn
}
