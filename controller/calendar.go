package controller

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/nickchou/gocode/app"
	//"github.com/nickchou/gocode/dataaccess"
	"github.com/nickchou/gocode/model"
)

//CalendarController 日历控制器
type CalendarController struct {
	app.App
}

//Get 提交方式
func (con *CalendarController) Get() {

	//计算代码耗时
	defer timeCost(time.Now())

	//产品ID，项目归属，渠道ID，出发地
	productid, belong, channel, depid := 0, 0, 0, 0
	//方式一
	var buffer bytes.Buffer
	//拿参数
	productid, _ = strconv.Atoi(con.R().FormValue("productid"))
	belong, _ = strconv.Atoi(con.R().FormValue("belong"))
	channel, _ = strconv.Atoi(con.R().FormValue("channel"))
	depid, _ = strconv.Atoi(con.R().FormValue("depid"))
	//new 一个实体
	pro := model.SGPProductBaseInfo{PBIOldProductId: productid, PBIBelong: 2}
	//DB:获取基本信息
	//dataaccess.GetProduct(&pro)

	//buffer
	buffer.WriteString(strconv.Itoa(productid) + "<br/>")
	buffer.WriteString(strconv.Itoa(belong) + "<br/>")
	buffer.WriteString(strconv.Itoa(channel) + "<br/>")
	buffer.WriteString(strconv.Itoa(depid) + "<br/>")

	buffer.WriteString("PBIId:" + strconv.FormatInt(pro.PBIId, 10) + "<br/>")

	//respnse
	con.W().Header().Set("content-type", "text/html; charset=utf-8")
	io.WriteString(con.W(), buffer.String())
	//方式二
	//fmt.Fprintf(con.W(), lunarcal.ToLunarDate("2009-02-02").Format("2006-01-02")+"<br/>")
}
func timeCost(start time.Time) {
	terminal := time.Since(start)
	fmt.Println(terminal)
}
