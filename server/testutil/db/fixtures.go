package db

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/viper"
)

func Fixtures(dir string) error {

	c, err := createCon(viper.GetString("database.name"))
	defer c.Close()

	// seed images
	imgb, err := ioutil.ReadFile(fmt.Sprintf("%s/fixtures/image.sql", dir))
	if err != nil {
		return err
	}
	if _, err := c.Exec(string(imgb)); err != nil {
		return err
	}

	// seed galleries
	gb, err := ioutil.ReadFile(fmt.Sprintf("%s/fixtures/gallery.sql", dir))
	if err != nil {
		return err
	}
	if _, err := c.Exec(string(gb)); err != nil {
		return err
	}

	// seed image_ids
	imgib, err := ioutil.ReadFile(fmt.Sprintf("%s/fixtures/image_ids.sql", dir))
	if err != nil {
		return err
	}
	if _, err := c.Exec(string(imgib)); err != nil {
		return err
	}

	return nil
}
