package main

import (
	"fmt"
	"sync"
)

const numPhilo = 5

var wg sync.WaitGroup

type chopS struct {sync.Mutex}

var host = make(chan int, 2)

type Philo struct {
	numbers 			int
	leftCS, rightCS *chopS
}

func (p Philo) eat() {

	for i := 0; i < 3; i++{

		host <- 0

		fmt.Println("Starting to eat", p.numbers)

		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("finishing eating", p.numbers)

		p.leftCS.Unlock()
		p.rightCS.Unlock()

		<- host
	}
	wg.Done()
}

func main() {
	cSticks := make([]*chopS, numPhilo)
	for i := 0; i < 5; i++ {
		cSticks[i] = new(chopS)
	}

	philos := make([]*Philo, numPhilo)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i, cSticks[i], cSticks[(i+1)%5]}
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat()
	}
	wg.Wait()
}