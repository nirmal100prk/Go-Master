package main

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

// Example input struct with various validation tags
type User struct {
	Username    string    `json:"username" validate:"required,min=3,max=32"`
	Email       string    `json:"email" validate:"required,email"`
	Password    string    `json:"password" validate:"required,min=8,max=64"`
	Age         int       `json:"age" validate:"required,gte=18,lte=100"`
	Website     string    `json:"website" validate:"url"`
	Registered  bool      `json:"registered"`
	SignupDate  time.Time `json:"signup_date" validate:"required"`
	CustomField string    `json:"custom_field" validate:"customValidation"`
}

// Custom validation function
func customValidation(fl validator.FieldLevel) bool {
	// Example custom validation: Check if the field contains only alphanumeric characters
	match, _ := regexp.MatchString("^[a-zA-Z0-9]+$", fl.Field().String())
	return match
}

func main() {
	// Initialize the validator
	validate := validator.New()

	// Register the custom validation function
	validate.RegisterValidation("customValidation", customValidation)

	// Example user input
	user := User{
		Username:    "JohnDoe",
		Email:       "john.doe@example.com",
		Password:    "password123",
		Age:         25,
		Website:     "https://example.com",
		Registered:  true,
		SignupDate:  time.Now(),
		CustomField: "CustomValue1", // Valid because it contains only alphanumeric characters
	}

	// Validate the user input
	err := validate.Struct(user)
	if err != nil {
		// If there is a validation error, print it
		for _, err := range err.(validator.ValidationErrors) {
			log.Println("Validation failed:", err.Namespace(), err.Tag(), err.Param(), err.Value())
		}
	} else {
		fmt.Println("Validation successful!")
	}
}
