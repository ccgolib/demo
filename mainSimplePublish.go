package main

import (
	"fmt"
	"github.com/my/repo/RabbitMQ"
)

func main() {

	for i:=1; i<20; i++ {
		rabbitmq := RabbitMQ.NewRabbitMQSimple(fmt.Sprintf("testmq%d", 1))
		rabbitmq.PublishSimple(fmt.Sprintf("Hello-%d", i))
		fmt.Println("发送成功")
	}


}