package pdns

import (
	"fmt"
	"net/http"

	jsoniter "github.com/json-iterator/go"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
)

var client *Client

type Client struct {
	baseURL    string
	httpClient *fasthttp.Client
	apikey     string
}

func NewClient() *Client {
	return &Client{
		baseURL:    viper.GetString("app.powerdnsserver"),
		httpClient: &fasthttp.Client{},
		apikey:     viper.GetString("app.powerdnskey"),
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
	req.Header.Set("X-API-Key", c.apikey)

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
	req.Header.Set("X-API-Key", c.apikey)
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
	req.Header.Set("X-API-Key", c.apikey)

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
	req.Header.Set("X-API-Key", c.apikey)
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
func (c *Client) patch(endpoint string, payload interface{}) ([]byte, error) {
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
	req.Header.Set("X-API-Key", c.apikey)
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
