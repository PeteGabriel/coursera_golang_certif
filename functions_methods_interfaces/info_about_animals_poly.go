package main

import (
	"errors"
	"fmt"
	"strings"
)

const (
	CMD   = "newanimal"
	QUERY = "query"
	COW   = "cow"
	BIRD  = "bird"
	SNAKE = "snake"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

type Bird struct {
	name string
}

type Snake struct {
	name string
}

func (Cow) Move() {
	fmt.Println("walk")
}

func (Cow) Eat() {
	fmt.Println("grass")
}

func (Cow) Speak() {
	fmt.Println("moo")
}

func (Bird) Move() {
	fmt.Println("fly")
}

func (Bird) Eat() {
	fmt.Println("worms")
}

func (Bird) Speak() {
	fmt.Println("peep")
}

func (Snake) Move() {
	fmt.Println("slither")
}

func (Snake) Eat() {
	fmt.Println("mice")
}

func (Snake) Speak() {
	fmt.Println("hsss")
}

var animals map[string]Animal

func main() {

	animals = make(map[string]Animal, 0)
	var op, name, typeOrAction string
	for {
		fmt.Print("> ")
		fmt.Scan(&op, &name, &typeOrAction)

		if strings.ToLower(op) == CMD {
			//create new animal
			an, err := newAnimal(name, typeOrAction)
			if err != nil {
				fmt.Println("something wrong with this command !")
			} else {
				animals[name] = an
				fmt.Println("Created it!")
			}
		} else if strings.ToLower(op) == QUERY {
			//ask for info about animal
			queryAnimal(name, typeOrAction)
		} else {
			fmt.Println("unknown operation")
		}
	}
}

func newAnimal(name, typeOrAction string) (Animal, error) {
	switch typeOrAction {
	case COW:
		return Cow{name: name}, nil
	case BIRD:
		return Bird{name: name}, nil
	case SNAKE:
		return Snake{name: name}, nil
	default:
		return nil, errors.New("unknown name in command")
	}
}

func queryAnimal(name, action string) {
	a := animals[name]

	if a == nil {
		fmt.Println("unknown name in query")
	} else {
		act(a, action)
	}
}

func act(a Animal, action string) {
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
