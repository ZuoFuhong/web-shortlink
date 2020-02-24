package web

import (
	"log"
	"net/http"
	"time"
	"web-shortlink/defs"
)

type Middleware struct {
}

func (m Middleware) LoggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[%s] %q %v", r.Method, r.URL.String(), time.Now().Sub(startTime))
	}
	return http.HandlerFunc(fn)
}

func (m Middleware) RecoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recover from panic:%+v", err)
				sendErrorResponse(w, defs.ErrorInternalFaults)
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
