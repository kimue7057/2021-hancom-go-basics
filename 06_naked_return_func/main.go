package main

import "fmt"

//naked return func
func split(sum int) (x, y int) {
	x = sum * 5 / 4
	y = sum - x
	return
}

func main() {
	x, y := split(100)
	fmt.Println(x, y)
}

func favoriteColor(color ...string) {
	fmt.Println(color)
}

// func main() {
// 	favoriteColor("yellow", "blue", "red", "dark --grey", "black")
// }
