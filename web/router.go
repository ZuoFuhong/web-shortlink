package web

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"web-shortlink/app/interfaces"
)

type Router struct {
	*mux.Router
	ch *alice.Chain
}

func NewRouter() *Router {
	mw := NewMiddleware()
	chain := alice.New(mw.RequestMetricHandler)
	return &Router{
		Router: mux.NewRouter(),
		ch:     &chain,
	}
}

func (r *Router) registerHandler(shortService *interfaces.ShortLinkServiceImpl) {
	r.Handle("/ping", r.ch.ThenFunc(shortService.Ping)).Methods("GET")
	r.Handle("/api/v1/shorten", r.ch.ThenFunc(shortService.CreateShortUrl)).Methods("POST")
	r.Handle("/api/v1/short_info/{eid:[a-zA-Z0-9]{1,10}}", r.ch.ThenFunc(shortService.GetShortLinkInfo)).Methods("GET")
	r.Handle("/{eid:[a-zA-Z0-9]{1,10}}", r.ch.ThenFunc(shortService.Redirect)).Methods("GET")
}
