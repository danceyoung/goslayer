package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var mysqldb *sql.DB

//user:password@tcp(127.0.0.1:3306)/goslayer_db?charset=utf8
const mysqlconnstring string = ""

func init() {
	if len(mysqlconnstring) > 0 {
		dbconn, err := sql.Open("mysql", mysqlconnstring)
		if err != nil {
			panic(err)
		}
		dbconn.SetMaxOpenConns(20)
		dbconn.SetMaxIdleConns(200)
		mysqldb = dbconn
	}

	log.Println("mysql is opened successfully")
}

//NewMysql returns mysql connection instance
func NewMysql() *sql.DB {
	return mysqldb
}
