package main

import "fmt"

/*
This program is non-deterministic. Running the program multiple times will give you different outputs.
This happens because the main thread spawns multiple goroutines that potentially will access a variable
in unordered way producing then non-deterministic results.
*/
func main() {

	inc := 2
	numbers := []int{1, 1, 1, 1, 1}

	for _, n := range numbers {
		go func() {
			inc = incCounter(n) + inc
		}()
	}

	fmt.Println(inc)
}

func incCounter(n int) int {
	return n + 1
}
