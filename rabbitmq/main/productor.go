package main

import (
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	//conn, err := amqp.Dial("amqp://bIndoMiddleWare:Bin$Mid!lePoIsAGh@35.194.205.200:5672")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:56721/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	body := bodyFrom(os.Args)
	body = `{"sid":7090,"oid":2700319,"n":"201809271411374640401600","et":1538030891,"paid_amount":"15.0"}`
	// 将消息发送到延时队列上
	err = ch.Publish(
		"",           // exchange 这里为空则不选择 exchange
		"test_delay", // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
			Expiration:  "5000", // 设置五秒的过期时间
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
