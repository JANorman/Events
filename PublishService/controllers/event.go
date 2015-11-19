package controllers

import (
	"github.com/astaxie/beego"
	"github.com/streadway/amqp"
	"log"
	"time"
	"app/conf"
)

type EventController struct {
	beego.Controller
}

type TransactionSuccessReport struct {
  Published bool `json:"published"`
  PublishedAt time.Time `json:"published_at"`
}

type TransactionFailureReport struct {
  Published bool   `json:"published"`
  Message   string `json:"message"`
}

var queueName string = "events"

func ConnectToRabbitMq() (*amqp.Channel, error) {
    // Connects opens an AMQP connection from the credentials in the URL.
    conn, err := amqp.Dial(conf.GetRabbitMqDsn())

    if err != nil {
        log.Fatalf("connection.open: %s", err)
    }

    defer conn.Close()

    c, err := conn.Channel()
    if err != nil {
        log.Fatalf("channel.open: %s", err)
    }

    err = c.ExchangeDeclare(queueName, "topic", true, false, false, false, nil)
    if err != nil {
        log.Fatalf("exchange.declare: %v", err)
    }

    return c, err
}

func (controller *EventController) Post() {

    c, err := ConnectToRabbitMq();

    // TODO: Do something on error

    timestamp := time.Now()
    msg := amqp.Publishing{
        DeliveryMode: amqp.Persistent,
        Timestamp:    timestamp,
        ContentType:  "text/plain",
        Body:         []byte("Go Go AMQP!"),
    }

    // This is not a mandatory delivery, so it will be dropped if there are no
    // queues bound to the logs exchange.
    err = c.Publish(queueName, "info", true, false, msg)
    if err != nil {
        // Since publish is asynchronous this can happen if the network connection
        // is reset or if the server has run out of resources.
        log.Fatalf("basic.publish: %v", err)

    }
    report := TransactionSuccessReport{Published: false, PublishedAt: time.Time{}}
    RespondWithReport(controller, report)
    return
	//report := TransactionReport{Published: true, PublishedAt: timestamp}
    //RespondWithReport(controller, report)
}



func RespondWithReport(controller *EventController, report TransactionSuccessReport) {
    controller.Data["json"] = &report
    controller.ServeJson()
}