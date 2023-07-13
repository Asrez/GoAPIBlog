package routers

import (
	handler "github.com/Asrez/GoAPIBlog/api/handlers"
	"github.com/Asrez/GoAPIBlog/config"
	"github.com/gin-gonic/gin"
)

func Auth(router *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewUsersHandler(cfg)
	router.POST("/register", h.Register)
	router.POST("/login", h.Login)
}