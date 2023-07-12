package main

import (
	"github.com/Asrez/GoAPIBlog/api"
	"github.com/Asrez/GoAPIBlog/config"
	"github.com/Asrez/GoAPIBlog/data/db"
	"github.com/Asrez/GoAPIBlog/data/migration"
	"github.com/Asrez/GoAPIBlog/pkg/logging"
)

func main(){
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	err := db.InitDb(cfg)
	defer db.CloseDb()

	if err != nil {
		logger.Fatal(logging.Postgres , logging.Startup , err.Error(),nil)
	}

	migration.Up()
	api.InitServer()
}