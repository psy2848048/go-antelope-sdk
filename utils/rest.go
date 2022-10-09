package utils

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	CodeOK = 0

	ErrStatusUnauthorized = "rest: unauthorized"
	ErrUnknown            = "rest: unknown error"

	HeaderContentType     = "Content-Type"
	HeaderContentTypeJson = "application/json"

	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"

	REST_TIMEOUT = 3 // in sec
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stderr, "xxx: ", log.Ldate|log.Ltime|log.Lshortfile)
}

type Method string

func RESTCallWithJson(endpoint string, method Method, param []byte) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Duration(time.Second * REST_TIMEOUT),
	}

	// TODO: will be random based
	domain := "https://api.eoseoul.io"
	url := strings.Join([]string{domain, endpoint}, "/")

	req, err := http.NewRequest(string(method), url, bytes.NewReader(param))
	if err != nil {
		return []byte{}, err
	}

	req.Header.Set(HeaderContentType, HeaderContentTypeJson)

	res, err := restCall(client, req)
	if os.IsTimeout(err) {
		// TODO: will be replaced into internal BP operation counter
		return []byte{}, err
	} else if err != nil {
		return []byte{}, err
	}

	return res, nil
}

func restCall(client *http.Client, req *http.Request) ([]byte, error) {
	res, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	err = restErrorHandler(res)
	if err != nil {
		return []byte{}, err
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			logger.Printf("res.body.close, %s", err)
		}
	}()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return resBody, nil
}

func restErrorHandler(res *http.Response) error {
	switch res.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusUnauthorized:
		return errors.New(ErrStatusUnauthorized)
	default:
		return errors.New(ErrUnknown)
	}
}
