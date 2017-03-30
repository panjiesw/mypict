package db

import (
	"github.com/jackc/pgx"
	"gopkg.in/nullbio/null.v6"
	"panjiesw.com/mypict/server/errs"
	"panjiesw.com/mypict/server/model"
)

func (d *Database) GallerySave(g model.GalleryS, uid null.String) (*model.GalleryR, *errs.AError) {
	tx, err := d.pool.Begin()
	if err != nil {
		d.log.Error("Transaction not acquired", "err", err, "gallery", g, "uid", uid)
		return nil, errs.NewDB("Failed to acquire tx")
	}
	defer tx.Rollback()

	id, err := d.sgid.Generate()
	if err != nil {
		d.log.Error("Failed to generate gallery id", "err", err, "gallery", g, "uid", uid)
		return nil, errs.NewDB("Failed to generate gid")
	}

	var created string
	//noinspection SqlResolve
	query := `
	INSERT INTO gallery (id, title, uid, cp, created)
	VALUES ($1, $2, $3, $4, DEFAULT)
	RETURNING created`

	if err := tx.QueryRow(query, id, g.Title, uid, g.ContentPolicy).
		Scan(&created); err != nil {
		d.log.Error("Gallery insert error", "err", err, "gallery", g, "uid", uid)
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
		d.log.Error("Transaction failed to be committed", "err", err, "gallery", g, "uid", uid)
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

func (d *Database) GalleryByID(id string, g *model.GalleryR) *errs.AError {
	// noinspection SqlResolve
	query := `
	SELECT row_to_json(gr)
	FROM (
		SELECT id, title, uid, cp, created,
			(
				SELECT array_to_json(array_agg(row_to_json(img)))
				FROM (
					SELECT *
					FROM image i
					JOIN image_ids ii ON (i.id = ii.iid)
					WHERE ii.gid = g.id
				) as img
			) as images
		FROM gallery g
		WHERE g.id = $1
	) gr`

	if err := d.pool.QueryRow(query, id).Scan(g); err != nil {
		if err == pgx.ErrNoRows {
			return errs.ErrDBIDNotExists
		}
		d.log.Error("Failed to query gallery", "err", err, "id", id)
		return errs.ErrDBUnknown
	}

	return nil
}
