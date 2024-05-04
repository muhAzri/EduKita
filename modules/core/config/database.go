package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitializeDatabase() (*sql.DB, error) {

	connStr := GetConnectionString()

	var err error
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetConnectionString() string {
	dbUsername := GetEnvValue("DB_USER")
	dbPassword := GetEnvValue("DB_PASS")
	dbName := GetEnvValue("DB_NAME")
	dbHost := GetEnvValue("DB_HOST")
	dbPort := GetEnvValue("DB_PORT")
	dbSSLMode := GetEnvValue("DB_SSL_MODE")

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",
		dbHost, dbUsername, dbPassword, dbName, dbPort, dbSSLMode)

	return connStr
}
