package main

func main() {
	println(ErrorMessageToCode("OK"))
	println(ErrorMessageToCode("CANCELLED"))
	println(ErrorMessageToCode("UNKNOWN"))
	println(ErrorMessageToCode("INVALID"))
}

func ErrorMessageToCode(msg string) int {
	switch msg {
	case "OK":
		return 0
	case "CANCELLED":
		return 1
	case "UNKNOWN":
		return 2
	default:
		return 2
	}
}
