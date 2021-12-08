package main

import "fmt"

//Method (메서드) => func과 비슷한 모양을 띈다, 함수명 앞에 타입과 함수명이 붙어있는 것을 리시버(receiver)라고 한다.
type Number struct {
	A, B int
}

//method (1. Value receiver) => Number라고 정의된 struct에 아무것도 붙어 있지 않는 것을 밸루 리시버라고 한다

func (n Number) printNumber() {
	fmt.Println(n)
}

//메서드 (pointer receiver) => n은 실질적인 값
func (n *Number) addPointer(_number int) {
	n.A += _number
	n.B += _number
}

//메서드 (value receiver) => n 자체가 복제된 값을 의미한다.
//즉 완전히 다른값을 의미하기 때문에 n을 출력을 하더라도 해당 부분의 결과 값이 반영되지 않는다.
func (n Number) addValue(_number int) {
	n.A += _number
	n.B += _number
}

func main() {
	n := Number{10, 20}
	// n.printNumber()
	// fmt.Println(n.A)
	// fmt.Println(n.B)
	// fmt.Println(n.c)
	// fmt.Println(n.d)
	n.addPointer(100)
	n.addValue(200)
	fmt.Println(n)
}
