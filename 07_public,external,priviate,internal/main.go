package main

import (
	"fmt"
	"strings"
)

//public vs private zoo 폴더 참고

// public, external, priviate, internal

func lenAndUpper(name string) (length int, upperacase string) {
	length = len(name)
	upperacase = strings.ToUpper(name)
	return
}

func main() {
	totalLength, uppercase := lenAndUpper("kim")
	fmt.Println(totalLength, uppercase)

}
