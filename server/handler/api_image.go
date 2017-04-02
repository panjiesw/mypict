package handler

import (
	"net/http"

	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
	"panjiesw.com/mypict/server/model"
	"panjiesw.com/mypict/server/util/errs"
)

func (h *H) imageRoutes() chi.Router {
	r := chi.NewRouter()
	r.Get("/:id", h.ImageMetaByID)
	return r
}

func (h *H) ImageMetaByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		RenderAError(w, r, errs.ErrRequestBadParam)
		return
	}

	var img model.ImageDTO
	if err := h.ds.ImageByID(id, &img); err != nil {
		RenderError(w, r, err)
	} else {
		render.JSON(w, r, img)
	}
}
