package db

import "github.com/spf13/viper"

func init() {
	viper.SetEnvPrefix("mpt")
	viper.BindEnv("database.name", "MPT_DB_NAME")
	viper.BindEnv("database.host", "MPT_DB_HOST")
	viper.BindEnv("database.port", "MPT_DB_PORT")
	viper.BindEnv("database.user", "MPT_DB_USER")
	viper.BindEnv("database.password", "MPT_DB_PASS")
	viper.BindEnv("database.pool.max_con", "MPT_DB_POOL_CON")
	viper.BindEnv("log.level.db", "MPT_DB_LOG")
}
