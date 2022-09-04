package main

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(name string) string {
	data, err := os.ReadFile(name)
	check(err)
	stringInput := string(data)
	return stringInput

}

func modifyInput() {

}
