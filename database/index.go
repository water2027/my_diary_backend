package database

func InitDatabase() error {
	err := initMysqlDB()
	if err != nil {
		return err
	}
	return nil
}