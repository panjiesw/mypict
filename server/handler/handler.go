package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pressly/chi"
	"go.uber.org/zap"
	"panjiesw.com/mypict/server/db"
	"panjiesw.com/mypict/server/util/config"
)

func New(c *config.Conf, ds db.Datastore, z *zap.SugaredLogger) *H {
	r := chi.NewRouter()

	h := &H{Mux: r, z: z, ds: ds, c: c}
	h.initialize()
	return h
}

type H struct {
	*chi.Mux
	z  *zap.SugaredLogger
	ds db.Datastore
	c  *config.Conf
}

func (h *H) AddRootCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), RootCtxKey, &RootCtx{})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *H) Start() {
	addr := fmt.Sprintf("%s:%d", h.c.Http.Host, h.c.Http.Port)
	h.z.Infof("Server listening on %s", addr)
	http.ListenAndServe(addr, h)
}

func (h *H) initialize() {
	h.Use(h.AddRootCtx)
	h.Use(h.RequestID)
	h.Use(h.LoggerMiddleware)
	h.Mount("/_", h.apiRoutes())
}
