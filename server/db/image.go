package db

import (
	"github.com/jackc/pgx"
	"gopkg.in/nullbio/null.v6"
	"panjiesw.com/mypict/server/errs"
	"panjiesw.com/mypict/server/model"
)

func (d *Database) imageBSave(tx *pgx.Tx, imgs []model.ImageS, uid null.String) *errs.AError {
	inputRows := [][]interface{}{}
	columns := []string{"id", "title", "uid", "cp"}

	for _, img := range imgs {
		inputRows = append(inputRows, []interface{}{img.ID, img.Title, uid, img.ContentPolicy})
	}

	n, err := tx.CopyFrom(model.ImageTableI, columns, pgx.CopyFromRows(inputRows))
	if err != nil {
		d.l.Error("Failed to save image", "error", err, "uid", uid, "imgs", imgs)
		return errs.NewDB("Failed to save image")
	} else if n != len(imgs) {
		d.l.Error("Saved image count not match", "want", len(imgs), "got", n, "uid", uid, "imgs", imgs)
		return errs.NewDB("Invalid saved image state")
	}
	return nil
}

func (d *Database) imageIDBSave(tx *pgx.Tx, imgs []model.ImageS, gid null.String) ([]model.ImageSR, *errs.AError) {
	inputRows := [][]interface{}{}
	columns := []string{"iid", "sid", "gid"}

	results := []model.ImageSR{}

	for _, img := range imgs {
		sid, err := d.ssid.Generate()
		if err != nil {
			d.l.Error("Failed to generate image sid", "err", err, "img", img, "gid", gid)
			return nil, errs.NewDB("Failed to generate sid")
		}
		inputRows = append(inputRows, []interface{}{img.ID, sid, gid})
		results = append(results,
			model.ImageSR{ID: img.ID, Title: img.Title, SID: null.StringFrom(sid)})
	}

	n, err := tx.CopyFrom(model.ImageIDTableI, columns, pgx.CopyFromRows(inputRows))
	if err != nil {
		d.l.Error("Bulk insert image ids failed", "err", err)
		return nil, errs.NewDB("Failed to save image ids")
	} else if n != len(imgs) {
		d.l.Error("Saved image ids count not match", "want", len(imgs), "got", n)
		return nil, errs.NewDB("Invalid saved image ids state")
	}

	return results, nil
}

func (d *Database) ImageBSave(imgs []model.ImageS, uid null.String) *errs.AError {
	tx, err := d.pool.Begin()
	if err != nil {
		d.l.Error("Transaction not acquired", "err", err, "uid", uid)
		return errs.NewDB("Failed to acquire tx")
	}
	defer tx.Rollback()

	if err := d.imageBSave(tx, imgs, uid); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		d.l.Error("Transaction failed to be committed", "err", err, "uid", uid)
		return errs.NewDB("Invalid saved image state")
	}

	return nil
}

func (d *Database) ImageByID(id string, img *model.ImageR) *errs.AError {
	//noinspection SqlResolve
	query := `
	SELECT row_to_json(img)
	FROM (
		SELECT
			img.id as id,
			img.title as title,
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
		d.l.Error("Failed to query image", "err", err, "id", id)
		return errs.ErrDBUnknown
	}

	return nil
}
