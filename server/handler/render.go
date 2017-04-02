package handler

import (
	"net/http"

	"github.com/pressly/chi/render"
	"panjiesw.com/mypict/server/util/errs"
)

func RenderAError(w http.ResponseWriter, r *http.Request, err *errs.AError) {
	render.Status(r, err.Status)
	render.JSON(w, r, err)
}

func RenderError(w http.ResponseWriter, r *http.Request, err error) {
	switch e := err.(type) {
	case *errs.AError:
		if e != nil {
			RenderAError(w, r, e)
		}
	default:
		RenderAError(w, r, errs.New("server", e.Error(), http.StatusInternalServerError))
	}
}
