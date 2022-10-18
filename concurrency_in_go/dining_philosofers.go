/*
Implement the dining philosopher's problem with the following constraints/modifications.
There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)

The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture). TODO review

In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
The host allows no more than 2 philosophers to eat concurrently.
Each philosopher is numbered, 1 through 5.
When a philosopher starts eating (after it has obtained necessary locks) it prints "starting to eat <number>" on a line by itself, where <number> is the number of the philosopher.
When a philosopher finishes eating (before it has released its locks) it prints "finishing eating <number>"" on a line by itself, where <number> is the number of the philosopher.
*/
package main

import (
	"fmt"
	"sync"
)

type philosopher struct {
	num             int
	rightCs, leftCs *chopStick //mutexes not not be coppied. Use pointers always.
}

type chopStick struct {
	mu sync.Mutex
}

func (p philosopher) eat(permission chan int, allow chan bool) {
	for i := 0; i < 3; i++ {
		p.leftCs.mu.Lock()
		p.rightCs.mu.Lock()

		fmt.Printf("starting to eat %d\n", p.num)

		//ask for permission to eat. (send philo number to host)
		permission <- p.num

		//wait for permission or denial
		select {
		case <-allow:
			fmt.Println(fmt.Sprintf("Philo #%d allowed to eat", p.num))
		}

		fmt.Printf("finishing eating %d\n", p.num)
		p.leftCs.mu.Unlock()
		p.rightCs.mu.Unlock()
	}
}

func main() {

	//init sticks
	CSticks := make([]*chopStick, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(chopStick)
	}
	//init philos
	philos := make([]*philosopher, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &philosopher{num: (i + 1), rightCs: CSticks[i], leftCs: CSticks[(i+1)%5]}
	}

	//init host
	askToEat := make(chan int, 2)
	allowToEat := make(chan bool)

	go func(perm chan int, allow chan bool) {
		for {
			select {
			case <-perm:
				allow <- true
			}
		}
	}(askToEat, allowToEat)

	for i := 0; i < 5; i++ {
		philos[i].eat(askToEat, allowToEat)
	}

}
