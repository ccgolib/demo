package main

import (
	"fmt"
	"time"
)

func main() {
	//ch := make(chan int, 100)
	//start := time.Now()
	//for i := 0; i < 20; i++ {
	//	ch <- i
	//}
	//close(ch)
	//for v := range ch{
	//	fmt.Println(v)
	//}
	//end := time.Now()
	//fmt.Println(end.Sub(start))

	// goroutine有时候会进入阻塞情况，那么如何避免由于channel阻塞导致整个程序阻塞的发生那？解决方案：通过select设置超时处理，具体程序如下
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case i := <-c:
				fmt.Println(i)
			case <-time.After(time.Duration(3) * time.Second):    //设置超时时间为３ｓ，如果channel　3s钟没有响应，一直阻塞，则报告超时，进行超时处理．
				fmt.Println("timeout")
				o <- true
				break
			}
		}
	}()
	<-o
}