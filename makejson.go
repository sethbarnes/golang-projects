package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var name string
	var address string

	person := make(map[string]string)

	fmt.Print("Please, enter a name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name = scanner.Text()

	fmt.Printf("name is %s\n", name)

	person["name"] = name

	fmt.Print("Please, enter an address: ")
	scanner.Scan()
	address = scanner.Text()

	fmt.Printf("address is %s\n", address)

	person["address"] = address

	m, err := json.Marshal(person)
	if err != nil {
		fmt.Println("error:", err)
	}

	os.Stdout.Write(m)

}
