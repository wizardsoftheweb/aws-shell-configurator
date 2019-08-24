package main

import (
	"fmt"

	"github.com/prometheus/common/log"
)

func main() {
	fmt.Println("rad")
}

func fatalCheck(err error) {
	log.Fatal(err)
}
