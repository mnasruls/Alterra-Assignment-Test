package validation

import (
	"alterra/entities"
	"alterra/entities/web"
	"reflect"

	"github.com/go-playground/validator/v10"
)

/*
 * User Validation - Error Message
 * -------------------------------
 * Kumpulan custom error message yang ditampilkan
 * ke response berdasarkan struct field dan validate tagnya
 */
var userErrorMessages = map[string]string{
	"Name|required":        "Name field must be filled",
	"Email|required":       "Email field must be filled",
	"Email|email":          "Email field is not an email",
	"Password|required":    "Password field must be filled",
	"Gender|required":      "Gender field must be filled",
	"DOB|required":         "Date of birth field must be filled",
	"Address|required":     "Address must be filled",
	"PhoneNumber|required": "Phone number must be filled",
	"Role|required":        "Role must be filled",
}

/*
 * User Validation - Validate Create User Request
 * -------------------------------
 * Validasi user saat registrasi berdasarkan validate tag
 * yang ada pada user request dan file rules diatas
 */
func ValidateCreateUserRequest(validate *validator.Validate, userReq entities.CreateUserRequest) error {

	errors := []web.ValidationErrorItem{}

	validateUserStruct(validate, userReq, &errors)

	if len(errors) > 0 {
		return web.ValidationError{
			Code:    400,
			Message: "Validation error",
			Errors:  errors,
		}
	}
	return nil
}

func validateUserStruct(validate *validator.Validate, userReq entities.CreateUserRequest, errors *[]web.ValidationErrorItem) {
	err := validate.Struct(userReq)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field, _ := reflect.TypeOf(userReq).FieldByName(err.Field())
			*errors = append(*errors, web.ValidationErrorItem{
				Field: field.Tag.Get("form"),
				Error: userErrorMessages[err.Field()+"|"+err.ActualTag()],
			})
		}
	}
}
