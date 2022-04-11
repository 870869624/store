package main

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"jinghaijun.com/store/constants"
)

type DBconnectConfig struct {
	Host     string
	Db       string
	User     string
	Password string
	Port     string
}

func (c *DBconnectConfig) connect() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Db)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		return nil
	}
	return db
}

var instance *gorm.DB
var once sync.Once

func GET_DB() *gorm.DB {
	once.Do(func() {
		config := DBconnectConfig{
			Host:     constants.DB_HOST,
			Port:     constants.DB_PORT,
			Db:       constants.DB_NAME,
			User:     constants.DB_USER,
			Password: constants.DB_PASSWORD,
		}
		instance = config.connect()
		//开启调试模式
		DEBUG := os.Getenv("DEBUG")
		if DEBUG != "" {
			instance = instance.DEBUG()
		}
	})
	return instance
}
