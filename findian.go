package main

import(
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	running := true
	scanner := bufio.NewScanner(os.Stdin)
	for running {
		fmt.Printf("Please enter a word: ")
		if scanner.Scan() {
			line := scanner.Text()
			fmt.Printf("Input is %s\n", line)
			a := strings.ToLower(line)
			b := strings.TrimSpace(a)
			if strings.HasPrefix(b, "i") && strings.Contains(b[1:len(b)-1], "a") && strings.HasSuffix(b, "n") {
				fmt.Println("Found!")
				running = false
			} else {
				fmt.Println("Not Found!")
			}

		}

		}

	}



