package main

import (
	"fmt"

	"github.com/nattaponra/my-go/concurrency"
)

func main() {

	//Ping Pong
	// var c chan string = make(chan string)
	// go concurrency.Pinger(c)
	// go concurrency.Ponger(c)
	// go concurrency.Printer(c)

	concurrency.Select()

	var input string
	fmt.Scanln(&input)
}
