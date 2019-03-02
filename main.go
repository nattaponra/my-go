package main

import (
	"github.com/nattaponra/my-go/concurrency"
)

func main() {

	//Ping Pong
	// var c chan string = make(chan string)
	// go concurrency.Pinger(c)
	// go concurrency.Ponger(c)
	// go concurrency.Printer(c)

	//Select
	//concurrency.Select()

	//Bugffer
	concurrency.FindMaximumSize()
	// var input string
	// fmt.Scanln(&input)
}
