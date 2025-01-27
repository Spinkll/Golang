package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Greetings("ІВАН"))
}

func Greetings(name string) string {

	trimmed := strings.TrimSpace(name)
	lowered := strings.ToLower(trimmed)
	titled := strings.Title(lowered)

	return "Привіт, " + titled + "!"
}
