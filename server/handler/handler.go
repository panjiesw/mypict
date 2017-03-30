package handler

import (
	"context"
	"net/http"

	log "github.com/inconshreveable/log15"
	"github.com/pressly/chi"
	"panjiesw.com/mypict/server/config"
	"panjiesw.com/mypict/server/db"
)

type H struct {
	*chi.Mux
	log log.Logger
	DS  db.Datastore
}

func New(c *config.Conf) *H {
	d, err := db.Open(c)
	if err != nil {
		panic(err)
	}

	l := log.New("module", "server")

	r := chi.NewRouter()

	h := &H{Mux: r, log: l, DS: d}
	h.initialize()
	return h
}

func (h *H) initialize() {
	h.Use(h.AddRootCtx)
	h.Use(h.RequestID)
	h.Use(h.LoggerMiddleware)
	h.Mount("/_", h.apiRoutes())
}

func (h *H) AddRootCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), RootCtxKey, &RootCtx{})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
