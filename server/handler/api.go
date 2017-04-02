package handler

import "github.com/pressly/chi"

func (h *H) apiRoutes() chi.Router {
	r := chi.NewRouter()
	r.Mount("/imgs", h.imageRoutes())
	return r
}
