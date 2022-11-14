package commonUtils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// GetErrorMap returns errors in a map[string]string
func GetErrorMap(paramType reflect.Type, err error, mapType ...string) map[string]string {
	respErr := make(map[string]string)
	if len(mapType) == 0 {
		mapType = []string{"query"}
	}
	switch err.(type) {
	case validator.ValidationErrors:
		for _, validationError := range err.(validator.ValidationErrors) {
			field, _ := paramType.Elem().FieldByName(validationError.Field())
			fieldName, _ := field.Tag.Lookup(mapType[0])
			respErr[fieldName] = errorMapMessage(validationError)
		}
	default:
		respErr["Unknown field error"] = err.Error()
	}
	return respErr
}

func errorMapMessage(validationError validator.FieldError) string {
	switch validationError.Tag() {
	case "required":
		return "This field is required"
	case "numeric":
		return "This field should be numeric"
	case "oneof":
		replacer := *strings.NewReplacer(" ", "/")
		return fmt.Sprintf("Value must be one of (%s)", replacer.Replace(validationError.Param()))
	case "min":
		return fmt.Sprintf("Cannot be less than %s", validationError.Param())
	case "max":
		return fmt.Sprintf("Cannot be more than %s", validationError.Param())
	default:
		return "Unknown Error"
	}
}
