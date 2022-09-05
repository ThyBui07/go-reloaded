package main

import (
	"strings"
)

func replaceWhiteSpace(s string) string {
	a := strings.ReplaceAll(s, " ", "")
	return a
}
