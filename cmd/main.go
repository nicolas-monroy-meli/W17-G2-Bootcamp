package main

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"

	server "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/application"
)

func main() {
	err := godotenv.Load("docs/db/dev.env")
	if err != nil {
		log.Println("No .env found!")
	}
	cfg := &server.SQLConfig{
		Database: mysql.Config{
			User:      os.Getenv("DB_USER"),
			Passwd:    "",
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
