package main

import (
	"fmt"

	server "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/application"
)

func main() {
	cfg := &server.ConfigServerChi{
		ServerAddress:  ":8080",
		LoaderFilePath: "docs/db/",
	}
	app := server.NewServerChi(cfg)
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
