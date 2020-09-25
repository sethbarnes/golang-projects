package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	running := true
	scanner := bufio.NewScanner(os.Stdin)
	for running {
		fmt.Printf("Please enter a number: ")
		if scanner.Scan() {
			line := scanner.Text()
			f, err := strconv.ParseFloat(line, 64)
			if err == nil {
				i := int64(f)
				fmt.Println(i)
				running = false
			} else {
				fmt.Println("Input is not a valid number")
			}
		}
	}
}






