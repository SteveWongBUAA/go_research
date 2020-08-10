package main

import (
	"fmt"
	"strings"
)

func AddUpper() func(int) (int, string) {
	n := 10
	str := "hello"
	return func(x int) (int, string) {
		n += x
		str += "asd"
		return n, str
	}
}

func main() {
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))

	f2 := addSuffix(".jpg")
	fmt.Println(f2("bird"))
	fmt.Println(f2("hasSuffix.jpg"))
}

func addSuffix(suffix string) func(string) string {
	return func(ori string) string {
		if strings.HasSuffix(ori, suffix) {
			return ori
		}
		return ori + suffix
	}
}
