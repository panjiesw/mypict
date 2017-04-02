package db

import (
	"github.com/jackc/pgx"
	"panjiesw.com/mypict/server/model"
	"panjiesw.com/mypict/server/util/errs"
)

func (d *Database) imageBSave(tx *pgx.Tx, imgs []*model.ImageDTO) error {
	inputRows := [][]interface{}{}
	columns := []string{"id", "title", "filename", "uid", "cp"}

	for _, img := range imgs {
		inputRows = append(inputRows, []interface{}{
			img.Image.ID,
			img.Image.Title,
			img.Image.FileName,
			img.Image.UserID,
			img.Image.ContentPolicy,
		})
	}

	n, err := tx.CopyFrom(model.ImageTableI, columns, pgx.CopyFromRows(inputRows))
	if err != nil {
		d.l.Error("Failed to save image", "err", err, "imgs", imgs)
		return errs.NewDB("Failed to save image")
	} else if n != len(imgs) {
		d.l.Error("Saved image count not match", "want", len(imgs), "got", n, "imgs", imgs)
		return errs.NewDB("Invalid saved image state")
	}
	return nil
}

func (d *Database) ImageBSave(imgs []*model.ImageDTO) error {
	tx, err := d.pool.Begin()
	if err != nil {
		d.l.Error("Transaction not acquired", "err", err, "imgs", imgs)
		return errs.NewDB("Failed to acquire tx")
	}
	defer tx.Rollback()

	if err := d.imageBSave(tx, imgs); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		d.l.Error("Transaction failed to be committed", "err", err, "imgs", imgs)
		return errs.NewDB("Invalid saved image state")
	}

	return nil
}

func (d *Database) ImageByID(id string, img *model.ImageDTO) error {
	//noinspection SqlResolve
	query := `
	SELECT row_to_json(img)
	FROM (
		SELECT
			img.id as id,
			img.title as title,
			img.filename as filename,
			img.uid as uid,
			ids.sid as sid,
			ids.gid as gid,
			g.title as gtitle,
			img.cp as cp,
			img.created as created
		FROM image img
		LEFT JOIN image_ids ids ON (img.id = ids.iid)
		LEFT JOIN gallery g ON (ids.gid = g.id)
		WHERE
				img.id = $1 OR
				(ids.iid IS NOT NULL AND (ids.iid = $1 OR ids.sid = $1))
	) img`

	if err := d.pool.QueryRow(query, id).Scan(img); err != nil {
		if err == pgx.ErrNoRows {
			return errs.ErrDBIDNotExists
		}
		d.l.Error("Failed to query image", "err", err, "id", id)
		return errs.ErrDBUnknown
	}

	return nil
}

func (d *Database) ImageGenerateID() (string, error) {
	s, err := d.siid.Generate()
	if err != nil {
		d.l.Error("Failed to generate image id", "err", err)
		return "", errs.NewDB("Failed to generate id")
	}
	return s, nil
}
