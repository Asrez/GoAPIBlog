package api

import (
	"fmt"

	"github.com/Asrez/GoAPIBlog/config"
	"github.com/Asrez/GoAPIBlog/api/routers"
	"github.com/gin-gonic/gin"
)

func InitServer(){
	config := config.GetConfig()
	r := gin.New()
	r.Use(gin.Logger()  ,gin.Recovery())

	v1 := r.Group("api/v1")
	{
		auth := v1.Group("auth")
		routers.Auth(auth , config)
	}

	r.Run(fmt.Sprintf(":%s", config.Server.InternalPort))
}