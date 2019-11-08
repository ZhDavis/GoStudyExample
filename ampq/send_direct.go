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
        log.Fatalf("%s: %s", msg, err)
        panic(fmt.Sprintf("%s: %s", msg, err))
    }
}
// 生产者
func main(){
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    ch, err := conn.Channel()
    failOnError(err, "Failed to open an channel")
    defer ch.Close()

    err = ch.ExchangeDeclare(
        "testdirect",     //name
        "direct",         //type
        true,    
        false,
        false,
        false,
        nil,
    )
    failOnError(err, "Failed to declare an exchange")

    body := bodyFrom(os.Args)
    err = ch.Publish(
        "testdirect",            // exchange
        severityFrom(os.Args),   // 通过routing key来发送到不同的queue中
        false,
        false,
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(body),
        })
    failOnError(err, "Failed to publish a message")

    log.Printf(" [x] sent %s", body)
}

func bodyFrom(args []string) string{
    var s string
    if(len(args) < 3) || os.Args[2] == "" {
        s = "hello"
    }else{
        s = strings.Join(args[2:], " ")
    }

    return s
}

func severityFrom(args []string) string {
    var s string
    if len(args) < 2 || args[1] == "" {
        s = "hello"
    }else {
        s = "other"
    }

    return s
}
