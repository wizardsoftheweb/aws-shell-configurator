package main

import (
	"fmt"
)

func main() {
	fmt.Println("rad")
}

func nilErrorOrPanic(err error) {
	panic(err)
}
