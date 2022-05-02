package api

import (
	"github.com/cs-tungthanh/Bank_Golang/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check that currency is support or not
		return util.IsSupportCurrency(currency)
	}
	return false
}
