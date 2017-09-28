package controller

import (
	"bytes"
	"fmt"
	"io"

	"github.com/nickchou/gocode/app"
	"github.com/nickchou/gocode/comm"
)

type CalController struct {
	app.App
}

func (con *CalController) Index() {

	//方式一
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("阳转阴历:%v -> %v <br/>", "2001-02-02",
		lunarcal.ToLunarDate("2001-02-02").Format("2006-01-02")))
	buffer.WriteString(fmt.Sprintf("阳转阴历:%v -> %v <br/>", "2009-02-02",
		lunarcal.ToLunarDate("2009-02-02").Format("2006-01-02")))
	buffer.WriteString(fmt.Sprintf("阳转阴历:%v -> %v <br/>", "2010-02-02",
		lunarcal.ToLunarDate("2010-02-02").Format("2006-01-02")))
	buffer.WriteString(fmt.Sprintf("阳转阴历:%v -> %v <br/>", "2017-02-02",
		lunarcal.ToLunarDate("2017-02-02").Format("2006-01-02")))
	buffer.WriteString(fmt.Sprintf("阳转阴历:%v -> %v <br/>", "2017-09-26",
		lunarcal.ToLunarDate("2017-09-26").Format("2006-01-02")))
	//respnse
	con.W().Header().Set("content-type", "text/html; charset=utf-8")
	io.WriteString(con.W(), buffer.String())
	//方式二
	//fmt.Fprintf(con.W(), lunarcal.ToLunarDate("2009-02-02").Format("2006-01-02")+"<br/>")

}
