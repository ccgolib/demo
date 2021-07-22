package main

import (
	"fmt"
	"github.com/my/repo/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple(fmt.Sprintf("testmq%d", 1))
	rabbitmq.ConsumeSimple()

}