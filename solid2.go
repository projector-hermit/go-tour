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
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))
}
