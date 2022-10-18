package main

import (
	"fmt"
	"math"
)

func main2() {

	//enter values for acceleration, initial velocity, and initial displacement
	var acc, vel, initDisp, t float64
	fmt.Println("enter a value for acceleration:")
	fmt.Scanln(&acc)
	fmt.Println("enter a value for initial velocity:")
	fmt.Scanln(&vel)
	fmt.Println("enter a value for initial displacement:")
	fmt.Scanln(&initDisp)

	fn := GenDisplaceFn2(acc, vel, initDisp)

	fmt.Println("enter a value for time:")
	fmt.Scanln(&t)

	fmt.Println(fn(t))
}

//GenDisplaceFn takes acceleration a, initial velocity v_o, and initial
//displacement s_o and return a new function which computes displacement as a function of time.
func GenDisplaceFn2(a, v_o, s_o float64) func(float64) float64 {
	return func(t float64) float64 {
		return (0.5 * a * (math.Pow(t, 2))) + (v_o * t) + s_o
	}
}
