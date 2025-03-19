package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	DbUser string
	DbPass string
	DBHost string
	DbPort int
	DbName string

	MaxOpenConn     int
	MaxIdleConn     int
	MaxLifetimeConn time.Duration

	UseOdbc bool
	DsnName string

	ServerPort string
}

func NewConfig() (*Config, error) {

	DbUser := os.Getenv("DBUSER")
	DbPass := os.Getenv("DBPASS")
	DbHost := os.Getenv("DBHOST")
	DbPort := os.Getenv("DBPORT")
	DbName := os.Getenv("DBNAME")
	PortInt, err := strconv.Atoi(DbPort)
	if err != nil {
		return nil, fmt.Errorf("failed to convert value to int for param `_DBPORT`: %w", err)
	}

	MaxOpenConn := os.Getenv("MAX_OPEN_CONN")
	OpenConnInt, err := strconv.Atoi(MaxOpenConn)
	if err != nil {
		return nil, fmt.Errorf("failed to convert value to int for param `_MAX_OPEN_CONN`: %w", err)
	}

	MaxIdleConn := os.Getenv("MAX_IDLE_CONN")
	IdleConnInt, err := strconv.Atoi(MaxIdleConn)
	if err != nil {
		return nil, fmt.Errorf("failed to convert value to int for param `_MAX_IDLE_CONN`: %w", err)
	}

	MaxLifetimeConn := os.Getenv("MAX_LIFETIME_CONN")
	LifetimeInt, err := strconv.Atoi(MaxLifetimeConn)
	if err != nil {
		return nil, fmt.Errorf("failed to convert value to int for param `_MAX_IDLE_CONN`: %w", err)
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "5374"
	}

	useOdbc := false
	if os.Getenv("USE_ODBC") == "Y" || os.Getenv("USE_ODBC") == "" {
		useOdbc = true
	}

	dsnName := os.Getenv("DSN_NAME")
	if dsnName == "" {
		dsnName = "mydsn"
	}

	return &Config{
		DbUser:          DbUser,
		DbPass:          DbPass,
		DBHost:          DbHost,
		DbPort:          PortInt,
		DbName:          DbName,
		MaxOpenConn:     OpenConnInt,
		MaxIdleConn:     IdleConnInt,
		MaxLifetimeConn: time.Duration(LifetimeInt),
		ServerPort:      serverPort,
		UseOdbc:         useOdbc,
		DsnName:         dsnName,
	}, nil
}
