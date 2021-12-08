package main

import "fmt"

func main() {
	switch day := 5; day {
	case 0:
		fmt.Println("sunday")
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("tuesday")
	case 3:
		fmt.Println("wednesday")
	case 4:
		fmt.Println("thursday")
	case 5:
		fmt.Println("friday")
	case 6:
		fmt.Println("saturday")
	}
}

// func canIDrive(age int) bool {
// 	switch {
// 	case age < 20:
// 		return false
// 	case age > 80:
// 		return false
// 	case age > 20:
// 		return true
// 	}
// 	return false
// }
// func main() {
// 	fmt.Println(canIDrive(25))
//}
