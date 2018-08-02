package main

import (
	"github.com/nickchou/gocode/app"
	"github.com/nickchou/gocode/controller"
)

func main() {
	app.Static["/js"] = "static/js"
	app.AutoRouter(&controller.IndexController{})
	app.AutoRouter(&controller.CalController{})
	app.AutoRouter(&controller.CommController{})
	app.AutoRouter(&controller.CalendarController{})
	app.AutoRouter(&controller.ZhuzherController{})
	app.Router("login", &controller.LoginController{})
	app.RunOn(":9090")
}
