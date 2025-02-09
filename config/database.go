package config

import (
	"line-bk-api/pkg/logs"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func ConnectDB() {
	var err error
	DB, err = sqlx.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		logs.Error(err)
	}

	if err = DB.Ping(); err != nil {
		logs.Error(err)
	}

	logs.Info("Connected to Database")
}

func GetDBInstance() *sqlx.DB {
	return DB
}
