package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func sqliteInit(config DataConfig, autoMigrateTable ...interface{}) (*gorm.DB, error) {

	if db, err := gorm.Open(sqlite.Open(config.Database), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用自动创建外键约束
	}); err != nil {
		return nil, err
	} else {
		//sqlDB, _ := db.DB()
		//sqlDB.SetMaxIdleConns(*config.MaxIdleConns)
		//sqlDB.SetMaxOpenConns(config.MaxOpenConns)
		err := db.AutoMigrate(autoMigrateTable...)
		if err != nil {
			return nil, err
		}
		return db, nil
	}
}
