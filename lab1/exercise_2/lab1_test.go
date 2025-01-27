package main

import (
	"fmt"
	"log"
	"os/exec"
	"testing"
)

func TestVariables(t *testing.T) {
	cmd := exec.Command("go", "run", "./main.go")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("error on running the command: %v/n", err)
	}
	fmt.Println(string(out))

}
