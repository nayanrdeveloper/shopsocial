package responses

import (
	"net/http"

	"shopsocial-backend/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// APIResponse - Standard response format for both success & error
type APIResponse struct {
	Success bool        `json:"success"`           // Indicates if the request was successful
	Status  int         `json:"status"`            // HTTP status code
	Message string      `json:"message,omitempty"` // Human-friendly message for the frontend
	Data    interface{} `json:"data,omitempty"`    // Actual payload (for success)
	Error   string      `json:"error,omitempty"`   // Internal error details (for debugging)
}

// SendSuccess is a helper to send a 200 OK response
func SendSuccess(c *gin.Context, message string, data interface{}) {
	response := APIResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
	}
	c.JSON(http.StatusOK, response)
}

// SendCreated is a helper to send a 201 Created response
func SendCreated(c *gin.Context, message string, data interface{}) {
	response := APIResponse{
		Success: true,
		Status:  http.StatusCreated,
		Message: message,
		Data:    data,
	}
	c.JSON(http.StatusCreated, response)
}

// SendDeleted is a helper to send a 200 OK response
// or 204 No Content if you prefer not returning any body.
func SendDeleted(c *gin.Context, message string) {
	response := APIResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: message,
	}
	c.JSON(http.StatusOK, response)
}

// SendError is a helper to send an error response with a custom status code.
// 'err' is logged internally but not fully exposed to the client.
func SendError(c *gin.Context, status int, userMsg string, err error) {
	var internalError string
	if err != nil {
		internalError = err.Error()
		// Log the full internal error for debugging
		logger.Log.Error("API Error",
			zap.Int("status", status),
			zap.String("userMsg", userMsg),
			zap.Error(err),
		)
	}

	response := APIResponse{
		Success: false,
		Status:  status,
		Message: userMsg,       // user-friendly
		Error:   internalError, // internal debugging
	}
	c.JSON(status, response)
}
