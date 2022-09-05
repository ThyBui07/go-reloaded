package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func findNumOfInstance(reg string, s string) (int, string) {
	var num int
	var instance string
	regCompile := regexp.MustCompile(reg)
	arrayOfInstance := regCompile.FindAllStringSubmatch(s, -1)

	for _, e := range arrayOfInstance {
		instance = e[1]
		if e[3] != "" {
			var err error
			num, err = strconv.Atoi(e[3])
			if err != nil {
				fmt.Print(err)
			}
		} else {
			num = 0
		}
	}

	return num, instance
}
