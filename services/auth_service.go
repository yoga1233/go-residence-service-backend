package service

import (
	"errors"

	"github.com/yoga1233/go-residence-service-backend/helper"
	model "github.com/yoga1233/go-residence-service-backend/models"
	"github.com/yoga1233/go-residence-service-backend/repositories"
	"github.com/yoga1233/go-residence-service-backend/utils"
)

type AuthService interface {
	Register(user *model.User) error
	Login(email, password string) (helper.UserResponse, error)
}

type authService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{
		userRepository: userRepo,
	}
}

// Login implements AuthService.
func (s *authService) Login(email string, password string) (helper.UserResponse, error) {
	//find user by username
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return helper.UserResponse{}, errors.New("invalid username or password")
	}

	//check password
	if !utils.CheckPasswordHash(password, user.Password) {
		return helper.UserResponse{}, errors.New("invalid username or password")
	}

	//generate jwt
	token, err := utils.GenerateJWT(email)
	if err != nil {
		return helper.UserResponse{}, errors.New("failed to generate token")
	}

	return helper.UserResponse{
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}, nil

}

// Register implements AuthService.
func (s *authService) Register(user *model.User) error {
	// check if user already exists
	existingUser, _ := s.userRepository.FindByEmail(user.Email)
	if existingUser != nil {
		return errors.New("email already registered")
	}

	// hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return errors.New("failed to hash password")
	}
	user.Password = hashedPassword

	return s.userRepository.CreateUser(user)

}
