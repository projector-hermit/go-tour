package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}
func split(sum int) (x, y int) {
	x = sum * 1 / 2
	y = sum - x
	return
}

func main() {
	var (
		peace, door, ball bool
		t, j              int = 1, 2
		i                 int
		c, python, java   = true, false, "no!"
		a, b              = swap("вселенная", "бесконечна")
	)
	fmt.Println(i, peace, door, ball)
	fmt.Println(t, j, c, python, java)
	fmt.Println(add(60, 9))
	fmt.Println(a, b)
	fmt.Println(split(176))

}
