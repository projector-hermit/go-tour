package main

import "fmt"

func main() {
	var x [6]float64
	x[0] = 99
	x[1] = 98
	x[2] = 77
	x[3] = 97
	x[4] = 87
	x[5] = 5

	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}
