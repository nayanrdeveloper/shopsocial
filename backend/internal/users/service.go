package users

import (
	"errors"
	"shopsocial-backend/pkg/jwt"
	"shopsocial-backend/pkg/logger"
	"shopsocial-backend/pkg/utils"

	"go.uber.org/zap"
)

type UserService struct {
    Repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
    return &UserService{Repo: repo}
}

// RegisterUser handles user registration (hash password, store in DB)
func (s *UserService) RegisterUser(fullName, email, plainPassword string) (*User, error) {
    // Check if email already exists
    existingUser, _ := s.Repo.FindByEmail(email)
    if existingUser != nil {
        logger.Log.Warn("Email already in use", zap.String("email", email))
        return nil, errors.New("email already in use")
    }

    // Hash password
    hashedPassword, err := utils.HashPassword(plainPassword)
    if err != nil {
        return nil, errors.New("failed to hash password")
    }

    // Create User object
    user := &User{
        FullName: fullName,
        Email:    email,
        Password: hashedPassword,
    }

    // Save to database
    createdUser, err := s.Repo.CreateUser(user)
    if err != nil {
        logger.Log.Error("Failed to create user", zap.Error(err))
        return nil, errors.New("failed to create user")
    }

    return createdUser, nil
}

// LoginUser validates user credentials and returns a JWT token
func (s *UserService) LoginUser(email, plainPassword string) (string, error) {
    user, err := s.Repo.FindByEmail(email)
    if err != nil {
        return "", errors.New("invalid email or password")
    }

    // Compare password
    if ok := utils.CheckPasswordHash(plainPassword, user.Password); !ok {
        return "", errors.New("invalid email or password")
    }

    // Generate JWT token
    token, err := jwt.GenerateJWT(user.ID.Hex())
    if err != nil {
        return "", errors.New("failed to generate token")
    }

    return token, nil
}

// GetUserByEmail can be used for other operations
func (s *UserService) GetUserByEmail(email string) (*User, error) {
    return s.Repo.FindByEmail(email)
}
