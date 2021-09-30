package services

import (
	"context"
	conversionModels "currency-converter/common/models"
	currency_converter "currency-converter/common/services/converter"
	httpModule "currency-converter/common/services/http"
	"fmt"
	"log"
)

var service CurrencyConverterService

type CurrencyConverterService struct {
	httpService      httpModule.HttpInterface
	converterService currency_converter.ConverterInterface
	config           *ServiceConfig
}

type ServiceConfig struct {
	HTTPConfig               *httpModule.Config         `required:"true" split_words:"true"`
	CurrencyConversionConfig *currency_converter.Config `required:"true" split_words:"true"`
}

//Initialize required services for given configuration
func Initialize(c *ServiceConfig) (err error) {

	service = CurrencyConverterService{
		config: c,
	}

	service.httpService = httpModule.NewHttpServiceA(c.HTTPConfig)

	service.converterService = currency_converter.NewConverterServiceA(c.CurrencyConversionConfig, service.httpService)

	return nil
}

//using a very simple cache, should be improved before prod.
var cachedConversions = make(map[string]float64, 0)

func GetConverted(ctx context.Context, convertTo string, valueToBeConverted conversionModels.ValueWithCurrency) (convertedValue conversionModels.ValueWithCurrency) {
	conversionID := fmt.Sprintf("%s_%s", valueToBeConverted.Currency, convertTo)

	if cachedConversion, ok := cachedConversions[conversionID]; ok {
		convertedValue = conversionModels.ValueWithCurrency{
			Currency: convertTo,
			Value:    cachedConversion * valueToBeConverted.Value,
		}
	} else {
		rate, err := service.converterService.GetCurrencyRate(ctx, conversionID)
		if err != nil {
			// here to handle if connection fails, should act as customer desires
			log.Printf("an error occurred while trying to get conversion %s", conversionID)
			convertedValue = conversionModels.ValueWithCurrency{
				Currency: "Error",
				Value:    0,
			}
		} else {
			cachedConversions[conversionID] = rate
			convertedValue = conversionModels.ValueWithCurrency{
				Currency: convertTo,
				Value:    rate * valueToBeConverted.Value,
			}
		}
	}

	return convertedValue
}
