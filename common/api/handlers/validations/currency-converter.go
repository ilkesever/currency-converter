package validations

import "fmt"

var validCurrencyCodes = map[string]string{
	"USD": "USD",
	"EUR": "EUR",
	"JPY": "JPY",
}

// validator for currency code
func ValidateCurrencyCode(code string) error {
	if _, ok := validCurrencyCodes[code]; !ok {
		return fmt.Errorf("%s is not a valid currency code. Must be: USD, EUR, JPY", code)
	}
	return nil
}

// validator for value
func ValidateAmount(value float64) error {
	if value < 0 {
		return fmt.Errorf("value can not be negative!")
	}
	return nil
}
