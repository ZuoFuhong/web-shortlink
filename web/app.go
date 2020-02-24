package web

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"log"
	"net/http"
)

// App encapsulates Router
type App struct {
	Router      *mux.Router
	Middlewares *Middleware
}

// Initalize is initialization of app
func (a *App) Initalize() {
	a.Router = mux.NewRouter()
	a.Middlewares = &Middleware{}
	a.registerHandler()
}

func (a *App) registerHandler() {
	chain := alice.New(a.Middlewares.LoggingHandler, a.Middlewares.RecoverHandler)

	a.Router.Handle("/api/shorten", chain.ThenFunc(createShortUrl)).Methods("POST")
	a.Router.Handle("/api/info", chain.ThenFunc(getShortlinkInfo)).Methods("GET")
	a.Router.Handle("/{shortUrl:[a-zA-Z0-9]{1,11}}", chain.ThenFunc(redirect)).Methods("GET")
}

// Run starts listen and server
func (a *App) Run(addr string) {
	err := http.ListenAndServe(addr, a.Router)
	if err != nil {
		log.Printf("Server startup failed.")
		panic(err)
	}
}
