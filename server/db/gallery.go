package db

import (
	"github.com/jackc/pgx"
	"panjiesw.com/mypict/server/model"
	"panjiesw.com/mypict/server/util/errs"
)

func (d *Database) imageIDBSave(tx *pgx.Tx, imgs []*model.ImageDTO, gid string) error {
	inputRows := [][]interface{}{}
	columns := []string{"iid", "sid", "gid"}

	for _, img := range imgs {
		sid, err := d.ssid.Generate()
		if err != nil {
			d.l.Error("Failed to generate image sid", "err", err, "img", img, "gid", gid)
			return errs.NewDB("Failed to generate sid")
		}
		img.SID = sid
		img.GID = gid
		inputRows = append(inputRows, []interface{}{img.Image.ID, img.SID, img.GID})
	}

	n, err := tx.CopyFrom(model.ImageIDTableI, columns, pgx.CopyFromRows(inputRows))
	if err != nil {
		d.l.Error("Bulk insert image ids failed", "err", err)
		return errs.NewDB("Failed to save image ids")
	} else if n != len(imgs) {
		d.l.Error("Saved image ids count not match", "want", len(imgs), "got", n)
		return errs.NewDB("Invalid saved image ids state")
	}

	return nil
}

func (d *Database) GallerySave(g *model.GalleryDTO) error {
	tx, err := d.pool.Begin()
	if err != nil {
		d.l.Error("Transaction not acquired", "err", err, "gallery", g)
		return errs.NewDB("Failed to acquire tx")
	}
	defer tx.Rollback()

	id, err := d.sgid.Generate()
	if err != nil {
		d.l.Error("Failed to generate gallery id", "err", err, "gallery", g)
		return errs.NewDB("Failed to generate gid")
	}
	g.Gallery.ID = id

	//var created string
	//noinspection SqlResolve
	query := `
	INSERT INTO gallery (id, title, uid, cp, created)
	VALUES ($1, $2, $3, $4, DEFAULT)
	RETURNING created`

	if err := tx.QueryRow(query, g.Gallery.ID, g.Gallery.Title, g.Gallery.UserID, g.Gallery.ContentPolicy).
		Scan(&g.Gallery.CreatedAt); err != nil {
		d.l.Error("Gallery insert error", "err", err, "gallery", g)
		return errs.NewDB("Failed to save gallery")
	}

	if err := d.imageBSave(tx, g.Images); err != nil {
		return err
	}

	if err := d.imageIDBSave(tx, g.Images, g.Gallery.ID); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		d.l.Error("Transaction failed to be committed", "err", err, "gallery", g)
		return errs.NewDB("Invalid saved gallery state")
	}

	return nil
}

func (d *Database) GalleryByID(id string, g *model.GalleryDTO) *errs.AError {
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
		d.l.Error("Failed to query gallery", "err", err, "id", id)
		return errs.ErrDBUnknown
	}

	return nil
}
