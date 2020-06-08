package web

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"web-shortlink/web/handler"
)

type router struct {
	muxr *mux.Router
}

func NewRouter() *router {
	var router = new(router)
	router.muxr = mux.NewRouter()
	return router
}

func (r *router) RegisterHandler() {
	chain := alice.New(Middleware.LoggingHandler, Middleware.RecoverPanic, Middleware.CORSHandler)
	r.muxr.Handle("/api/shorten", chain.ThenFunc(handler.OpenAPI.CreateShortUrl)).Methods("POST")
	r.muxr.Handle("/api/info", chain.ThenFunc(handler.OpenAPI.GetShortlinkInfo)).Methods("GET").Queries("url", "{url}")
	r.muxr.Handle("/{url:[a-zA-Z0-9]{1,11}}", chain.ThenFunc(handler.OpenAPI.Redirect)).Methods("GET")
}
