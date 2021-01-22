package main

import (
	"fmt"
	"time"
)

var count int

func race() {
	fmt.Println(count)
	count++
}

func main() {

	count = 1
	go race()
	go race()
	time.Sleep(1 * time.Second)
}

//Race conditions occur when two functions with the same value race one another to be the first to finish.
//In this example each function has been assigned to the number "1" with an increment offset
//allowing for a race condition to be created.





