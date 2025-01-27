package main

import "fmt"

func main() {
	fmt.Print(DomainForLocale("site.com", "ua"))
}

func DomainForLocale(domain, locale string) string {
	if len(locale) == 0 {
		return "en" + "." + domain
	} else {
		return locale + "." + domain
	}
}
