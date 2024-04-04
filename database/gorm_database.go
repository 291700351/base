package database

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func ConnectDatabase(config DataConfig, autoMigrateTable ...interface{}) {

	if config.Database == "" {
		panic(errors.New("please set a database"))
	}
	var err error
	switch config.Driver {
	case "mysql":
		DB, err = mysqlInit(config, autoMigrateTable...)
	case "sqlite":
		DB, err = sqliteInit(config, autoMigrateTable...)
	default:
		DB, err = sqliteInit(config, autoMigrateTable...)
	}
	if nil != err {
		panic(err)
	}
}

func Close() {
	if nil != DB {
		s, err := DB.DB()
		if nil != err {
			fmt.Println("Get sql.DB error")
		} else {
			if err := s.Close(); nil != err {
				fmt.Println("Close database error")
			}
		}
	}
}
