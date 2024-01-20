package combahton

import (
	"encoding/base64"
	"fmt"
	"net/http"

	jsoniter "github.com/json-iterator/go"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
)

const (
	baseURL = "https://api.aurologic.com"
)

var client *Client

type Client struct {
	baseURL    string
	httpClient *fasthttp.Client
	username   string
	password   string
}

type Response[T any] struct {
	Status struct {
		Code  int    `json:"code"`
		Type  string `json:"type"`
		Error bool   `json:"error"`
	} `json:"status"`
	Paginated bool `json:"paginated,omitempty"`
	Count     int  `json:"count,omitempty"`
	PerPage   int  `json:"per_page,omitempty"`
	Page      int  `json:"page,omitempty"`
	Result    T    `json:"result"`
}

func NewClient() *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: &fasthttp.Client{},
		username:   viper.GetString("app.ddosusername"),
		password:   viper.GetString("app.ddospassword"),
	}
}

func Init() {
	client = NewClient()
}

func (c *Client) get(endpoint string) ([]byte, error) {
	url := c.baseURL + endpoint

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodGet)
	req.Header.Set("Accept", "application/json")
	req.Header.Set(
		"Authorization",
		"Basic "+base64.StdEncoding.EncodeToString([]byte(c.username+":"+c.password)),
	)

	resp := fasthttp.AcquireResponse()
	err := c.httpClient.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %s", err)
	}
	defer fasthttp.ReleaseResponse(resp)

	if resp.StatusCode() != http.StatusOK {
		log.Error().Bytes("body", resp.Body()).Int("status", resp.StatusCode()).Msg("request failed")
		return resp.Body(), fmt.Errorf("request failed with status code: %d", resp.StatusCode())
	}

	return resp.Body(), nil
}

func (c *Client) post(endpoint string, payload interface{}) ([]byte, error) {
	url := c.baseURL + endpoint

	data, err := jsoniter.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %s", err)
	}

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set(
		"Authorization",
		"Basic "+base64.StdEncoding.EncodeToString([]byte(c.username+":"+c.password)),
	)
	req.SetBody(data)

	resp := fasthttp.AcquireResponse()
	err = c.httpClient.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %s", err)
	}
	defer fasthttp.ReleaseResponse(resp)

	if resp.StatusCode() != http.StatusOK {
		log.Error().Bytes("body", resp.Body()).Int("status", resp.StatusCode()).Msg("request failed")
		return resp.Body(), fmt.Errorf("request failed with status code: %d", resp.StatusCode())
	}

	return resp.Body(), nil
}

func (c *Client) delete(endpoint string) ([]byte, error) {
	url := c.baseURL + endpoint

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodDelete)
	req.Header.Set("Accept", "application/json")
	req.Header.Set(
		"Authorization",
		"Basic "+base64.StdEncoding.EncodeToString([]byte(c.username+":"+c.password)),
	)

	resp := fasthttp.AcquireResponse()
	err := c.httpClient.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %s", err)
	}
	defer fasthttp.ReleaseResponse(resp)

	if resp.StatusCode() != http.StatusOK {
		log.Error().Bytes("body", resp.Body()).Int("status", resp.StatusCode()).Msg("request failed")
		return resp.Body(), fmt.Errorf("request failed with status code: %d", resp.StatusCode())
	}

	return resp.Body(), nil
}

func (c *Client) put(endpoint string, payload interface{}) ([]byte, error) {
	url := c.baseURL + endpoint

	data, err := jsoniter.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %s", err)
	}

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodPut)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set(
		"Authorization",
		"Basic "+base64.StdEncoding.EncodeToString([]byte(c.username+":"+c.password)),
	)
	req.SetBody(data)

	resp := fasthttp.AcquireResponse()
	err = c.httpClient.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %s", err)
	}
	defer fasthttp.ReleaseResponse(resp)

	if resp.StatusCode() != http.StatusOK {
		log.Error().Bytes("body", resp.Body()).Int("status", resp.StatusCode()).Msg("request failed")
		return resp.Body(), fmt.Errorf("request failed with status code: %d", resp.StatusCode())
	}

	return resp.Body(), nil
}
