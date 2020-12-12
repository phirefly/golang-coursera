package main

import (
	"fmt"
)

func GenDisplaceFn(acceleration float64, initialVelocity float64,
	intitalDisplacement float64) func(float64) float64 {
	return func(time float64) float64 {
		displacementTravelled := 0.5*acceleration*(time*time) + (initialVelocity * time) + intitalDisplacement
		return displacementTravelled
	}
}

func main() {
	var acceleration, initialVelocity, intitalDisplacement, time float64
	fmt.Println("Enter acceleration value")
	fmt.Scanln(&acceleration)
	fmt.Println("Enter initial velocity value")
	fmt.Scanln(&initialVelocity)
	fmt.Println("Enter initial displacement value")
	fmt.Scanln(&intitalDisplacement)
	fn := GenDisplaceFn(acceleration, initialVelocity, intitalDisplacement)
	fmt.Println("Enter value for time")
	fmt.Scanln(&time)
	fmt.Println("Displacemet Travelled  :", fn(time))
}
