package utils

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func NewValidateStruct() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

func ParseErrorMsg(s interface{}, err error) string {
	for _, err := range err.(validator.ValidationErrors) {
		field, _ := reflect.TypeOf(s).FieldByName(err.StructField())
		msg := field.Tag.Get("errorMsg")
		return msg
	}

	return ""
}
