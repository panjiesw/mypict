package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
	"panjiesw.com/mypict/server/handler"
)

func init() {
	viper.SetEnvPrefix("mp")

	viper.BindEnv("http.host", "MP_HTTP_HOST")
	viper.BindEnv("http.port", "MP_HTTP_PORT")

	viper.BindEnv("database.name", "MP_DB_NAME")
	viper.BindEnv("database.host", "MP_DB_HOST")
	viper.BindEnv("database.port", "MP_DB_PORT")
	viper.BindEnv("database.user", "MP_DB_USER")
	viper.BindEnv("database.password", "MP_DB_PASS")
	viper.BindEnv("database.pool.max_con", "MP_DB_POOL_CON")

	viper.BindEnv("log.level.db", "MP_DB_LOG")

	viper.SetDefault("http.host", "localhost")
	viper.SetDefault("http.port", 3000)
}

func main() {
	addr := fmt.Sprintf("%s:%d", viper.GetString("http.host"), viper.GetInt("http.port"))

	s := handler.New()
	http.ListenAndServe(addr, s)
}
