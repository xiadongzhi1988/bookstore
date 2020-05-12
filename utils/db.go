package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db *sql.DB
	err error
)

func init()  {
	Db, err = sql.Open("mysql", "root:root@tcp(10.0.0.51:3306)/bookstore0612")
	if err != nil {
		panic(err.Error())
	}
}