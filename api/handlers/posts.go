package handler

import (
	"log"
	"net/http"

	"github.com/Asrez/GoAPIBlog/api/dto"
	helper "github.com/Asrez/GoAPIBlog/api/helpers"
	"github.com/Asrez/GoAPIBlog/config"
	"github.com/Asrez/GoAPIBlog/data/db"
	"github.com/Asrez/GoAPIBlog/data/models"
	"github.com/gin-gonic/gin"
)


type PostsHandler struct {
	cfg    *config.Config
}

func NewPostsHandler(cfg *config.Config) *PostsHandler {
	return &PostsHandler{}
}

func (p *PostsHandler) GetAllPost(c *gin.Context){
	db := db.GetDb()

	var posts []models.Post
	result := db.Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError,helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, result.Error))
		return
	}

	c.JSON(http.StatusOK,helper.GenerateBaseResponse(posts , true ,helper.Success))
}

func (p *PostsHandler) CreatePost(c *gin.Context){
	db := db.GetDb()
	
	req := new(dto.CreatePost)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}
	
	post := models.Post{Title: req.Title , Content: req.Content , AuthorId: req.AuthorId}
	tx := db.Begin()
	err = tx.Create(&post).Error
	if err != nil {
		tx.Rollback()
		log.Fatalln("error creating post")
		return 
	}
	tx.Commit()
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse("created", true, helper.Success))
}

func (p *PostsHandler) Get(c *gin.Context) {
	db := db.GetDb()
	id := c.Param("id")

	var post models.Post
	result := db.First(&post, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, helper.GenerateBaseResponseWithError(nil, false, helper.NotFoundError, result.Error))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(post, true, helper.Success))
}

func (p *PostsHandler) UpdatePost(c *gin.Context) {
	db := db.GetDb()
	id := c.Param("id")

	var post models.Post
	result := db.First(&post, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, helper.GenerateBaseResponseWithError(nil, false, helper.NotFoundError, result.Error))
		return
	}
	req := new(dto.UpdatePost)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	post.Title = req.Title
	post.Content = req.Content

	result = db.Save(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, result.Error))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(post, true, helper.Success))
}

func (p *PostsHandler) DeletePost(c *gin.Context) {
	db := db.GetDb()
	id := c.Param("id")

	var post models.Post
	result := db.First(&post, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, helper.GenerateBaseResponseWithError(nil, false, helper.NotFoundError, result.Error))
		return
	}

	result = db.Delete(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, result.Error))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse("deleted", true, helper.Success))
}
