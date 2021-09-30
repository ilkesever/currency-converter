package handlers

import (
	"context"
	"currency-converter/common/api/handlers/validations"
	conversionModels "currency-converter/common/models"
	"currency-converter/common/services"
	"encoding/json"
	"log"
	"net/http"
)

//Handles incoming http req for conversion
func ConversionHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	requestBody := conversionModels.ConversionRequestDto{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Printf("Error while unmarshalling currency conversion error %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = validations.ValidateCurrencyCode(requestBody.TargetCurrency)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var ConvertedValues []conversionModels.ValueWithCurrency

	for _, conversion := range requestBody.Data {
		err := validations.ValidateCurrencyCode(conversion.Currency)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = validations.ValidateAmount(conversion.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		conversion = services.GetConverted(ctx, requestBody.TargetCurrency, conversion)

		ConvertedValues = append(ConvertedValues, conversion)
	}

	responseBodyBytes, err := json.Marshal(ConvertedValues)
	if err != nil {
		log.Printf("Error while marshalling currency conversion error %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(responseBodyBytes)
	if err != nil {
		log.Printf("Error while marshalling healthcheck error %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
