package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreetings(t *testing.T) {
	a := assert.New(t)
	a.Equal("Привіт, Іван!", Greetings(" іван"))
	a.Equal("Привіт, Петро!", Greetings("ПЕТРО"))
	a.Equal("Привіт, Василь!", Greetings("	вАсиЛЬ	"))
}
