package db

func Cleanup() error {
	c, err := createCon("")
	if err != nil {
		return err
	}
	defer c.Close()

	if _, err := c.Exec("DROP DATABASE mypictdbtest"); err != nil {
		return err
	}

	return nil
}
