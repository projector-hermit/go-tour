package main

import "fmt"

func main() {
	x := make(map[string]int)
	y := make(map[int]int)
	x["key"] = 10
	x["key1"] = 99
	y[1] = 69
	y[2] = 88
	fmt.Println(x["key"])
	fmt.Println(x["key1"])
	fmt.Println(y[1])
	fmt.Println(y[2])
}
