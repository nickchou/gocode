package controller

import (
	//"html/template"

	"fmt"

	"github.com/nickchou/gocode/app"
)

type LoginController struct {
	app.App
}

func (i *LoginController) Login() {
	i.Data["name"] = "qibin"
	i.Data["email"] = "qibin0506@gmail.com"
	//i.Display("./view/info.tpl", "./view/header.tpl", "./view/footer.tpl")
	fmt.Fprintf(i.W(), "login")
	//i.DisplayWithFuncs(template.FuncMap{"look": funcs.Lookup}, "./view/info.tpl", "./view/header.tpl", "./view/footer.tpl")
}
func (i *LoginController) Logout() {
	i.Data["name"] = "qibin"
	i.Data["email"] = "qibin0506@gmail.com"
	//i.Display("./view/info.tpl", "./view/header.tpl", "./view/footer.tpl")
	fmt.Fprintf(i.W(), "logout")
	//i.DisplayWithFuncs(template.FuncMap{"look": funcs.Lookup}, "./view/info.tpl", "./view/header.tpl", "./view/footer.tpl")
}
