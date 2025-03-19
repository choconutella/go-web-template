package utils

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/alexbrainman/odbc"
	ora "github.com/sijms/go-ora/v2"
)

type DBConnector struct {
	Config Config
}

func (conn *DBConnector) Connect() (*sql.DB, error) {
	connStr := ora.BuildUrl(
		conn.Config.DBHost,
		conn.Config.DbPort,
		conn.Config.DbName,
		conn.Config.DbUser,
		conn.Config.DbPass, nil,
	)
	db, err := sql.Open("oracle", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	db.SetMaxOpenConns(conn.Config.MaxOpenConn)
	db.SetMaxIdleConns(conn.Config.MaxIdleConn)
	db.SetConnMaxLifetime(conn.Config.MaxLifetimeConn * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return db, nil
}

func (conn *DBConnector) ConnectOdbc() (*sql.DB, error) {

	connStr := fmt.Sprintf("DSN=%s", conn.Config.DsnName)
	db, err := sql.Open("odbc", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	db.SetMaxOpenConns(conn.Config.MaxOpenConn)
	db.SetMaxIdleConns(conn.Config.MaxIdleConn)
	db.SetConnMaxLifetime(conn.Config.MaxLifetimeConn * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return db, nil
}

func (conn *DBConnector) GetDBConnector(config *Config) (*sql.DB, error) {
	if config.UseOdbc {
		return conn.ConnectOdbc()
	}
	return conn.Connect()
}
