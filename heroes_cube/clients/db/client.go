package db

import (
	"fmt"
	"heroes_cube/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetMySqlClient(config configs.Config) (*gorm.DB, error) {

	dsn := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn = fmt.Sprintf(dsn, config.DbUser, config.DbPass, config.DbHost, config.DbPort, config.DbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil

}
