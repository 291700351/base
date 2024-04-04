package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

func mysqlInit(config DataConfig, autoMigrateTable ...interface{}) (*gorm.DB, error) {
	dsn := "root" + ":" + "root" + "@tcp(" + "127.0.0.1" + ":" + strconv.Itoa(3306) + ")/" +
		config.Database + "?charset=" + "utf8" + "&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
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
