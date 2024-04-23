package main

import "fmt"

func main() {
	myInts := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range myInts {
		if isEven(v) {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
}

func isEven(number int) bool {
	return number%2 == 0
}
