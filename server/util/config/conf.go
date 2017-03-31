package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/mypict/")
	viper.AddConfigPath("$HOME/.mypict")
}

func Parse(p string) (*Conf, error) {
	if p != "" {
		viper.SetConfigFile(p)
	}

	_ = viper.ReadInConfig()

	var c Conf
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

func BindEnv(prefix string) {
	viper.SetEnvPrefix(prefix)

	viper.BindEnv("env")

	viper.BindEnv("http.host", fmt.Sprintf("%s_HTTP_HOST", prefix))
	viper.BindEnv("http.port", fmt.Sprintf("%s_HTTP_PORT", prefix))

	viper.BindEnv("database.name", fmt.Sprintf("%s_DB_NAME", prefix))
	viper.BindEnv("database.host", fmt.Sprintf("%s_DB_HOST", prefix))
	viper.BindEnv("database.port", fmt.Sprintf("%s_DB_PORT", prefix))
	viper.BindEnv("database.user", fmt.Sprintf("%s_DB_USER", prefix))
	viper.BindEnv("database.password", fmt.Sprintf("%s_DB_PASS", prefix))
	viper.BindEnv("database.pool.max_con", fmt.Sprintf("%s_DB_POOL_CON", prefix))

	viper.BindEnv("log.default", fmt.Sprintf("%s_LOG", prefix))
	viper.BindEnv("log.level.db", fmt.Sprintf("%s_DB_LOG", prefix))

	viper.SetDefault("env", "development")
	viper.SetDefault("log.default", "info")
	viper.SetDefault("http.host", "localhost")
	viper.SetDefault("http.port", 3000)
}

type Conf struct {
	Env      string `mapstructure:"env"`
	Database DB     `mapstructure:"database"`
	Http     Http   `mapstructure:"http"`
	Log      Log    `mapstructure:"log"`
}

type DB struct {
	Name     string `mapstructure:"name"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Pool     struct {
		MaxCon int `mapstructure:"max_con"`
	}
}

type Http struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type Log struct {
	Default string            `mapstructure:"default"`
	Level   map[string]string `mapstructure:"level"`
}

func (l Log) Lvl(module string) string {
	level, ok := l.Level[module]
	if ok {
		return level
	}
	return l.Default
}
