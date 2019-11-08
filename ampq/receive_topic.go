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

func main(){
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

    err = ch.ExchangeDeclare(
        "testtopic",
        "topic",
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

    if len(os.Args) < 2 {
        log.Printf("Usage: %s [info] [warning] [error]", os.Args[0])
        os.Exit(0)
    }

    for _, s := range os.Args[1:] {
        log.Printf("Binding queue %s to exchange %s with routing key %s",
            q.Name, "testtopic", s)
        err = ch.QueueBind(
            q.Name, 
            s,
            "testtopic",
            false,
            nil)
        failOnError(err, "Failed to bind a queue")
    }

    msgs, err := ch.Consume(
        q.Name,
        "",
        true,
        false,
        false, 
        false,
        nil,
    )
    failOnError(err, "Failed to register a consumer")

    forever := make(chan bool)

    go func(){
        for d:= range msgs {
            log.Printf(" [x] %s", d.Body)
        }
    }()

    log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
    <-forever
}
