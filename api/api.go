package api

import (
	"fmt"

	"github.com/Asrez/GoAPIBlog/config"
	"github.com/gin-gonic/gin"
)

func InitServer(){
	config := config.GetConfig()
	r := gin.New()
	r.Use(gin.Logger()  ,gin.Recovery())

	// v1 := r.Group("api/v1")
	// {
		
	// }

	r.Run(fmt.Sprintf(":%s", config.Server.InternalPort))
}