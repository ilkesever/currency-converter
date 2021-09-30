package server

import (
	"currency-converter/common/config"
	"currency-converter/common/services"
	"fmt"
	"log"
	"net/http"
)

//main server struct
type Server struct {
	Configuration *config.WebServerConfig
	Router        *Router
}

//creates a new server with config
func NewServer(config *config.WebServerConfig) *Server {
	server := &Server{
		Configuration: config,
		Router:        NewRouter(),
	}

	return server
}

//RunApi initializes the server
func RunApi() (err error) {
	webServerConfig, err := config.FromEnv()
	if err != nil {
		return err
	}

	log.Printf("Starting HTTP server on port %s", webServerConfig.Port)

	err = services.Initialize(webServerConfig.Service)
	if err != nil {
		log.Printf("an error occurred while initializing services: %s", err.Error())
		return err
	}

	server := NewServer(webServerConfig)
	server.Router.InitializeRoutes()

	if err := http.ListenAndServe(fmt.Sprintf("%v:%v", "", webServerConfig.Port), *server.Router); err != nil {
		panic(err)
	}

	return nil
}
