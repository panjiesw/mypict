package db

import (
	"github.com/jackc/pgx"
	"github.com/ventu-io/go-shortid"
	"go.uber.org/zap"
	"panjiesw.com/mypict/server/util/config"
	"panjiesw.com/mypict/server/util/logging"
)

func Open(conf *config.Conf, z *zap.SugaredLogger) (*Database, error) {
	pgxLvl, err := pgx.LogLevelFromString(conf.Log.Lvl("db"))
	if err != nil {
		pgxLvl = pgx.LogLevelWarn
	}

	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:      conf.Database.Host,
			Port:      uint16(conf.Database.Port),
			User:      conf.Database.User,
			Password:  conf.Database.Password,
			TLSConfig: nil,
			Database:  conf.Database.Name,
			Logger:    logging.NewPGXLog(z),
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

	return &Database{
		z:    z,
		pool: pool,
		siid: siid,
		ssid: ssid,
		sgid: sgid,
	}, nil
}

type Database struct {
	z    *zap.SugaredLogger
	pool *pgx.ConnPool
	siid *shortid.Shortid
	ssid *shortid.Shortid
	sgid *shortid.Shortid
}

func (d *Database) Close() {
	d.pool.Close()
}
