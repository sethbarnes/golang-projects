package main

import (
	"fmt"
	"os"
	"sort"
	"sync"
)

const maxSorters = 4

func sorters(id int, wg * sync.WaitGroup, ch chan int, arr *[]int) {
	for val := range ch {
		*arr = append(*arr, val)
	}
	sort.Ints(*arr)
	wg.Done()
}

func main() {
	arrs := make([][]int, maxSorters)
	ch := make(chan int, maxSorters)
	wg := sync.WaitGroup{}
	for i := 0; i < maxSorters; i++ {
		wg.Add(1)
		go sorters(i, &wg, ch, &arrs[i])
	}

	var n int
	var err error
	inCount := 0
	for err == nil && inCount <= n {
		fmt.Fprintf(os.Stdout, "Please enter an integer (enter 'q' to quit): ")
		_, err = fmt.Scan(&n)
		if err == nil {
			inCount++
			ch <- n
		}
	}
	close(ch)
	wg.Wait()
	sortArr := []int{}
	for _, arr := range arrs {
		sortArr = append(sortArr, arr...)
		sort.Ints(sortArr)
	}
	fmt.Println("Sorted integers: ", sortArr)
}

