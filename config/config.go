package config

import (
	"encoding/json"
	"os"
	"sync"
)

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type RedisConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
}

type DatabaseConfig struct {
	MysqlConfig `json:"mysql"`
	RedisConfig `json:"redis"`
}

type SMTPConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Config struct {
	DatabaseConfig `json:"database"`
	SMTPConfig     `json:"smtp"`
}

var (
	mu         sync.Mutex
	mysqlConfig MysqlConfig
	redisConfig RedisConfig
	smtpConfig SMTPConfig
)

func GetMysqlConfig() MysqlConfig {
	mu.Lock()
	defer mu.Unlock()
	return mysqlConfig
}

func GetRedisConfig() RedisConfig {
	mu.Lock()
	defer mu.Unlock()
	return redisConfig
}

func GetSMTPConfig() SMTPConfig {
	mu.Lock()
	defer mu.Unlock()
	return smtpConfig
}

func loadConfig() error {
	file, err := os.Open("config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var appConfig Config
	err = decoder.Decode(&appConfig)
	if err != nil {
		return err
	}
	dbConfig := appConfig.DatabaseConfig
	mysqlConfig = dbConfig.MysqlConfig
	redisConfig = dbConfig.RedisConfig
	smtpConfig = appConfig.SMTPConfig
	return nil
}

func InitConfig() error {
	err := loadConfig()
	return err
}
