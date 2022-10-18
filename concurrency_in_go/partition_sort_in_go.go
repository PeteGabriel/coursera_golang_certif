package main

import (
	"fmt"
	"sort"
	"sync"
)

/*
import (
	"fmt"
	"sort"
	"sync"
)


Write a program to sort an array of integers.
The program should partition the array into 4 parts,
each of which is sorted by a different goroutine.

Each partition should be of approximately equal size.

Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers.
Each goroutine which sorts ¼ of the array should print the subarray that it will sort.

When sorting is complete, the main goroutine should print the entire sorted list.
*/

func main() {

	var elements []int
	var n int
	for i := 0; i < 10; i++ {
		fmt.Print("> ")
		fmt.Scan(&n)
		elements = append(elements, n)
	}

	var wg sync.WaitGroup
	ps := len(elements) / 4 //partition size
	ch := make(chan []int, ps*4)

	for i := 0; i < len(elements); i = i + ps {
		wg.Add(1)
		go sortPartition(elements[i:i+ps], &wg, ch)
	}

	wg.Wait()
	close(ch)
	var elemSorted []int
	for sorted := range ch {
		elemSorted = append(elemSorted, sorted...)
	}
	sort.Ints(elemSorted)

	fmt.Println(elemSorted)
}

func sortPartition(a []int, wg *sync.WaitGroup, ch chan []int) {
	defer wg.Done()
	fmt.Println(a)
	sort.Ints(a)
	ch <- a
}
