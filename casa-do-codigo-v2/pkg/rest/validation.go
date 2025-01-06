package rest

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

var v = validator.New()

type ErrorDetail map[string][]string

func ValidateJSON(req any) ErrorDetail {
	err := v.Struct(req)
	if err == nil {
		return nil
	}

	var ee map[string][]string = make(map[string][]string)
	for _, err := range err.(validator.ValidationErrors) {
		fld := reflect.ValueOf(req).Type()
		a, _ := fld.FieldByName(err.Field())

		var message string
		message = a.Tag.Get(err.Tag())
		if message == "" {
			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("%s is required", a.Tag.Get("json"))
			case "email":
				message = fmt.Sprintf("%s must be a valid email address", a.Tag.Get("json"))
			default:
				message = fmt.Sprintf("%s should be %s", a.Tag.Get("json"), err.Tag())
				if err.Param() != "" {
					message += fmt.Sprintf("=%s", err.Param())
				}

			}
		}

		ee[a.Tag.Get("json")] = append(ee[err.Field()], message)
	}

	return ee
}
