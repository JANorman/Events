package conf

import (
    "os"
    "fmt"
)

var dsn string = ""

func GetRabbitMqDsn() (string) {
    if dsn == "" {
        port := os.Getenv("RABBIT_PORT_5672_TCP_PORT")
        host := os.Getenv("RABBIT_PORT_5672_TCP_ADDR")
        dsn = fmt.Sprintf("amqp://%s:%s/", host, port)
    }
    fmt.Println(dsn)
    return dsn
}