package db

import (
	"gopkg.in/nullbio/null.v6"
	"panjiesw.com/mypict/server/errs"
	"panjiesw.com/mypict/server/model"
)

type Datastore interface {
	ImageBSave(imgs []model.ImageS, uid null.String) *errs.AError
	ImageByID(id string, img *model.ImageR) *errs.AError
	ImageGenerateID() (null.String, *errs.AError)

	GallerySave(g model.GalleryS, uid null.String) (*model.GalleryR, *errs.AError)
}
