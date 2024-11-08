package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Struct que representa o validator.
type XValidator struct {
	Validator *validator.Validate
}

// Struct que representa um erro de validação do validator.
type ValidationError struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

// Função responsável por validar um struct utilizando o validator.
func (v XValidator) Validate(data interface{}) []ValidationError {
	validate := validator.New()

	validationErrors := []ValidationError{}

	// Validando o struct e retornando os erros se houver
	errs := validate.Struct(data)
	if errs != nil {
		// Verificando se o erro é do tipo ValidationErrors
		if _, ok := errs.(validator.ValidationErrors); !ok {
			fmt.Printf("erro validando struct: %v", errs)
			panic(errs)
		}

		// Iterando sobre os erros e adicionando ao array de erros
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ValidationError

			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Value = err.Value()
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}
