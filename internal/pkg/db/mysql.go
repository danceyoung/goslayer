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

	dbconn, err := sql.Open("mysql", mysqlconnstring)
	if err != nil {
		panic(err)
	} else {
		if err = dbconn.Ping(); err != nil {
			panic(err)
		} else {
			dbconn.SetMaxOpenConns(20)
			dbconn.SetMaxIdleConns(200)
			mysqldb = dbconn

			log.Println("mysql has already prepared for user connection.")
		}
	}

}

//NewMysql returns mysql connection instance
func NewMysql() *sql.DB {
	return mysqldb
}
