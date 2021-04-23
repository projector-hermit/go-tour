package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "четное")
		} else {
			fmt.Println(i, "не четное")
		}
		switch i {
		case 0:
			fmt.Println(i, "ноль")
		case 1:
			fmt.Println("один")
		case 2:
			fmt.Println("два")
		case 3:
			fmt.Println("три")
		case 4:
			fmt.Println("четыре")
		case 5:
			fmt.Println("пять")
		default:
			fmt.Println("больше пяти")
		}
	}
	fmt.Println("wow!")
}
