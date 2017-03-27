package dbtest

import (
	"github.com/jackc/pgx"
	"github.com/spf13/viper"
)

func PgxConn(db string) (*pgx.Conn, error) {
	conf := pgx.ConnConfig{
		Host:      viper.GetString("database.host"),
		Port:      uint16(viper.GetInt("database.port")),
		User:      viper.GetString("database.user"),
		Password:  viper.GetString("database.password"),
		TLSConfig: nil,
	}

	if db != "" {
		conf.Database = db
	}

	c, err := pgx.Connect(conf)
	if err != nil {
		return nil, err
	}

	return c, nil
}
