package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Бросаю для вас кубики... Значение \"0\" означает, что кубик укатился под диван =)")
	a := rand.Intn(6)
	b := rand.Intn(6)
	fmt.Println("На первом кубике выпало", a)
	fmt.Println("На втором кубике выпало", b)
	fmt.Println("Всего получается", a+b)
}
