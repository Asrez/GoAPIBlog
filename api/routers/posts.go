package routers

import (
	handler "github.com/Asrez/GoAPIBlog/api/handlers"
	"github.com/Asrez/GoAPIBlog/config"
	"github.com/gin-gonic/gin"
)

func Posts(router *gin.RouterGroup, cfg *config.Config){
	p := handler.NewPostsHandler(cfg)
	router.GET("/" ,p.GetAllPost)
	router.POST("/",p.CreatePost)
	router.GET("/:id" , p.Get)
	router.PUT("/:id",p.UpdatePost)
	router.DELETE("/:id",p.DeletePost)
}