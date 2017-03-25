package db

import (
	"github.com/jackc/pgx"
	"github.com/ventu-io/go-shortid"
)

type DB struct {
	pool *pgx.ConnPool
	siid *shortid.Shortid
	ssid *shortid.Shortid
	sgid *shortid.Shortid
}

func New() (*DB, error) {
	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		MaxConnections: 10,
	})
	if err != nil {
		return nil, err
	}

	siid, err := shortid.New(1, shortid.DEFAULT_ABC, 10)
	if err != nil {
		return nil, err
	}

	ssid, err := shortid.New(2, shortid.DEFAULT_ABC, 10)
	if err != nil {
		return nil, err
	}

	sgid, err := shortid.New(3, shortid.DEFAULT_ABC, 10)
	if err != nil {
		return nil, err
	}

	return &DB{pool: pool, siid: siid, ssid: ssid, sgid: sgid}, nil
}
