package db

import (
	"github.com/jackc/pgx"
	"github.com/mgutz/logxi"
	"github.com/ventu-io/go-shortid"
	"panjiesw.com/mypict/server/util/config"
)

func Open(conf *config.Conf) (*Database, error) {

	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:      conf.Database.Host,
			Port:      uint16(conf.Database.Port),
			User:      conf.Database.User,
			Password:  conf.Database.Password,
			TLSConfig: nil,
			Database:  conf.Database.Name,
			Logger:    &PgxLogger{log: logxi.New("pgx")},
			LogLevel:  pgx.LogLevelInfo,
		},
		MaxConnections: conf.Database.Pool.MaxCon,
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

	return &Database{
		l:    logxi.New("db"),
		pool: pool,
		siid: siid,
		ssid: ssid,
		sgid: sgid,
	}, nil
}

type Database struct {
	l    logxi.Logger
	pool *pgx.ConnPool
	siid *shortid.Shortid
	ssid *shortid.Shortid
	sgid *shortid.Shortid
}

func (d *Database) Close() {
	d.pool.Close()
}
