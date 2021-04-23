package main

import (
	"fmt"
	"math"
)

func main() {
	var x = 69
	var a = math.Sqrt(float64(x))
	fmt.Println("Квадратный корень из числа 69 равен", math.Round(float64(a)))
	fmt.Printf("Если быть точным то %G.\n", math.Sqrt(float64(x)))
}
