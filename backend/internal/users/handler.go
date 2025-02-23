package users

import (
    "net/http"
    "shopsocial-backend/pkg/constants"
    "shopsocial-backend/pkg/logger"
    "shopsocial-backend/pkg/responses"

    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type UserHandler struct {
    Service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
    return &UserHandler{Service: service}
}

// RegisterUser handles POST /api/users/register
func (h *UserHandler) RegisterUser(c *gin.Context) {
    var body struct {
        FullName string `json:"full_name" binding:"required"`
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required,min=6"`
    }

    if err := c.ShouldBindJSON(&body); err != nil {
        responses.SendError(c, http.StatusBadRequest, constants.ErrInvalidRequest, err)
        return
    }

    createdUser, err := h.Service.RegisterUser(body.FullName, body.Email, body.Password)
    if err != nil {
        logger.Log.Warn("Failed to register user", zap.Error(err))
        responses.SendError(c, http.StatusConflict, err.Error(), nil)
        return
    }

    responses.SendCreated(c, "User registered successfully", createdUser)
}

// LoginUser handles POST /api/users/login
func (h *UserHandler) LoginUser(c *gin.Context) {
    var body struct {
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required,min=6"`
    }

    if err := c.ShouldBindJSON(&body); err != nil {
        responses.SendError(c, http.StatusBadRequest, constants.ErrInvalidRequest, err)
        return
    }

    token, err := h.Service.LoginUser(body.Email, body.Password)
    if err != nil {
        logger.Log.Warn("Login failed", zap.Error(err))
        responses.SendError(c, http.StatusUnauthorized, "Invalid credentials", nil)
        return
    }

    responses.SendSuccess(c, "Login successful", gin.H{"token": token})
}
