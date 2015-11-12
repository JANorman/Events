package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type EventController struct {
	beego.Controller
}

type TransactionReport struct {
  success bool
  name string
}

func (r TransactionReport) Name() string {
    return r.name
}

func (c *EventController) Get() {
	report := TransactionReport{success: true, name: "Bob"}
    fmt.Println(report)
    c.Data["json"] = &report
    c.ServeJson()
}
