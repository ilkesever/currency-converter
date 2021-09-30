package main

import (
	"currency-converter/common/cli"
	"log"
)

func main() {
	err := cli.RunCLI()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
