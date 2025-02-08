package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"

	"my_diary/config"
)

var (
	db   *sql.DB
	once sync.Once
	mu   sync.Mutex
)

func GetMysqlDB() *sql.DB {
	mu.Lock()
	defer mu.Unlock()
	return db
}

func initTable() error {
	userSql := `
CREATE TABLE IF NOT EXISTS users (
    userId INT AUTO_INCREMENT PRIMARY KEY,
    userEmail VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);
	`
	diarySql := `
	CREATE TABLE IF NOT EXISTS diaries (
    diaryId INT NOT NULL AUTO_INCREMENT,
    userId INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    createdAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    isPublic BOOLEAN DEFAULT FALSE,
    PRIMARY KEY (diaryId)
);
`
	_, err := db.Exec(userSql)
	if err != nil {
		return err
	}
	_, err = db.Exec(diarySql)
	if err != nil {
		return err
	}
	return nil
}

func initMysqlDB() error {
	var err error
	once.Do(func() {
		var mysqlConfig = config.GetDatabaseConfig()
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Name)
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			return
		}

		err = db.Ping()
		if err != nil {
			return
		}
		err = initTable()
		if err != nil {
			return
		}
		fmt.Println("成功连接到MySQL数据库")
	})
	return err
}
