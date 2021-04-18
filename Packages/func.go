package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}
func main() {
	fmt.Println(add(60, 9))
	a, b := swap("вселенная", "бесконечна")
	fmt.Println(a, b)
}
