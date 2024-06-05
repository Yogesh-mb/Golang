package db

import (
	"time"
	"gorm.io/gorm"
	"errors"
	"log"
	"config"
	"gorm.io/gorm/logger"
	"gorm.io/driver/mysql" 
)

type DB struct{
	Db *gorm.DB
	Prefix string
}

func New(conf config.DB) DB {
	username := conf.username
	password := conf.Password
	host := conf.host
	port := conf.port
	database := conf.Name
	dsn := username + ':' + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	gormDB, err :=  gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error)
	})
	if err != nil {
		log.Fatal(errors.New("unable to open a MySQL DB session. Error: " + err.Error()))
	}

	sqlDB, dbErr := gormDB.DB()
	if dbErr != nil {
		log.Fatal("unable to connect to MySQL DB. Error: " + dbErr.Error())
	}

	//SetMaxIdleConns sets the maximun no of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(conf.MaxIdleConnections)

	//SetMaxOpenConns sets the maximum no of open connections to the database.
	sqlDB.SetMaxOpenConns(conf.PoolSize)

	//SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(conf.ConnectionMaxTime) * time.Second)

	return DB{
		Db : gormDB,
		Prefix : conf.Prefix
	}
}

