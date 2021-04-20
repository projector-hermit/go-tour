package main

import (
	"fmt"
	"math/cmplx" // предоставляет основные константы и математические функции для комплексных чисел.
)

var (
	ToBe   bool       = true
	MaxInt uint64     = 20000 / 2 * 75
	z      complex128 = cmplx.Sqrt(69)
)

func main() {
	w := 7
	t := false
	f := " "
	fmt.Printf("Тип: %T Значение: %v\n", ToBe, ToBe)
	fmt.Printf("Тип: %T Значение: %v\n", MaxInt, MaxInt)
	fmt.Printf("Тип: %T Значение: %v\n", z, z)
	fmt.Printf("%v %v %q\n", w, t, f)
}
