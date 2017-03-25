package db

import (
	"time"

	"gopkg.in/nullbio/null.v6"
	"panjiesw.com/mypict/server/errs"
	"panjiesw.com/mypict/server/model"
)

func (d *DB) GallerySave(g model.GalleryS, uid null.String) (*model.GalleryR, *errs.AError) {
	tx, err := d.pool.Begin()
	if err != nil {
		return nil, errs.NewDB("Failed to acquire tx")
	}
	defer tx.Rollback()

	id, err := d.sgid.Generate()
	if err != nil {
		return nil, errs.NewDB("Failed to generate gid")
	}

	var created time.Time

	if err := tx.QueryRow(
		`INSERT INTO gallery
		(id, title, uid, cp, created)
				VALUES
		($1, $2, $3, $4, DEFAULT)
		RETURNING created`, id, g.Title, uid, g.ContentPolicy).Scan(&created); err != nil {
		return nil, errs.NewDB("Failed to save gallery")
	}

	if err := d.imageBSave(tx, g.Images, uid); err != nil {
		return nil, err
	}

	imgs, errr := d.imageIDBSave(tx, g.Images, null.StringFrom(id))
	if errr != nil {
		return nil, errr
	}

	if err := tx.Commit(); err != nil {
		return nil, errs.NewDB("Invalid saved gallery state")
	}

	return &model.GalleryR{
		ID:            id,
		Title:         g.Title,
		UserID:        uid,
		ContentPolicy: g.ContentPolicy,
		Images:        imgs,
		CreatedAt:     created,
	}, nil
}
