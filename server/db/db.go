package db

import (
	log "github.com/inconshreveable/log15"
	"github.com/jackc/pgx"
	"github.com/spf13/viper"
	"github.com/ventu-io/go-shortid"
)

func New() (*Database, error) {
	pgxLvl, err := pgx.LogLevelFromString(viper.GetString("log.level.db"))
	if err != nil {
		pgxLvl = pgx.LogLevelWarn
	}

	lvl, err := log.LvlFromString(viper.GetString("log.level.db"))
	if err != nil {
		lvl = log.LvlWarn
	}

	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:      viper.GetString("database.host"),
			Port:      uint16(viper.GetInt("database.port")),
			User:      viper.GetString("database.user"),
			Password:  viper.GetString("database.password"),
			TLSConfig: nil,
			Database:  viper.GetString("database.name"),
			Logger:    log.New("module", "pgx"),
			LogLevel:  pgxLvl,
		},
		MaxConnections: viper.GetInt("database.pool.max_con"),
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
		l:    l,
		pool: pool,
		siid: siid,
		ssid: ssid,
		sgid: sgid,
	}, nil
}

type Database struct {
	l    log.Logger
	pool *pgx.ConnPool
	siid *shortid.Shortid
	ssid *shortid.Shortid
	sgid *shortid.Shortid
}

func (d *Database) Close() {
	d.pool.Close()
}
