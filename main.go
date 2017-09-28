package main

import (
	//	"fmt"
	//	"reflect"
	//	"time"

	"github.com/nickchou/gocode/app"
	//"github.com/nickchou/gocode/comm"
	"github.com/nickchou/gocode/controller"
)

func main() {
	app.Static["/js"] = "static/js"
	app.AutoRouter(&controller.IndexController{})
	app.AutoRouter(&controller.CalController{})
	app.Router("login", &controller.LoginController{})
	app.RunOn(":9090")

	//	t := lunarcal.ToLunarDate("2017-09-26")
	//	fmt.Println(reflect.TypeOf(t))
	//	n := time.Now()
	//	fmt.Println(reflect.TypeOf(n))
	//	fmt.Println(n.Format("2006-01-02"))
	//	fmt.Println(lunarcal.ToLunarDate("2017-09-26"))
	//	fmt.Println(lunarcal.ToLunarStr("2009-02-02"))
	//	fmt.Println(lunarcal.ToLunarStr("2010-02-02"))
	//	fmt.Println(lunarcal.ToLunarStr("2017-02-02"))
	//	fmt.Println(lunarcal.ToLunarStr("2017-09-26"))
}
