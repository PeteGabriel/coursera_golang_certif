package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names.

Each line of the text file has a first name and a last name, in that order,
separated by a single space on the line.

Your program will define a name struct which has two fields,
fname for the first name and lname for the last name.
Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and
create a struct which contains the first and last names found in the file.

Each struct created will be added to a slice, and after all lines have been read from the file,
your program will have a slice containing one struct for each line in the file.

After reading all lines from the file,
your program should iterate through your slice of structs
and print the first and last names found in each struct.
*/

type Name struct {
	fname string
	lname string
}

const MAX_LEN = 20

func NewName(fn, ln string) Name {
	var fname, lname string
	if len(ln) > MAX_LEN {
		lname = ln[:MAX_LEN]
	} else {
		lname = ln
	}
	if len(fn) > MAX_LEN {
		fname = fn[:MAX_LEN]
	} else {
		fname = fn
	}
	return Name{
		fname: fname,
		lname: lname,
	}
}

func main() {
	var fn string

	fmt.Println("Enter file name:")
	fmt.Scanln(&fn)

	file, err := os.Open(fn)
	if err != nil {
		fmt.Printf("Could not open file with name %s\n", fn)
		return
	}
	defer file.Close()

	names := make([]Name, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content := strings.Split(scanner.Text(), " ")
		names = append(names, NewName(content[0], content[1]))
	}

	for _, name := range names {
		fmt.Printf("%s %s\n", name.fname, name.lname)
	}

}
