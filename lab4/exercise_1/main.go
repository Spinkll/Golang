package main

import (
	"fmt"
	"strings"
)

func main() {
	u := UserCreateRequest{FirstName: "", Age: 0}
	fmt.Println(Validate(u))
}

/*func Validate(req UserCreateRequest) string {
	if strings.Contains(req.FirstName, " ") {
		return "Ім'я не повинно містити пробілів"
	}
	if strings.TrimSpace(req.FirstName) == "" {
		return "Ім’я не може бути пустим"
	}
	if req.Age < 0 {
		return "Вік має бути більше нуля"
	}
	return ""

}*/

func Validate(req UserCreateRequest) string {
	if strings.TrimSpace(req.FirstName) == "" {
		return "invalid request"
	}

	if strings.Contains(req.FirstName, " ") {
		return "invalid request"
	}

	if req.Age <= 0 || req.Age > 150 {
		return "invalid request"
	}

	return ""
}

type UserCreateRequest struct {
	FirstName string
	Age       int
}
