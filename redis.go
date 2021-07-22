package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	t := time.Now()

	rdb.Set(ctx, "cc1", 1, time.Second*100)
	cc := rdb.Del(ctx, "cc1")
	cc2 := rdb.Del(ctx, "cc2")
	fmt.Println(cc, cc2)
	//for i:=0; i< 100; i++ {
	//	fmt.Println(i)
	//}
	//go func() {
	//	rdb.Set(ctx, "cc1", "c112", 0)
	//
	//	fmt.Println(rdb.Get(ctx, "cc1"))
	//}()

	endt := time.Since(t)
	fmt.Println(endt)

	/*// hset,hget  map数组使用
	rdb.HSet(ctx,"myhash", "key1", "123")
	rdb.HSet(ctx,"myhash", "key2", "234")
	rdb.HSet(ctx,"myhash", "key3", "345")
	rdb.Expire(ctx,"myhash", time.Minute)
	fmt.Println(rdb.HGetAll(ctx, "myhash").Result())


	// 管道-命令打包
	pipe := rdb.Pipeline()
	incr := pipe.Incr(ctx, "pipe_test")
	pipe.Set(ctx, "a", 1, time.Minute)
	pipe.Set(ctx, "b", 2, time.Minute)
	pipe.Set(ctx, "c", 3, time.Minute)
	pipe.Expire(ctx, "pipe_test", time.Hour)

	_, err := pipe.Exec(ctx)

	fmt.Println(incr.Val(), err)
	fmt.Println(rdb.Get(ctx, "pipe_test"))
	fmt.Println(rdb.Get(ctx, "a").Val())
	fmt.Println(rdb.Get(ctx, "b"))
	fmt.Println(rdb.Get(ctx, "c"))*/

}
