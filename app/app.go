package app

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
)

const (
	defController = "index"
	defMethod     = "Index"
)

type IApp interface {
	Init(ctx *Context)
	W() http.ResponseWriter
	R() *http.Request
	Display(tpls ...string)
	DisplayWithFuncs(funcs template.FuncMap, tpls ...string)
}
type App struct {
	ctx  *Context
	Data map[string]interface{}
}

func (a *App) Init(ctx *Context) {
	a.ctx = ctx
	a.Data = make(map[string]interface{})
}

func (a *App) W() http.ResponseWriter {
	return a.ctx.w
}

func (a *App) R() *http.Request {
	return a.ctx.r
}

func (a *App) Display(tpls ...string) {
	if len(tpls) == 0 {
		return
	}

	name := filepath.Base(tpls[0])

	t := template.Must(template.ParseFiles(tpls...))
	t.ExecuteTemplate(a.W(), name, a.Data)
}

func (a *App) DisplayWithFuncs(funcs template.FuncMap, tpls ...string) {
	if len(tpls) == 0 {
		return
	}

	name := filepath.Base(tpls[0])
	t := template.Must(template.New(name).Funcs(funcs).ParseFiles(tpls...))
	t.ExecuteTemplate(a.W(), name, a.Data)
}

var Static map[string]string = make(map[string]string)

func serveStatic(w http.ResponseWriter, r *http.Request) bool {
	for prefix, static := range Static {
		if strings.HasPrefix(r.URL.Path, prefix) {
			file := static + r.URL.Path[len(prefix):]
			http.ServeFile(w, r, file)
			return true
		}
	}

	return false
}
func RunOn(port string) {
	server := &http.Server{
		Handler: newHandler(),
		Addr:    port,
	}

	//fmt.Println("listen port" + port)
	log.Println("listen port" + port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("端口被占用，请检查。" + err.Error())
	}
}
func newHandler() *handler {
	h := &handler{}
	h.p.New = func() interface{} {
		return &Context{}
	}

	return h
}

type handler struct {
	p sync.Pool
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if serveStatic(w, r) {
		return
	}

	ctx := h.p.Get().(*Context)
	defer h.p.Put(ctx)
	ctx.Config(w, r)

	controllerName, methodName := h.findControllerInfo(r)
	controllerT, ok := mapping[controllerName]
	if !ok {
		http.NotFound(w, r)
		return
	}

	refV := reflect.New(controllerT)
	method := refV.MethodByName(methodName)
	if !method.IsValid() {
		http.NotFound(w, r)
		return
	}

	controller := refV.Interface().(IApp)
	controller.Init(ctx)
	method.Call(nil)
}
func (h *handler) findControllerInfo(r *http.Request) (string, string) {
	path := r.URL.Path
	if strings.HasSuffix(path, "/") {
		path = strings.TrimSuffix(path, "/")
	}
	pathInfo := strings.Split(path, "/")

	controllerName := defController
	if len(pathInfo) > 1 {
		controllerName = pathInfo[1]
	}

	methodName := defMethod
	if len(pathInfo) > 2 {
		methodName = strings.Title(strings.ToLower(pathInfo[2]))
	}

	return controllerName, methodName
}
