package controller

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/nickchou/gocode/app"
)

type CommController struct {
	app.App
}

func (con *CommController) Quxian() {
	var buffer bytes.Buffer
	con.W().Header().Set("content-type", "text/html; charset=utf-8")
	//goquery http
	doc, err := goquery.NewDocument("http://studygolang.com/topics")
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".topics .topic").Each(func(i int, contentSelection *goquery.Selection) {
		title := contentSelection.Find(".title a").Text()
		buffer.WriteString(fmt.Sprintf("第%v个帖子标题：%v<br/>", i+1, title))
	})
	io.WriteString(con.W(), buffer.String())
}
