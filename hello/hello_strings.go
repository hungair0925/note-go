package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		firstname string = "Nuko"
		lastname  string = "Grammer"
		fullname  string
	)
	fmt.Printf("data type: %T\n", fullname)
	fmt.Println(fullname)

	firstname = strings.Replace(firstname, "u", "e", 1)
	firstname = strings.Replace(firstname, "Neko", "Pro", 1)
	lastname = strings.ToLower(lastname)

	fullname = firstname + lastname

	fmt.Println(fullname)
}
