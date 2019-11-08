package main

import (
    "fmt"
    "log"
    "os"
    "strings"
    "github.com/streadway/amqp"
)

func failOnError(err error, msg string){
    if err != nil {
        log.Fatalf("%s:%s", msg, err)
        panic(fmt.Sprintf("%s:%s", msg, err))
    }
}

func bodyFrom(args []string) string {
    var s string
    if (len(args) < 2) || os.Args[1] == "" {
        s = "hoge"
    } else {
        s = strings.Join(args[1:], " ")
    }

    return s
}

func main() {
    // step1: 作为生产者与amqp server 建立一个连接，rabbitmq的提供的默认端口是5672
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    // step2: 在这个连接上面创建channel
    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

    // step3: 创建一个名字叫做hello的queue
    q, err := ch.QueueDeclare(
    "hellodurable", //name
    true,  //durable
    false,  //delete when unused
    false,  //exclusive
    false,  //no wait
    nil,    //arguments
    )
    failOnError(err, "Failed to declare q queue")

    err = ch.Qos(
            1,     // prefetch count
            0,     // prefetch size
            false, // global
    )
    failOnError(err, "Failed to set QoS")
    // step4: 向rabbitmq server发送"Hello"消息
    body := bodyFrom(os.Args)
    err = ch.Publish(
        "",     //exchange
        q.Name,     // routing key
        false,  //mandatory
        false, //immediate
        amqp.Publishing{
            ContentType: "text/plain",
            Body :      []byte(body),
        }) 
    failOnError(err, "Failed to publish a message")
    log.Printf(" [x] Sent %s", body)

    return
}
