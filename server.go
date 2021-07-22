package main
// 导入两个包fmt和time
import (
	"fmt"
	"time"
)

// 定义一个say函数，接受一个字符串参数s
// 这个函数的作用就是循环打印5次字符串参数s
func say(s string) {
	// 循环5次
	for i := 0; i < 5; i++ {
		// 协程休眠100毫秒
		time.Sleep(100 * time.Millisecond)
		// 打印字符串s
		fmt.Println(s)
	}
}

// 入口函数main
func main() {
	// 通过go关键词，开启一个新的协程执行say函数
	go say("world")
	// main函数继续执行say函数，这里其实是在主协程执行
	say("hello")
}