package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainForLocale(t *testing.T) {
	a := assert.New(t)
	a.Equal("en.goolang.com", DomainForLocale("goolang.com", ""))
	a.Equal("ua.goolang.com", DomainForLocale("goolang.com", "ua"))
	a.Equal("vn.goolang.com", DomainForLocale("goolang.com", "vn"))
}
