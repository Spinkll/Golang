package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(IntToString(42))
}

func IntToString(number int) string {
	return strconv.Itoa(number)
}
