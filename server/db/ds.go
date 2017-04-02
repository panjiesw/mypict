package db

import "panjiesw.com/mypict/server/model"

type Datastore interface {
	ImageBSave(imgs []*model.ImageDTO) error
	ImageByID(id string, img *model.ImageDTO) error
	ImageGenerateID() (string, error)

	GallerySave(g *model.GalleryDTO) error
}
