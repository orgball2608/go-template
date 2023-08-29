package validations

import (
	"github.com/dlclark/regexp2"

	"github.com/go-playground/validator/v10"
)

var CustomPassword validator.Func = func(fl validator.FieldLevel) bool {
	matched, _ := regexp2.MustCompile(`^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[a-zA-Z]).{8,30}$`, 0).MatchString(fl.Field().String())
	return matched
}
