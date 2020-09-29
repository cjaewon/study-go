package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("숫자를 입력해주세요\n")
	reader := bufio.NewReader(os.Stdin)
	// os.Stdin 은 표준 입력

	line, _ := reader.ReadString('\n') // \n 이 올때까지 읽는다
	line = strings.TrimSpace(line)     // 개행문자 또는 공백 제거

	n1, _ := strconv.Atoi(line)

	line, _ = reader.ReadString('\n') // \n 이 올때까지 읽는다
	line = strings.TrimSpace(line)    // 개행문자 또는 공백 제거

	n2, _ := strconv.Atoi(line)

	fmt.Printf("입력하신 숫자는 %v, %v 입니다\n", n1, n2)
	fmt.Println("연산자를 입력해주세요")

	line, _ = reader.ReadString('\n') // \n 이 올때까지 읽는다
	line = strings.TrimSpace(line)

	switch line { // 다른 언어들과는 다르게 break 를 하지 않음, fallthrough 으로 다음으로 넘어감
	case "+":
		fmt.Printf("%v + %v = %v", n1, n2, n1+n2)
	case "-":
		fmt.Printf("%v - %v = %v", n1, n2, n1-n2)
	case "*":
		fmt.Printf("%v * %v = %v", n1, n2, n1*n2)
	case "/":
		fmt.Printf("%v / %v = %v", n1, n2, n1-n2)
	default:
		fmt.Print("잘못입력하셨습니다.")
	}

	if line == "+" {
	} else if line == "-" {
	} else if line == "*" {
	} else if line == "/" {
	} else {
	}
}
