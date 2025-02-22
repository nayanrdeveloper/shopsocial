package constants

import "fmt"

// ✅ Common success messages with placeholders
const (
	SuccessCreated = "%s created successfully"
	SuccessUpdated = "%s updated successfully"
	SuccessDeleted = "%s deleted successfully"
	SuccessFetched = "%s retrieved successfully"
)

// ✅ Common error messages
const (
	ErrInvalidRequest = "Invalid request body"
	ErrNotFound       = "%s not found"
	ErrCreationFailed = "Failed to create %s"
	ErrUpdateFailed   = "Failed to update %s"
	ErrDeletionFailed = "Failed to delete %s"
	ErrRouteNotFound  = "Route not found: %s"
)

// ✅ Generic function to format messages dynamically
func FormatMessage(template string, entity string) string {
	return fmt.Sprintf(template, entity)
}
