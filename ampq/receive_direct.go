package main

import(
    "fmt"
    "log"
    "os"

    "github.com/streadway/amqp"
)

func failOnError(err error, msg string){
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
        panic(fmt.Sprintf("%s: %s", msg, err))
    }
}
// 消费者
func main(){
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

    err = ch.ExchangeDeclare(
        "testdirect",
        "direct",
        true,
        false, 
        false,
        false,
        nil,
    )
    failOnError(err, "Failed to declare an exchange")

    q, err := ch.QueueDeclare(
        "",     //name
        false,
        false,
        true,
        false,
        nil,
    )
    failOnError(err, "Failed to declare a queue")

    for _, s := range os.Args[1:] {
        log.Printf("Binding queue %s to exchange %s with routing key %s",
        q.Name, "testdirect", s)
        err = ch.QueueBind(
            q.Name, //queue name
            s,                //routing key
            "testdirect",    //exchange
            false,
            nil,
        )
        failOnError(err, "Failed to bind a queue")
    }

    msgs, err := ch.Consume(
        q.Name,     //name
        "",                //consumer
        true,
        false,
        false,
        false,
        nil,
    )
    failOnError(err, "Failed to register a consumer")

    forever := make(chan bool)

    go func(){
        for d:= range msgs{
            log.Printf(" [x] %s", d.Body)
        }
    }()

    log.Printf(" [*] Waiting for logs. To Exit press Ctrl+c")
    <-forever
}
