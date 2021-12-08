package main

import "fmt"

func add(numbers ...int) {
	for index, number := range numbers {
		fmt.Println(index, number)
	}
}

func main() {
	add(1, 2, 3, 4, 5)
}
