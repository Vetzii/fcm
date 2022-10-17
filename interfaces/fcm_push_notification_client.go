package interfaces

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/vetzii/fcm/domain"
)

const (
	HOST         = "https://fcm.googleapis.com/fcm/send"
	BASE_API_URI = ""
	MIN_ATTEMPT  = 100 * time.Millisecond
	MAX_ATTEMPT  = 1 * time.Minute
)

type Config struct {
	ServerKey  string
	ApiUri     string
	HttpClient *http.Client
}

func FcmPushNotificationClient(config *Config) (*Config, error) {

	if config.ServerKey == "" {
		return nil, errors.New("API key is required")
	}

	//	instance
	cC := &Config{
		ServerKey:  config.ServerKey,
		ApiUri:     HOST,
		HttpClient: &http.Client{},
	}

	//	validate custom endpoint
	if config.ApiUri != "" {
		//	set custom endpoint
		cC.ApiUri = config.ApiUri
	}

	return cC, nil

}

func (cC *Config) Send(msg *domain.Message) (*domain.Response, error) {
	var data []byte
	var err error

	if err = msg.MessageValidate(); err != nil {
		return nil, err
	}

	if data, err = json.Marshal(msg); err != nil {
		return nil, err
	}

	return cC.submitProcess(data)
}

func (cC *Config) submitProcess(data []byte) (*domain.Response, error) {

	var err error
	var req *http.Request
	var resp *http.Response

	//	build request

	if req, err = http.NewRequest("POST", cC.ApiUri, bytes.NewBuffer(data)); err != nil {
		return nil, err
	}

	//	set headers
	req.Header.Add("Authorization", fmt.Sprintf("key=%s", cC.ServerKey))
	req.Header.Add("Content-Type", "application/json")

	// execute request
	if resp, err = cC.HttpClient.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// build return
	response := new(domain.Response)
	response.StatusCode = resp.StatusCode

	// check response status
	if resp.StatusCode != http.StatusOK {
		return response, errors.New(fmt.Sprintf("%d error: %s", resp.StatusCode, resp.Status))
	}

	if err = json.NewDecoder(resp.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}

func (cC *Config) SendWithRetry(msg *domain.Message, attemptsQuantity int) (*domain.Response, error) {

	var data []byte
	var err error
	var resp *domain.Response

	if err = msg.MessageValidate(); err != nil {
		return nil, err
	}

	data, err = json.Marshal(msg)

	if err != nil {
		return nil, err
	}

	resp = new(domain.Response)
	err = retry(func() error {
		var err error
		resp, err = cC.submitProcess(data)
		return err
	}, attemptsQuantity)

	return resp, err
}

func retry(fn func() error, attemptsQuantity int) error {

	var currentAttempts int
	var err error
	var backOff time.Duration
	var flagStatus bool
	var networkError net.Error

	for {

		if err = fn(); err == nil {
			return nil
		}

		if networkError, flagStatus = err.(net.Error); !flagStatus || !networkError.Temporary() {
			return err
		}

		//	set attemps counter
		currentAttempts++

		backOff = MIN_ATTEMPT * time.Duration(currentAttempts*currentAttempts)
		if currentAttempts > attemptsQuantity || backOff > MAX_ATTEMPT {
			return err
		}

		time.Sleep(backOff)
	}
}
