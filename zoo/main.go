package main

import (
	"fmt"

	animal "github.com/jcs1910/animals"
	animal_sounds "github.com/jcs1910/animals/animal_sounds"
)

func main() {
	fmt.Println(animal.Cat("What is animals sounds?"))
	fmt.Println(animal_sounds.DogSounds)
	fmt.Println(animal_sounds.CatSounds)
}
