package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func newstorage(cfg mysql.Config) (*sql.DB, error){

	db ,err := sql.Open("mysql",cfg.FormatDSN())

	if err != nil{
		return nil,err
	}

	err := db.Ping()

	if err != nil{
		return nil,err
	}

	return db, nil
}