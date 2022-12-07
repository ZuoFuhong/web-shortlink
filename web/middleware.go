package web

import (
	"context"
	"github.com/google/uuid"
	"log"
	"net/http"
	"web-shortlink/consts"
)

type Middleware struct {
}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) RequestMetricHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Inject traceId to context
		traceId, _ := uuid.NewUUID()
		r = r.WithContext(context.WithValue(r.Context(), consts.TraceKey, traceId.String()))
		log.Printf("%s %s %q\n", traceId, r.Method, r.URL.String())
		w.Header().Add("Trace_id", traceId.String())
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
