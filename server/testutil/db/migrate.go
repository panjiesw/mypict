package db

import (
	"fmt"

	_ "github.com/mattes/migrate/driver/postgres"
	"github.com/mattes/migrate/migrate"
	"github.com/spf13/viper"
)

func Migrate(dir string) []error {

	if err := createDB(); err != nil {
		return []error{err}
	}

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/mypictdbtest?sslmode=disable",
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.port"))
	errors, ok := migrate.UpSync(url, fmt.Sprintf("%s/migrations", dir))
	if !ok {
		return errors
	}

	return nil
}

func createDB() error {
	c, err := createCon("")
	if err != nil {
		return err
	}
	defer c.Close()

	_, err = c.Exec(`CREATE DATABASE mypictdbtest`)
	if err != nil {
		return err
	}

	return nil
}
