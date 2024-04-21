package main

import "fmt"

func main() {
	colors := map[string]int{
		"hello": 4,
		"bye":   5,
		"what":  6,
	}

	delete(colors, "what")
	printMap(colors)
}

func printMap(c map[string]int) {
	for k, v := range c {
		fmt.Printf("Key: %s Value: %d\n", k, v)
	}
}
