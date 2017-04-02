package handler

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
	"gopkg.in/nullbio/null.v6"
	"panjiesw.com/mypict/server/model"
	"panjiesw.com/mypict/server/util/errs"
)

func (h *H) uploadRoutes() chi.Router {
	r := chi.NewRouter()
	r.Post("/images", h.UploadImages)
	r.Post("/gallery", h.UploadGallery)
	return r
}

func (h *H) UploadGallery(w http.ResponseWriter, r *http.Request) {
	ctx := GetContext(r)

	reader, err := r.MultipartReader()
	if err != nil {
		ctx.log.Error("Failed to read multipart request", "err", err)
		RenderAError(w, r, errs.ErrUnknown)
		return
	}

	gallery, err := h.uploader(reader, ctx)
	if err != nil {
		RenderError(w, r, err)
		return
	}

	if err := h.ds.GallerySave(gallery); err != nil {
		ctx.log.Error("Failed to save gallery data", "err", err)
		RenderError(w, r, err)
		return
	}
	render.JSON(w, r, gallery)
}

func (h *H) UploadImages(w http.ResponseWriter, r *http.Request) {
	ctx := GetContext(r)

	reader, err := r.MultipartReader()
	if err != nil {
		ctx.log.Error("Failed to read multipart request", "err", err)
		RenderAError(w, r, errs.ErrUnknown)
		return
	}

	gallery, err := h.uploader(reader, ctx)
	if err != nil {
		RenderError(w, r, err)
		return
	}

	if err := h.ds.ImageBSave(gallery.Images); err != nil {
		ctx.log.Error("Failed to save image data", "err", err)
		RenderError(w, r, err)
		return
	}
	render.JSON(w, r, map[string][]*model.ImageDTO{"data": gallery.Images})
}

func (h *H) uploader(reader *multipart.Reader, ctx *RootCtx) (*model.GalleryDTO, error) {
	imgs := []*model.ImageDTO{}
	gallery := &model.Gallery{UserID: ctx.UID()}

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		if part.FileName() == "" {
			if part.FormName() == "cp" {
				buff := h.bp.Get()
				if _, err := buff.ReadFrom(part); err != nil && err != io.EOF {
					ctx.log.Error("Failed to read cp", "err", err)
					//RenderAError(w, r, errs.ErrRequestBadParam)
					return nil, errs.ErrRequestBadParam
				}

				if cpi, err := strconv.Atoi(buff.String()); err != nil {
					ctx.log.Error("Failed to convert cp", "err", err)
					//RenderAError(w, r, errs.ErrRequestBadParam)
					return nil, errs.ErrRequestBadParam
				} else if cpi > 1 {
					return nil, errs.ErrRequestBadParam
				} else {
					gallery.ContentPolicy = uint8(cpi)
				}
				h.bp.Put(buff)
			} else if part.FormName() == "gtitle" {
				buff := h.bp.Get()
				if _, err := buff.ReadFrom(part); err != nil && err != io.EOF {
					ctx.log.Error("Failed to read gallery title", "err", err)
					return nil, errs.ErrRequestBadParam
				}

				var gt null.String
				gtf := buff.String()
				if gtf == "" {
					gt = null.NewString("", false)
				} else {
					gt = null.StringFrom(gtf)
				}
				gallery.Title = gt
				h.bp.Put(buff)
			}
			continue
		}

		iid, err := h.ds.ImageGenerateID()
		if err != nil {
			ctx.log.Error("Failed to generate image id")
			//RenderError(w, r, err)
			return nil, err
		}

		dst, err := os.Create(part.FileName())
		if err != nil {
			ctx.log.Error("Failed to create image file", "err", err)
			//RenderError(w, r, err)
			return nil, err
		}
		//noinspection GoDeferInLoop
		defer dst.Close()

		if _, err := io.Copy(dst, part); err != nil {
			ctx.log.Error("Failed to write image file", "err", err)
			//RenderError(w, r, err)
			return nil, err
		}

		imgs = append(imgs, &model.ImageDTO{Image: &model.Image{
			ID:            iid,
			Title:         null.StringFrom(part.FileName()),
			UserID:        gallery.UserID,
			ContentPolicy: gallery.ContentPolicy,
		}})
	}
	return &model.GalleryDTO{Gallery: gallery, Images: imgs}, nil
}
