package main

import "fmt"

// Slice
func main() {
	// Slice
	rainbow := []string{"red", "orange", "yellow", "green", "blue", "indigo", "violet"}
	rainbow = append(rainbow, "sky", "pink", "purple")
	fmt.Println(rainbow)
}
