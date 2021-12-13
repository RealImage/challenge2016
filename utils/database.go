package utils

import (
	"fmt"
	"sync"

	"github.com/rs/zerolog/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var once sync.Once
var DB *gorm.DB

func GetInstancemysql() (dba *gorm.DB) {
	once.Do(func() {
		user := "root"
		password := "rootpassword"
		host := "127.0.0.1"
		port := "3306"
		dbname := "qubecinema"
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
		db, err := gorm.Open(mysql.Open(dsn+"?parseTime=true"), &gorm.Config{})
		//close connection - cleanup and close
		dba = db
		if err != nil {
			log.Panic().Msgf("Error connecting to the database at %s:%s/%s", host, port, dbname)
			log.Info().Msgf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
		}

		sqlDB, err := dba.DB()
		if err != nil {
			log.Panic().Msgf("Error getting GORM DB definition")
		}
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(10)
		//defer sqlDB.Close()

		log.Info().Msgf("Successfully established connection to %s:%s/%s", host, port, dbname)

	})

	DB = dba
	return dba
}
