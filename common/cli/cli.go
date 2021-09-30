package cli

import (
	"bufio"
	"context"
	"currency-converter/common/api/handlers/validations"
	"currency-converter/common/config"
	"currency-converter/common/models"
	"currency-converter/common/services"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

func RunCLI() (err error) {
	webServerConfig, err := config.FromEnv()
	if err != nil {
		return err
	}

	err = services.Initialize(webServerConfig.Service)
	if err != nil {
		log.Printf("an error occurred while initializing services: %s", err.Error())
		return err
	}

	filePath := flag.String("file", "", "Full path to json file")
	targetCurrency := flag.String("target-currency", "", "Currency file should be converted to")

	flag.Parse()

	if *filePath == "" {
		return fmt.Errorf("file is required. Pass in as --file=<path-to-file>")
	}
	if *targetCurrency == "" {
		return fmt.Errorf("target-currency is required. Pass in as --target-currency=EUR")
	}

	err = validations.ValidateCurrencyCode(*targetCurrency)
	if err != nil {
		return err
	}

	f, err := os.Open(*filePath)
	if err != nil {
		log.Printf("an error occurred while reading file with error: %s", err.Error())
		return err
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	//allows to reads line by line
	s := bufio.NewScanner(f)

	ctx := context.Background()

	for s.Scan() {
		var conversion models.ValueWithCurrency
		err := json.Unmarshal([]byte(s.Text()), &conversion)
		if err != nil {
			log.Printf("an error occurred while unmarshalling with error: %s", err.Error())
			return err
		}

		conversion = services.GetConverted(ctx, *targetCurrency, conversion)

		outputBytes, err := json.Marshal(conversion)
		if err != nil {
			log.Printf("an error occurred while marshalling output with error: %s", err.Error())
			continue
		}
		fmt.Println(string(outputBytes))
	}

	return nil
}
