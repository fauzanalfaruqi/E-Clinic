package dto

import "database/sql"

type ConfigData struct {
	DbConfig dbConfig
	AppConfig appConfig
}

type dbConfig struct {
	Host string
	Port string
	User string
	Pass string
	Database string
	MaxIdle int
	MaxConn int
	MaxLifeTime string
	LogMode int
}

type appConfig struct {
	Port string
}

type Db struct {
	*sql.DB
}