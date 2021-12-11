package main

import "fmt"

const kinds = "Cat"

func main() {
	var (
		name     string  = "Nuko Grammer"
		gender   string  = "Male"
		year     int     = 8
		weight   float64 = 13.5
		is_alive bool    = true
	)

	fmt.Printf("KINDS:%s\n", kinds)
	fmt.Printf("GENDER:%s\n", gender)
	fmt.Printf("YEAR:%d\n", year)
	fmt.Printf("WEIGHT:%f\n", weight)
	fmt.Printf("NAME:%s\n", name)
	fmt.Printf("ALIVE:%t\n", is_alive)
}
