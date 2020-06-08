package web

import (
	"log"
	"net/http"
	"strconv"
	"web-shortlink/config"
)

type App struct {
	conf   *config.Conf
	router *router
}

func NewApp() *App {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	app := new(App)
	app.conf = config.LoadConf()
	app.router = NewRouter()
	return app
}

func (app *App) Run() {
	addr := app.conf.Server.Addr + ":" + strconv.Itoa(app.conf.Server.Port)
	log.Println("server runs on http://" + addr)
	app.router.RegisterHandler()
	err := http.ListenAndServe(addr, app.router.muxr)
	if err != nil {
		log.Printf("Server startup failed.")
		panic(err)
	}
}
