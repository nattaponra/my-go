package concurrency

import (
	"fmt"
	"time"
)

func Pinger(c chan int) {
	for i := 0; ; i++ {
		c <- i
	}
}

func Printer(c chan int) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
