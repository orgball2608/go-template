package error

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type FieldError struct {
	Code  string `json:"code"`
	Field string `json:"field"`
}

type ErrorCodeValidate struct {
	HTTPCode    int          `json:"-"`
	Type        ErrorType    `json:"type"`
	FieldErrors []FieldError `json:"fieldErrors,omitempty"`
}

func CreateErrCodeValidate(ctx context.Context, input interface{}, err error) ErrorCodeValidate {
	errorCodeValidate := ErrorCodeValidate{
		HTTPCode:    http.StatusBadRequest,
		Type:        ErrorTypeValidate,
		FieldErrors: make([]FieldError, 0),
	}
	validationErrors, _ := err.(validator.ValidationErrors)
	for _, validationError := range validationErrors {
		field := lowerCaseFieldName(validationError.Namespace())
		errorCodeValidate.FieldErrors = append(errorCodeValidate.FieldErrors, FieldError{
			Code: fmt.Sprintf("%s_%s-%s-%s",
				ErrorTypeValidate,
				MapCodeValidate[validationError.ActualTag()].Code,
			),
			Field: field,
		})
	}
	return errorCodeValidate
}

func lowerCaseFieldName(field string) string {
	arrayString := strings.Split(field, ".")
	var newArrayString []string
	for i, v := range arrayString {
		if i == 0 {
			continue
		}
		bts := []byte(v)
		lc := bytes.ToLower([]byte{bts[0]})
		rest := bts[1:]
		newArrayString = append(newArrayString, string(bytes.Join([][]byte{lc, rest}, nil)))
	}
	return strings.Join(newArrayString, ".")
}
