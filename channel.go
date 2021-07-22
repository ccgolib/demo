package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"runtime"
)

type Job struct {
	Name string
}

func (j Job) AFun() {
	fmt.Println("isa")
}
var rctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//rdb.HSet(rctx, "channel", "a", 1)
	//rdb.Expire(rctx, "channel", time.Second * 10)
	fmt.Println(rdb.HGetAll(rctx, "channel").Val())

	//这里我们假设数据是int类型，缓存格式设为100
	dataChan := make(chan Job, 10)
	go func() {
		for {
			select {
			case data := <-dataChan:
				fmt.Println(data.Name)
				rdb.HSet(rctx, "channel", data.Name, 1)
				data.AFun()
				//time.Sleep(1 * time.Second) //这里延迟是模拟处理数据的耗时
			}
		}
	}()

	//填充数据
	for i := 0; i < 20; i++ {
		job := Job{
			Name: fmt.Sprintf("abc, %d", i),
		}
		dataChan <- job
	}

	//这里循环打印查看协程个数
	for {
		fmt.Println("runtime.NumGoroutine() :", runtime.NumGoroutine())
		//time.Sleep(2 * time.Second)
	}
}
