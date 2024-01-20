package dynadot

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	jsoniter "github.com/json-iterator/go"

	"github.com/spf13/viper"
)

const (
	baseURL = "https://api.dynadot.com/api3.json"
)

var client *Client

type Client struct {
	baseUrl    string
	httpClient *http.Client
	apiKey     string
}

type ErrorResponse struct {
	Response struct {
		ResponseCode string
		Error        string
	}
}

func NewClient() *Client {
	return &Client{
		baseUrl:    baseURL,
		httpClient: &http.Client{},
		apiKey:     viper.GetString("app.dynapikey"),
	}
}

func Init() {
	client = NewClient()
}

func (c *Client) sendParams(command string, output interface{}, params map[string]string) error {
	values := url.Values{}
	values.Add("key", c.apiKey)
	values.Add("command", command)

	for key, val := range params {
		values.Add(key, val)
	}

	url, err := url.ParseRequestURI(c.baseUrl)
	if err != nil {
		return fmt.Errorf("failed to parse url: %s", err)
	}
	url.RawQuery = values.Encode()

	resp, err := c.httpClient.Get(url.String())
	if err != nil {
		return fmt.Errorf("request failed: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %s", err)
	}

	var r ErrorResponse
	if err := jsoniter.Unmarshal(body, &r); err != nil {
		return fmt.Errorf("failed to decode error response: %s", err)
	}

	if r.Response.ResponseCode == "-1" {
		return fmt.Errorf("received error from dynadot api: %s", r.Response.Error)
	}

	if err := jsoniter.Unmarshal(body, &output); err != nil {
		return fmt.Errorf("failed to decode response: %s", err)
	}

	return nil
}

func (c *Client) send(command string, output interface{}) error {
	return c.sendParams(command, &output, map[string]string{})
}
