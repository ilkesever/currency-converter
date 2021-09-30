package http

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

type Service struct {
	HTTPClient *http.Client
	Config     *Config
}

type Config struct {
	ClientTimeoutInSeconds int `required:"true" split_words:"true" default:"30"`
}

type APIError struct {
	StatusCode int
	Error      string
}

func NewHttpServiceA(config *Config) *Service {
	s := &Service{
		Config: config,
	}
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 10 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 10 * time.Second,

		ExpectContinueTimeout: 10 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			MinVersion:         tls.VersionTLS11,
		},
	}

	c := &http.Client{
		Timeout:   time.Second * time.Duration(config.ClientTimeoutInSeconds),
		Transport: transport,
	}

	s.HTTPClient = c
	return s
}

//Returns response body in bytes or error, with status code
func (httpService *Service) Do(ctx context.Context, request *http.Request) ([]byte, int, error) {

	request = request.WithContext(ctx)

	response, err := httpService.HTTPClient.Do(request)
	if err != nil {
		log.Printf("an error occurred while trying to make api request with error %s", err)
		return nil, http.StatusInternalServerError, err
	}
	if response == nil {
		log.Printf("received a nil response from %s", request.URL.String())
		err = fmt.Errorf("an error occurred trying to call http do request")
		return nil, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("an error occurred while trying to read response body from url %s with error %s", request.URL.String(), err.Error())
		return nil, http.StatusInternalServerError, err
	}

	return bodyBytes, response.StatusCode, nil
}

//adds timeout to context
func (httpService *Service) WithTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, time.Duration(httpService.Config.ClientTimeoutInSeconds)*time.Second)
}
