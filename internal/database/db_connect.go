package database

import (
	"database/sql"
	"fmt"
	"go-database/config"
	"log"
	// "time"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDb() {
	// Capture connection properties
	cfg := mysql.NewConfig()
	cfg.User = config.GetEnv("DB_USER")
	cfg.Passwd = config.GetEnv("DB_PWD")
	cfg.Net = "tcp"
	cfg.Addr = config.GetEnv("DB_HOST")
	cfg.DBName = config.GetEnv("DB_NAME")
	cfg.ParseTime = true

	fmt.Println(config.GetEnv("DB_HOST"))

	var err error
	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	// defer DB.Close()

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Database connected successful!")

	// DB.SetConnMaxLifetime(time.Minute * 3)
	// DB.SetMaxOpenConns(10)
	// DB.SetMaxIdleConns(10)
}
