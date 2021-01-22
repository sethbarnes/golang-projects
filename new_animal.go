package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Speak() string
	Eat() string
	Move() string
}

type Cow struct{}

func (c Cow) Speak() string {
	return "moo"
}

func (c Cow) Eat() string {
	return "grass"
}

func (c Cow) Move() string {
	return "walk"
}

type Bird struct{}

func (b Bird) Speak() string {
	return "peep"
}

func (b Bird) Eat() string {
	return "worms"
}

func (b Bird) Move() string {
	return "fly"
}

type Snake struct{}

func (s Snake) Speak() string {
	return "hsss"
}

func (s Snake) Eat() string {
	return "mice"
}

func (s Snake) Move() string {
	return "slither"
}

func newAnimal(typ string) Animal {
	var animal Animal = nil

	switch typ {
	case "bird":
		animal = Bird{}
	case "cow":
		animal = Cow{}
	case "snake":
		animal = Snake{}
	}

	return animal
}

func main() {
	m := make(map[string]Animal)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		str, err := reader.ReadString('\n')
		if err == nil {
			cmds := strings.Split(str, " ")
			if len(cmds) == 3 {
				action := strings.ToLower(cmds[0])
				name := strings.ToLower(cmds[1])
				attOrType := strings.ToLower(strings.TrimRight(cmds[2], "\n"))

				switch action {
				case "newanimal":
					animal := newAnimal(attOrType)
					if animal == nil {
						fmt.Printf("Invalid animal type: '%s'\n", attOrType)
					} else {
						_, exists := m[name]
						if exists {
							fmt.Printf("Animal name already exists: '%s'\n", name)
						} else {
							m[name] = animal
							fmt.Println("Created It!")
						}
					}
				case "query":
					animal, exists := m[name]
					if exists {
						switch attOrType {
						case "speak":
							fmt.Println(animal.Speak())
						case "eat":
							fmt.Println(animal.Eat())
						case "move":
							fmt.Println(animal.Move())
						default:
							fmt.Printf("Invalid animal query: '%s'\n", attOrType)
						}
					} else {
						fmt.Printf("Animal to be queried does not exist: '%s'\n", name)
					}
				default:
					fmt.Printf("Invalid animal action: '%s'\n", action)
				}
			} else {
				fmt.Printf("Invalid input: '%s'\n", str)
			}
		} else {
			panic(err.Error())
		}
	}
}