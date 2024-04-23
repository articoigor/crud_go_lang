package validator

import (
	"encoding/json"
	"errors"

	error_map "github.com/articoigor/crud_go_lang/src/configuration/error_mapping"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validationError error) *error_map.CurrError {
	var jsonError *json.UnmarshalTypeError
	var jsonErrorvalidation validator.ValidationErrors

	if errors.As(validationError, &jsonError) {
		return error_map.NewErrorBadRequest("Invalid field type")
	} else if errors.As(validationError, &jsonErrorvalidation) {
		errorsCauses := []error_map.Cause{}

		for _, e := range validationError.(validator.ValidationErrors) {
			cause := error_map.Cause{
				Cause: e.Translate(transl),
				Field: e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return error_map.NewErrorValidationBadRequest("Some fields are invalid", errorsCauses)
	} else {
		return error_map.NewErrorBadRequest("Error trying to convert fields")
	}
}
