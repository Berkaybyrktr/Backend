package Api

import (
	"github.com/go-playground/validator/v10"
	"github/techschool/simplebank/Util"
)

var validCurrency validator.Func = func(fieldlevel validator.FieldLevel) bool {
	if currency, ok := fieldlevel.Field().Interface().(string); ok {
		return Util.IsSupportedCurrency(currency)
	}
	return false
}
