package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/TrNgTien/new-feed-go/internal/configs"
	"github.com/doug-martin/goqu"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func NewMysqlClient(databaseConfig configs.Database, l *zap.Logger) (*sql.DB, func(), error) {
	l.Info("[InitMysql] connecting to the database!")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		databaseConfig.Username,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.Database,
	)

	db, err := sql.Open("mysql", connectionString)
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	if err != nil {
		l.With(zap.Error(err)).Error("[InitMysql] Error connecting to the database!")
		return nil, nil, err
	}

	l.Info("[InitMysql] connected to the database!")

	cleanup := func() {
		db.Close()
	}

	return db, cleanup, nil
}

func InitializeSqlBuilder(db *sql.DB) (*goqu.Database, error) {
	return goqu.New("mysql", db), nil
}
