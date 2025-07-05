package validation

import (
	"reflect"
	"strings"
)

type customValidation struct{}

const (
	enumKey = "enum"
)

func (re customValidation) SetTagName() {
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "" {
			return field.Name
		}

		if name == "-" {
			return ""
		}

		return name
	})
}
