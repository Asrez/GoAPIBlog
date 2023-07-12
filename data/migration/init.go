package migration

import (
	"github.com/Asrez/GoAPIBlog/config"
	"github.com/Asrez/GoAPIBlog/data/db"
	"github.com/Asrez/GoAPIBlog/pkg/logging"
	"gorm.io/gorm"
)


var logger = logging.NewLogger(config.GetConfig())


func Up(){
	database := db.GetDb()
	createTables(database)
	logger.Info(logging.Postgres, logging.Migration, "UP", nil)

}


func addNewTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}

func createTables(database *gorm.DB){
	tables := []interface{}{}

	// tables = addNewTable(database , models.Categories{} , tables)


	err := database.Migrator().CreateTable(tables...)
	if err != nil {
		logger.Error(logging.Postgres, logging.Migration, err.Error(), nil)
	}
	logger.Info(logging.Postgres, logging.Migration, "tables created", nil)

}
