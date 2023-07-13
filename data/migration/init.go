package migration

import (
	"github.com/Asrez/GoAPIBlog/config"
	"github.com/Asrez/GoAPIBlog/constants"
	"github.com/Asrez/GoAPIBlog/data/db"
	"github.com/Asrez/GoAPIBlog/data/models"
	"github.com/Asrez/GoAPIBlog/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


var logger = logging.NewLogger(config.GetConfig())


func Up(){
	database := db.GetDb()
	createTables(database)
	createDefaultUserInformation(database)
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
	
	tables = addNewTable(database, models.User{} , tables)
	tables = addNewTable(database , models.Post{},tables)

	err := database.Migrator().CreateTable(tables...)
	if err != nil {
		logger.Error(logging.Postgres, logging.Migration, err.Error(), nil)
	}
	logger.Info(logging.Postgres, logging.Migration, "tables created", nil)

}


func createDefaultUserInformation(database *gorm.DB) {
	u := models.User{Username: constants.DefaultUserName, Email: constants.DefaultEmail}
	pass := constants.DefaultPassword
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	u.Password = string(hashedPassword)
	createAdminUserIfNotExists(database, &u)
}


func createAdminUserIfNotExists(database *gorm.DB, u *models.User) {
	exists := 0
	database.
		Model(&models.User{}).
		Select("1").
		Where("username = ?", u.Username).
		First(&exists)
	if exists == 0 {
		database.Create(u)
	}
}
