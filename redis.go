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

	// lua
	//redislua(rdb)

	cc := rdb.Set(ctx, "cc", "cc", time.Second*100)
	cc1 := rdb.Del(ctx, "cc1")
	cc2 := rdb.Del(ctx, "cc2")
	fmt.Println(cc, cc1, cc2)

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

/**
http://www.redis.cn/commands/eval.html
redis执行lua脚本
KEYS[i] 全局变量，对应Eval keys参数，i从1开始
ARGV[1] 局部变量，对应Eval args多参数，从1开始
redis.call()
tonumber() 数据类型转换
*/
func redislua(rdb *redis.Client) {
	arrStr := []string{"cc", "cc1"}
	RandomPushScript := `
	local res
	local i = tonumber(ARGV[1])
	while(i>0) do
		redis.call('SET',KEYS[i],KEYS[i])
		i=i-1
	end
	res = redis.call('SET',ARGV[2],222)
	return res
	`
	res := rdb.Eval(ctx, RandomPushScript, arrStr, len(arrStr), "cc2")

	fmt.Println(res.Val())
	fmt.Println(rdb.Get(ctx, "cc"))
	fmt.Println(rdb.Get(ctx, "cc1"))
	fmt.Println(rdb.Get(ctx, "cc2"))
}
