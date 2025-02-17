package validationerror

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationResponse struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

var ErrValidator = map[string]string{}

// Fungsi untuk mengubah error validasi menjadi response JSON
func ErrValidationResponse(err error) []ValidationResponse {
	var validationResponses []ValidationResponse
	var fieldErrors validator.ValidationErrors

	// Mengecek apakah error berasal dari validasi
	if errors.As(err, &fieldErrors) {
		for _, err := range fieldErrors {
			switch err.Tag() {
			case "required":
				validationResponses = append(validationResponses, ValidationResponse{
					Field:   err.Field(),
					Message: fmt.Sprintf("%s is required", err.Field()),
				})
			case "email":
				validationResponses = append(validationResponses, ValidationResponse{
					Field:   err.Field(),
					Message: fmt.Sprintf("%s is not a valid email", err.Field()),
				})
			default:
				errValidator, ok := ErrValidator[err.Tag()]	
				if ok {
					count := strings.Count(errValidator, "%s")
					if count == 1 {
						va
					}
				}
			}

		}
	}

	return validationResponses
}
