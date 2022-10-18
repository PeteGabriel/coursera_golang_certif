/*
Write a program which allows the user to get information about a predefined set of animals.
Three animals are predefined, cow, bird, and snake.
Each animal can eat, move, and speak.

The user can issue a request to find out one of three things about an animal:
	1) the food that it eats
	2) its method of locomotion
	3) the sound it makes when it speaks.


An example of the command could be: "bird speak"
*/

package main

import (
	"fmt"
)

type Animal struct {
	kind, food, locomotion, sound string
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Speak() {
	fmt.Println(a.sound)
}

func main() {
	cow := Animal{
		kind:       "cow",
		food:       "grass",
		locomotion: "walk",
		sound:      "moo",
	}

	bird := Animal{
		kind:       "bird",
		food:       "worms",
		locomotion: "fly",
		sound:      "peep",
	}

	snake := Animal{
		kind:       "snake",
		food:       "mice",
		locomotion: "slither",
		sound:      "hsss",
	}

	var kind, action string
	for {
		fmt.Print("> ")
		fmt.Scan(&kind, &action)
		switch kind {
		case "cow":
			cow.act(action)
		case "snake":
			snake.act(action)
		case "bird":
			bird.act(action)
		default:
			fmt.Println("unknown animal kind")
		}
	}
}

func (a Animal) act(action string) {
	switch action {
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	case "eat":
		a.Eat()
	default:
		fmt.Println("unknown action")
	}
}
