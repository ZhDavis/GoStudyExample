package main

import (
    "fmt"
    "log"
    "time"
    "bytes"
    "github.com/streadway/amqp"
)

func failOnError(err error, msg string){
    if err != nil {
        log.Fatalf("%s:%s", msg, err)
        panic(fmt.Sprintf("%s:%s", msg, err))
    }
}

func main() {

    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
    failOnError(err, "Failed to connect to server")
    defer conn.Close();

    ch, err := conn.Channel()
    failOnError(err, "Failed to connect to channel")
    defer ch.Close()

    q, err := ch.QueueDeclare(
        "hello",    //name
        false,      //durable
        false,      //delete when usused
        false,      // exclusive
        false,      //no-wait
        nil,        // arguments
    )
    failOnError(err, "Failed to declare a queue")

    msgs, err := ch.Consume(
        q.Name,     // queue
        "",         // consumer
        false,       // auto-ack
        false,      // exclusive
        false,      // no-local
        false,      // no-wait
        nil,        // arguments
    )
    failOnError(err, "Failed to register a consumer")

    forever := make(chan bool)
    go func(){
        for d:= range msgs{
            log.Printf("Received a message : %s", d.Body)
            dot_count := bytes.Count(d.Body, []byte("."))
            t := time.Duration(dot_count)
            time.Sleep(t* time.Second)
            log.Printf("Done")
           // d.Ack(false) 
        }
    }()

    log.Printf(" [*] Waiting for messages, To exit press CTRL+C")
    <-forever
    return
}
