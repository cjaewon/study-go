package main

import "fmt"

func main() {
	s := "Hello 월드"
	// s2 := []rune(s)

	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]))
	}

	for t := range s {
		fmt.Print(string(t))
	}
}
