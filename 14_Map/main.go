package main

import "fmt"

func main() {
	myDog := map[string]string{"name": "doy", "gender": "male", "age": "3"}
	for key, value := range myDog {
		fmt.Println(key, value)
	}
}
