package main

func main() {
	var i int = 500
	var u uint = uint(i)
	var f float32 = float32(i)

	println(f, u)

	str := "ABCDEF"
	bytes := []byte(str)
	str2 := string(bytes)
	println(bytes, str2)
}
