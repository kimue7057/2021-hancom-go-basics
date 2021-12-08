package main

import "fmt"

func main() {
	const firstname string = "Mary"
	fmt.Println(firstname)

	//var
	var firstname2 string = "Mary"
	firstname2 = "john"
	fmt.Println(firstname2)

	// 축약 문법 (축약문법은 오로지 함수 내에서만 사용가능)
	firstname3 := "mary" //var firstname2 string = "Mary" 와 같다.
	firstname3 = "john"
	age := 30
	fmt.Println(firstname3, age)

}
