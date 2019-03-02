package main

import (
	"fmt"

	"github.com/nattaponra/my-go/concurrency"
)

func main() {
	var c chan string = make(chan string)
	go concurrency.Pinger(c)
	go concurrency.Ponger(c)
	go concurrency.Printer(c)

	var input string
	fmt.Scanln(&input)
}
