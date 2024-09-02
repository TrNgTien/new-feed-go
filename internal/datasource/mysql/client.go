package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewMysqlClient() *sql.DB {
	fmt.Println("[InitMysql] Connecting Mysql!!")
	db, err := sql.Open("mysql", "root:new_feed@/new-feed")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	fmt.Println("[InitMysql] Connected Mysql!!")
	return db
}
