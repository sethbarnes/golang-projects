package main

import (
	"fmt"
	"sort"
	"sync"
)

const maxSorters = 4

var arr[5] int

func sorters(id int, wg * sync.WaitGroup, ch chan int) {
	for val := range ch {
		fmt.Printf("Sorter #%d: recieved %d\n", id, val)
	}
	wg.Done()
}

func main() {
	ch := make(chan int, maxSorters)
	wg := sync.WaitGroup{}
	for i := 0; i < maxSorters; i++ {
		fmt.Scanf("%d", &ch)
		wg.Add(1)
		go sorters(i, &wg, ch)
	}
	sort.Ints(arr[:])

	var n int
	var err error
	inCount := 0
	for err == nil && inCount < 6 {
		fmt.Print("Please enter some integers: ")
		_, err = fmt.Scan(&n)
		if err == nil {
			inCount++
			ch <- n
		}
	}
	close(ch)
	wg.Wait()
	fmt.Println(ch)
}

