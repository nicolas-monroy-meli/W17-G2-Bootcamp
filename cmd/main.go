package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	server "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/application"
)

func main() {
	err := godotenv.Load("dev.env.example")
	if err != nil {
		log.Println("No .env found!")
	}
	cfg := &server.SQLConfig{
		Database: mysql.Config{
			User:      os.Getenv("DB_USER"),
			Passwd:    os.Getenv("PASSWD"),
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
		fmt.Println(err)
		return
	}
}
