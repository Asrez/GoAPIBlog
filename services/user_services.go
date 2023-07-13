package services

import (
	"github.com/Asrez/GoAPIBlog/api/dto"
	"github.com/Asrez/GoAPIBlog/config"
	"github.com/Asrez/GoAPIBlog/data/db"
	"github.com/Asrez/GoAPIBlog/data/models"
	"github.com/Asrez/GoAPIBlog/pkg/logging"
	"github.com/Asrez/GoAPIBlog/pkg/service_errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	logger       logging.Logger
	cfg          *config.Config
	tokenService *TokenService
	database     *gorm.DB
}

func NewUserService(cfg *config.Config) *UserService {
	database := db.GetDb()
	logger := logging.NewLogger(cfg)
	return &UserService{
		cfg:          cfg,
		database:     database,
		logger:       logger,
		tokenService: NewTokenService(cfg),
	}
}

// Login 
func (s *UserService) LoginByUsername(req *dto.LoginByUsernameRequest) (*dto.TokenDetail, error) {
	var user models.User
	err := s.database.
		Model(&models.User{}).
		Where("username = ?", req.Username).
		Find(&user).Error
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	tdto := tokenDto{UserId: user.Id, Username: user.Username,
		Email: user.Email}

	token, err := s.tokenService.GenerateToken(&tdto)
	if err != nil {
		return nil, err
	}
	return token, nil

}


// Register
func (s *UserService) RegisterByUsername(req *dto.RegisterUserByUsernameRequest) error {
	u := models.User{Username: req.Username, Email: req.Email}

	exists, err := s.existsByEmail(req.Email)
	if err != nil {
		return err
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.EmailExists}
	}
	exists, err = s.existsByUsername(req.Username)
	if err != nil {
		return err
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.UsernameExists}
	}

	bp := []byte(req.Password)
	hp, err := bcrypt.GenerateFromPassword(bp, bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error(logging.General, logging.HashPassword, err.Error(), nil)
		return err
	}
	u.Password = string(hp)

	tx := s.database.Begin()
	err = tx.Create(&u).Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return err
	}
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return err
	}
	tx.Commit()
	return nil
}

func (s *UserService) existsByEmail(email string) (bool, error) {
	var exists bool
	if err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&exists).
		Error; err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (s *UserService) existsByUsername(username string) (bool, error) {
	var exists bool
	if err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("username = ?", username).
		Find(&exists).
		Error; err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}
