package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"log"
)

func InitDatabase() (*gorm.DB, *sql.DB) {
	if os.Getenv("DB_CONNECTION") == "" {
		os.Setenv("DB_CONNECTION", "mysql")
		os.Setenv("DB_HOST", "localhost")
		os.Setenv("DB_PORT", "3306")
		os.Setenv("DB_DATABASE", "tenor")
		os.Setenv("DB_USERNAME", "root")
		os.Setenv("DB_PASSWORD", "")
	}

	dbConnection := os.Getenv("DB_CONNECTION")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	if dbConnection == "mysql" {
		// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True"
		dsn := fmt.Sprint(dbUsername, ":", dbPassword, "@tcp(", dbHost, ":", dbPort, ")/", dbName, "?charset=utf8mb4&parseTime=True")
		db, err := gorm.Open(mysql.New(mysql.Config{
			DSN: dsn,
		}), &gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
		})

		if err != nil {
			log.Panic(err)
		}

		sqlDB, err := db.DB()

		if err != nil {
			log.Panic(err)
		}

		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(1000)
		sqlDB.SetConnMaxLifetime(time.Hour)

		return db, sqlDB
	} else if dbConnection == "pgsql" {
		// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable"
		dsn := fmt.Sprint("host=", dbHost, " user=", dbUsername, " password=", dbPassword, " dbname=", dbName, " port=", dbPort, " sslmode=disable")
		db, err := gorm.Open(postgres.New(postgres.Config{
			DSN: dsn,
		}), &gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
		})

		if err != nil {
			log.Panic(err)
		}

		sqlDB, err := db.DB()

		if err != nil {
			log.Panic(err)
		}

		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(999)
		sqlDB.SetConnMaxLifetime(time.Hour)

		return db, sqlDB
	}

	fmt.Println("Init database failed")
	return nil, nil
}
