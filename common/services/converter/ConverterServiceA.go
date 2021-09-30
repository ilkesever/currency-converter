package converter

import (
	"context"
	httpService "currency-converter/common/services/http"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Service struct {
	HttpService httpService.HttpInterface
	Config      *Config
}

type Config struct {
	HostName string `required:"true" split_words:"true" `
	APIKey   string `required:"true" split_words:"true" `
}

func NewConverterServiceA(config *Config, HttpService httpService.HttpInterface) *Service {
	s := &Service{
		Config:      config,
		HttpService: HttpService,
	}

	return s
}

func (converterService *Service) GetCurrencyRate(ctx context.Context, conversion string) (float64, error) {
	url := fmt.Sprintf("%s?q=%s&compact=ultra&apiKey=%s", converterService.Config.HostName, conversion, converterService.Config.APIKey)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("An error occurred while building http request with error: %s", err.Error())
		return 0, err
	}

	timeout, cancel := converterService.HttpService.WithTimeout(ctx)
	defer cancel()

	bodyBytes, statusCode, err := converterService.HttpService.Do(timeout, request)
	if err != nil {
		log.Printf("An error occurred while making api call to %s with error: %s", request.URL.String(), err.Error())
		return 0, err
	}

	if statusCode < 200 && statusCode > 299 {
		log.Printf("Received an unexpected status code %d with body: %s", statusCode, string(bodyBytes))
		return 0, err
	}

	response := map[string]float64{}
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		log.Printf("An error occurred while unmarshalling response body with err %s and body %s", err.Error(), string(bodyBytes))
		return 0, err
	}

	if val, ok := response[conversion]; ok {
		return val, nil
	}
	err = fmt.Errorf("conversion rate was not found. Expected response object is not correct from body %s", string(bodyBytes))
	log.Printf(err.Error())
	return 0, err
}
