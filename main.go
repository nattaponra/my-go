package main

import (
	"fmt"

	"github.com/nattaponra/my-go/concurrency"
)

func main() {
	var c chan int = make(chan int)
	go concurrency.Pinger(c)
	go concurrency.Printer(c)

	var input string
	fmt.Scanln(&input)
}
