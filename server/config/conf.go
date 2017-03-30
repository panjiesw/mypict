package config

import "github.com/spf13/viper"

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

type Conf struct {
	Database DB   `mapstructure:"database"`
	Http     Http `mapstructure:"http"`
	Log      Log  `mapstructure:"log"`
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
