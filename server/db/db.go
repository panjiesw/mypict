package db

import (
	log "github.com/inconshreveable/log15"
	"github.com/jackc/pgx"
	"github.com/ventu-io/go-shortid"
	"panjiesw.com/mypict/server/config"
)

func Open(conf *config.Conf) (*Database, error) {
	pgxLvl, err := pgx.LogLevelFromString(conf.Log.Lvl("db"))
	if err != nil {
		pgxLvl = pgx.LogLevelWarn
	}

	lvl, err := log.LvlFromString(conf.Log.Lvl("db"))
	if err != nil {
		lvl = log.LvlWarn
	}

	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:      conf.Database.Host,
			Port:      uint16(conf.Database.Port),
			User:      conf.Database.User,
			Password:  conf.Database.Password,
			TLSConfig: nil,
			Database:  conf.Database.Name,
			Logger:    log.New("module", "pgx"),
			LogLevel:  pgxLvl,
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

	l := log.New("module", "db")
	hs := log.CallerStackHandler("%+v", log.StdoutHandler)
	hlf := log.LvlFilterHandler(lvl, hs)
	l.SetHandler(hlf)

	return &Database{
		log:  l,
		pool: pool,
		siid: siid,
		ssid: ssid,
		sgid: sgid,
	}, nil
}

type Database struct {
	log  log.Logger
	pool *pgx.ConnPool
	siid *shortid.Shortid
	ssid *shortid.Shortid
	sgid *shortid.Shortid
}

func (d *Database) Close() {
	d.pool.Close()
}
