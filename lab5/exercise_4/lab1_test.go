package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWork(t *testing.T) {
	expected := []string{
		"Goroutine 0: counter = 0",
		"Goroutine 0: counter = 1",
		"Goroutine 0: counter = 2",
		"Goroutine 0: counter = 3",
		"Goroutine 0: counter = 4",
		"Goroutine 0: counter = 5",
	}

	results := work(0)

	assert.Equal(t, expected, results, "Work function did not produce expected output")
}
