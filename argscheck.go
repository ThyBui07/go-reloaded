package main

import (
	"fmt"
	"os"
)

func argsCheck() []string {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("no")
		os.Exit(0)
	} else {
		fmt.Println("ok to process")
	}

	return args
}
