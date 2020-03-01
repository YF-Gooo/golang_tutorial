package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 1000; i++ {
		go func(ch chan int, i int) {
			ch <- i
		}(ch, i)
	}
	go func(ch chan int) {
		for {
			select {
			case a := <-ch:
				fmt.Print(a)
			}
		}
	}(ch)
	time.Sleep(time.Second * 1)
}
