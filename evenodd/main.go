package main

import (
	"fmt"
)

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, n := range numbers {
		checkOddEven(n)
	}
	fmt.Println("================================================")
	for i := 0; i < 20; i++ {
		checkOddEven(i)
	}

}

func checkOddEven(n int) {
	if n%2 == 0 {
		fmt.Printf("%v is even\n", n)
	} else {
		fmt.Printf("%v is odd\n", n)
	}
}
