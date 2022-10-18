/*
Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing.
The BubbleSort() function should modify the slice so that the elements are in sorted
order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.
*/

package main

import "fmt"

func main() {

	var seq []int
	var input int

	for i := 0; i < 10; i++ {
		fmt.Println("Enter next integer:")
		fmt.Scanln(&input)
		seq = append(seq, input)
	}

	BubbleSort(seq)
	fmt.Println(seq)

}

//BubbleSort sorts in place
func BubbleSort(seq []int) {

	for idx := range seq {

		if (idx + 1) >= len(seq) {
			break
		}

		for idx2, val2 := range seq {

			if (idx2 + 1) >= len(seq) {
				break
			}
			if val2 > seq[idx2+1] {
				Swap(idx2, &seq)
			}
		}
	}
}

//Swap in place two adjacent elements in the slice 'values'
func Swap(idx int, values *[]int) {
	tmp := (*values)[idx]
	(*values)[idx] = (*values)[idx+1]
	(*values)[idx+1] = tmp
}
