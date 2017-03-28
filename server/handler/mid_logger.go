package handler

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"time"
)

func (h *H) LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rctx := r.Context().Value(RootCtxKey).(*RootCtx)
		logger := NewRequestLogger(rctx.reqID)
		logger.Start(r)
		ww := NewWrapResponseWriter(w, r.ProtoMajor)

		t1 := time.Now()
		defer func() {
			t2 := time.Now()

			// Recover and record stack traces in case of a panic
			if rec := recover(); rec != nil {
				logger.Error("that happens", "panic", fmt.Sprintf("%+v", rec), "stack", string(debug.Stack()))
				http.Error(ww, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}

			// Log the entry, the request is complete.
			logger.End(ww.Status(), ww.BytesWritten(), t2.Sub(t1))
		}()

		rctx.log = logger
		next.ServeHTTP(ww, r)
	})
}
