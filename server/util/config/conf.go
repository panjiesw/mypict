package config

import (
	"fmt"

	"os"

	"github.com/spf13/viper"
	"gopkg.in/nullbio/null.v6"
)

var PS = string(os.PathSeparator)

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

	viper.BindEnv("picture.dir", fmt.Sprintf("%s_PIC_DIR", prefix))

	viper.SetDefault("env", "development")
	viper.SetDefault("log.default", "info")
	viper.SetDefault("http.host", "localhost")
	viper.SetDefault("http.port", 3000)
	viper.SetDefault("picture.dir", fmt.Sprintf(".%spictures", string(os.PathSeparator)))
}

type Conf struct {
	Env      string  `mapstructure:"env"`
	Database DB      `mapstructure:"database"`
	Http     Http    `mapstructure:"http"`
	Log      Log     `mapstructure:"log"`
	Picture  Picture `mapstructure:"picture"`
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

type Picture struct {
	Dir string `mapstructure:"dir"`
}

func (p Picture) PictLoc(pictDir, file string) string {
	return fmt.Sprintf("%s%s%s", pictDir, PS, file)
}

func (p Picture) ThumbLoc(thumbDir, file string, size int) string {
	return fmt.Sprintf("%s%sth_%d_%s", thumbDir, PS, size, file)
}

func (p Picture) PictDir(user null.String, id string) string {
	if user.Valid {
		return fmt.Sprintf("%s%susers%s%s%s%s", p.Dir, PS, PS, user.String, PS, id)
	}
	return fmt.Sprintf("%s%sunknown%s%s", p.Dir, PS, PS, id)
}

func (p Picture) ThumbDir(pictDir string) string {
	return fmt.Sprintf("%s%sthumb", pictDir, PS)
}
