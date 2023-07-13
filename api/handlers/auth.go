package handler

import (
	"net/http"

	"github.com/Asrez/GoAPIBlog/api/dto"
	helper "github.com/Asrez/GoAPIBlog/api/helpers"
	"github.com/Asrez/GoAPIBlog/config"
	"github.com/Asrez/GoAPIBlog/services"
	"github.com/gin-gonic/gin"
	
)

type AuthHandler struct {
	service *services.UserService
}

func NewUsersHandler(cfg *config.Config) *AuthHandler {
	service := services.NewUserService(cfg)
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Login(c *gin.Context) {
	req := new(dto.LoginByUsernameRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}
	token, err := h.service.LoginByUsername(req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(token, true, helper.Success))
}

func (h *AuthHandler) Register(c *gin.Context) {
	req := new(dto.RegisterUserByUsernameRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}
	err = h.service.RegisterByUsername(req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, helper.Success))
}
