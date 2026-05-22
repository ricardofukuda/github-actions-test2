package main

import (
	"app/math"
	"fmt"
	"os"
)

func main() {
	println("hello world!")

	fmt.Printf("MESSAGE: '%s'\r\n", os.Getenv("MESSAGE"))

	a := 1.0
	b := 2.0
	fmt.Printf("%f + %f = %f\r\n", a, b, math.Sum(a, b))
}
