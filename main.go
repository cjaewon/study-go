package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := x

	y[0] = 100

	fmt.Println(x) // [100, 2, 3, 4]
}
