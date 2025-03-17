package main

import "fmt"

func main() {
	fmt.Println(IsValid(0, "hello world"))
	fmt.Println(IsValid(-22, "hello world"))
	fmt.Println(IsValid(22, ""))
	fmt.Println(IsValid(225, "hello world"))
}

func IsValid(id int, text string) bool {
	if id > 0 && len(text) != 0 {
		return true
	} else {
		return false
	}
}
