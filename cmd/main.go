package main

import (
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	server "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/application"
)

func main() {
	//load .env
	err := godotenv.Load("dev.env")
	if err != nil {
		panic("No .env found!")
	}
	cfg := &server.SQLConfig{
		Database: mysql.Config{
			User:      os.Getenv("DB_USER"),
			Passwd:    os.Getenv("DB_PASSWORD"),
			Net:       "tcp",
			Addr:      os.Getenv("DB_ADDRESS"),
			DBName:    os.Getenv("DB_NAME"),
			ParseTime: true,
		},
		Address: os.Getenv("API_ADDRESS"),
	}
	app := server.NewSQLConfig(cfg)
	// - run
	if err := app.Run(); err != nil {
		panic(err)
	}
}
