package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice.

The program should be written as a loop.

Before entering the loop, the program should create an empty integer slice of size (length) 3.

During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.

The slice must grow in size to accommodate any number of integers which the user decides to enter.
The program should only quit (exiting the loop) when the user enters the character 'X' instead of an integer.
*/

const END_CHAR = 'X'

func main2() {

	numbers := make([]int, 0, 3)
	var input string

	fmt.Println("Enter next integer:")
	fmt.Scanln(&input)

	for input != string(END_CHAR) {

		num, err := strconv.Atoi(input)
		if err != nil || input == "" {
			fmt.Println("Input must be a valid integer or termination char 'X'")
		} else {
			numbers = append(numbers, num)
			sort.Slice(numbers, func(ix, jx int) bool {
				return numbers[ix] < numbers[jx]
			})
			fmt.Printf("Sorted Sequence: %v - Len: %d - Cap: %d\n", numbers, len(numbers), cap(numbers))
		}

		fmt.Println("Enter next integer:")
		fmt.Scanln(&input)
	}
}

func main() {

	sli := make([]int, 3)
	i := 0
	for {
		var num int
		fmt.Println("Please enter an integer:")
		_, err := fmt.Scanln(&num)
		if err != nil {
			break
		}
		if i < 3 {
			sli[i] = num
			i++
		} else {
			sli = append(sli, num)
		}
		target := append([]int{}, sli...)
		sort.Ints(target)
		fmt.Println(target)
	}
}
