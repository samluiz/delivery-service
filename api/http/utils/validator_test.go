package utils

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type ValidatedUserStruct struct {
	Name  string `validate:"required"`
	Age   int    `validate:"gte=18"`
	Email string `validate:"required,email"`
}

type NoValidationStruct struct {
	Name string
}

func TestXValidator_ValidStruct(t *testing.T) {
	validator := XValidator{Validator: validator.New()}
	data := ValidatedUserStruct{
		Name:  "John Doe",
		Age:   25,
		Email: "john.doe@example.com",
	}

	validationErrors := validator.Validate(data)
	assert.Empty(t, validationErrors, "Expected no validation errors for valid struct")
}

func TestXValidator_InvalidStruct(t *testing.T) {
	validator := XValidator{Validator: validator.New()}
	data := ValidatedUserStruct{
		Name:  "",
		Age:   16,
		Email: "invalid-email",
	}

	validationErrors := validator.Validate(data)

	assert.NotEmpty(t, validationErrors, "Expected validation errors for invalid struct")
	assert.Len(t, validationErrors, 3, "Expected three validation errors")

	assert.Equal(t, "Name", validationErrors[0].FailedField)
	assert.Equal(t, "required", validationErrors[0].Tag)

	assert.Equal(t, "Age", validationErrors[1].FailedField)
	assert.Equal(t, "gte", validationErrors[1].Tag)

	assert.Equal(t, "Email", validationErrors[2].FailedField)
	assert.Equal(t, "email", validationErrors[2].Tag)
}

func TestXValidator_NoValidationTags(t *testing.T) {
	validator := XValidator{Validator: validator.New()}
	data := NoValidationStruct{
		Name: "Some Name",
	}

	validationErrors := validator.Validate(data)
	assert.Empty(t, validationErrors, "Expected no validation errors for struct without validation tags")
}
