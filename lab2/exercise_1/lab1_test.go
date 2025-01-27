package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	a := assert.New(t)
	a.False(IsValid(0, ""))
	a.False(IsValid(5, ""))
	a.False(IsValid(-3, ""))
}
