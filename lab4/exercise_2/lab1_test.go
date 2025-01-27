package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumWorker(t *testing.T) {
	a := assert.New(t)

	cat := Cat{}
	dog := Dog{}
	cow := Cow{}

	a.Equal("Мяу", cat.Voice())

	a.Equal("Гав", dog.Voice())

	a.Equal("Мууу", cow.Voice())
}
