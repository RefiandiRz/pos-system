package services

import (
	"errors"

	"github.com/RefiandiRz/pos-system/internal/models"
	"github.com/RefiandiRz/pos-system/internal/repositories"
	"github.com/RefiandiRz/pos-system/utils"
	"golang.org/x/crypto/bcrypt"
)

func Register(req models.RegisterRequest) (*models.AuthResponse, error) {
	// Check if email already exist
	existingEmail, _ := repositories.FindUserByEmail(req.Email)
	if existingEmail != nil {
		return nil, errors.New("email already registered")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Default role to cashier if not provided
	role := req.Role
	if role == "" {
		role = models.RoleCashier
	}

	// Create user
	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     role,
	}

	if err := repositories.CreateUser(user); err != nil {
		return nil, errors.New("failed to create user")
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &models.AuthResponse{
		Token: token,
		User:  *user,
	}, nil

}

func Login(req models.LoginRequest) (*models.AuthResponse, error) {
	// Find user by email
	user, err := repositories.FindUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// compare password with hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &models.AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}
