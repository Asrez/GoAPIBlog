package api

import (
	"fmt"
	"log"

	"github.com/Asrez/GoAPIBlog/api/routers"
	"github.com/Asrez/GoAPIBlog/api/validations"
	"github.com/Asrez/GoAPIBlog/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer(){
	config := config.GetConfig()
	r := gin.New()
	r.Use(gin.Logger()  ,gin.Recovery())
	RegisterValidators()
	v1 := r.Group("api/v1")
	{
		auth := v1.Group("auth")
		routers.Auth(auth , config)

		posts := v1.Group("posts")
		routers.Posts(posts,config)
	}

	r.Run(fmt.Sprintf(":%s", config.Server.InternalPort))
}


func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("password", validation.PasswordValidator, true)
		if err != nil {
			// logger.Error(logging.Validation, logging.Startup, err.Error(), nil)
			log.Fatal("validation failed")
		}
	}
}
