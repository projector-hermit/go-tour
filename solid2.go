package main

import "fmt"

func main() {
	var x [5]float64
	x[0] = 92
	x[1] = 26
	x[2] = 77
	x[3] = 54
	x[4] = 87

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += x[i]
	}
	fmt.Println(total / 5)
}
