package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

/**
Write a program which prompts the user to first enter a name,
and then enter an address.

Your program should create a map and add the name and address to the map
using the keys "name" and "address", respectively.

Your program should use Marshal() to create a JSON object from the map,
and then your program should print the JSON object.
*/

func main2() {
	contacts := make(map[string]string)

	fmt.Println("Insert a name:")
	scn := bufio.NewScanner(os.Stdin)
	if !scn.Scan() {
		fmt.Println("an error occured while scanning name from user")
		return
	}

	contacts["name"] = scn.Text() //add name

	fmt.Println("Insert an address:")
	if !scn.Scan() {
		fmt.Println("an error occured while scanning address from user")
		return
	}

	contacts["address"] = scn.Text() //add address

	c, err := json.Marshal(contacts)
	if err != nil {
		fmt.Printf("an error occured while marshaling json: %d", err)
		return
	}
	fmt.Println(string(c))
}

func main() {
	bytes, err := json.Marshal(map[string]string{"name": inputName("Name"), "address": inputName("Address")})
	if err == nil {
		fmt.Println(string(bytes))
	} else {
		fmt.Printf("Error: %s\n", err.Error())
	}
}

func inputName(key string) string {
	var inputName string
	notSet := true
	for notSet {
		fmt.Printf("Input %s: ", key)
		lines, err := fmt.Scan(&inputName)
		if err == nil {
			if lines == 1 {
				notSet = false
			} else {
				fmt.Printf("Did not read one line, Read %d\n", lines)
			}
		} else {
			fmt.Printf("Error: %s\n", err.Error())
		}
	}
	return inputName
}
