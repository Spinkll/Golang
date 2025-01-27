package main

import "fmt"

func main() {
	cat := Cat{}
	dog := Dog{}
	cow := Cow{}

	fmt.Println(cat.Voice())
	fmt.Println(dog.Voice())
	fmt.Println(cow.Voice())
}

type Voicer interface {
	Voice() string
}

type Cat struct {
}

type Cow struct {
}

type Dog struct {
}

func (c Cat) Voice() string {
	return "Мяу"
}

func (c Cow) Voice() string {
	return "Мууу"
}
func (d Dog) Voice() string {
	return "Гав"
}
