package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	channelNums := 3
	ch := make(chan int, channelNums)
	defer close(ch)

	go func() {
		for {
			select {
			case a := <-ch:
				fmt.Println(a)
			default:
				fmt.Println("no data")
				return
			}
		}
	}()

	for i := 0; i < channelNums; i++ {
		ch <- i
	}

	end := time.Now()
	fmt.Println(end.Sub(start))
	//time.Sleep(2 * time.Second)
}
