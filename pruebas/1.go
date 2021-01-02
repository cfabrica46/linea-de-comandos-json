package main

import (
	"fmt"
	"regexp"
)

func main() {

	match, _ := regexp.MatchString("[a-zA-Z0-9]+", "cesar2")

	fmt.Println(match)

	r, _ := regexp.Compile("[a-zA-Z0-9]+")

}
