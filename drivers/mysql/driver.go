package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type ConfigDB struct {
	DB_Username string
	DB_Password string
	DB_Database string
}

func (config *ConfigDB) InitDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Database,
	)

	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
