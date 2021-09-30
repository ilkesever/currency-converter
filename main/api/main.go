package main

import (
	server "currency-converter/common/api"
	"log"
)

func main() {
	err := server.RunApi()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
