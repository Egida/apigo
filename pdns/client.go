package pdns

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

var client *Client

type Client struct {
	baseURL    string
	httpClient *http.Client
	apiKey     string
}

func NewClient() *Client {
	return &Client{
		baseURL:    viper.GetString("app.powerdnsserver"),
		httpClient: &http.Client{},
		apiKey:     viper.GetString("app.powerdnskey"),
	}
}

func Init() {
	client = NewClient()
}

func (c *Client) get(endpoint string) ([]byte, error) {
	url := c.baseURL + endpoint

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %s", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-API-Key", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		return body, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	return body, nil
}

func (c *Client) post(endpoint string, payload interface{}) ([]byte, error) {
	url := c.baseURL + endpoint

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %s", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %s", err)
	}

	if resp.StatusCode != http.StatusCreated {
		return body, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	return body, nil
}

func (c *Client) delete(endpoint string) ([]byte, error) {
	url := c.baseURL + endpoint

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %s", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		return body, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	return body, nil
}

func (c *Client) patch(endpoint string, payload interface{}) ([]byte, error) {
	url := c.baseURL + endpoint

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %s", err)
	}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %s", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		return body, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	return body, nil
}
