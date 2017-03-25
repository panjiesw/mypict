package db

import (
	"github.com/jackc/pgx"
	"gopkg.in/nullbio/null.v6"
	"panjiesw.com/mypict/server/errs"
	"panjiesw.com/mypict/server/model"
)

func (d *DB) imageBSave(tx *pgx.Tx, imgs []model.ImageS, uid null.String) *errs.AError {
	inputRows := [][]interface{}{}
	columns := []string{"id", "title", "uid", "cp"}

	for _, img := range imgs {
		inputRows = append(inputRows, []interface{}{img.ID, img.Title, uid, img.ContentPolicy})
	}

	n, err := tx.CopyFrom(model.ImageTableI, columns, pgx.CopyFromRows(inputRows))
	if err != nil {
		return errs.NewDB("Failed to save image")
	} else if n != len(imgs) {
		return errs.NewDB("Invalid saved image state")
	}
	return nil
}

func (d *DB) imageIDBSave(tx *pgx.Tx, imgs []model.ImageS, gid null.String) ([]model.ImageSR, *errs.AError) {
	inputRows := [][]interface{}{}
	columns := []string{"iid", "sid", "gid"}

	results := []model.ImageSR{}

	for _, img := range imgs {
		sid, err := d.ssid.Generate()
		if err != nil {
			return nil, errs.NewDB("Failed to generate sid")
		}
		inputRows = append(inputRows, []interface{}{img.ID, sid, gid})
		results = append(results,
			model.ImageSR{ID: img.ID, Title: img.Title, SID: null.StringFrom(sid)})
	}

	n, err := tx.CopyFrom(model.ImageIDTableI, columns, pgx.CopyFromRows(inputRows))
	if err != nil {
		return nil, errs.NewDB("Failed to save image ids")
	} else if n != len(imgs) {
		return nil, errs.NewDB("Invalid saved image ids state")
	}

	return results, nil
}

func (d *DB) ImageBSave(imgs []model.ImageS, uid null.String) *errs.AError {
	tx, err := d.pool.Begin()
	if err != nil {
		return errs.NewDB("Failed to acquire tx")
	}
	defer tx.Rollback()

	if err := d.imageBSave(tx, imgs, uid); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return errs.NewDB("Invalid saved image state")
	}

	return nil
}
