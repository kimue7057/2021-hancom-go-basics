package main

import "fmt"

func main() {
	fmt.Println("Hello, ", "I'm Kim 25 years old")

	fmt.Print("hELLO ~ ", "I'm Jack ") //인자마다 띄어쓰기를 하지않고
	//마지막에 줄바꿈을 하지않는다
	fmt.Printf("%d years old!", 50) // 형식 지정자를 사용할 수 있다.
}
