package main

import (
	"fmt"
	"os"
)

func main() {
	println("hello world!")

	fmt.Printf("MESSAGE: '%s'\r\n", os.Getenv("MESSAGE"))
}
