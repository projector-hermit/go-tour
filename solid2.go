package main

import "fmt"

func main() {
	x := [6]float64{
		99,
		55,
		75,
		98,
		25,
	}

	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}
