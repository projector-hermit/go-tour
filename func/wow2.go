package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println(i, "- кратно трем и пяти")
			} else {
				fmt.Println(i, "- кратно трем")
			}
		} else if i%5 == 0 {
			fmt.Println(i, "- кратно пяти")
		} else {
			fmt.Println(i)
		}
	}
}
