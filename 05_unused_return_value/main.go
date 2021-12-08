package main

import "fmt"

func KoreanLegalDrinkingAge(age int) (int, string) {
	return age + 1, "This is Only for Koreans"
}

func main() {
	// 	age, words := KoreanLegalDrinkingAge(30)
	// 	fmt.Println((age, words))
	// }
	age, _ := KoreanLegalDrinkingAge(30)
	fmt.Println(age)
}

func add(a int, b int) int {
	return a + b
}

func add2(a, b int) int {
	return a + b
}

func add3(a, b int) string {
	return "Kis" + "young-hee"
}

// func main() {
// 	fmt.Println(add3(10, 10))

// }
