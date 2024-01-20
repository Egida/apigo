package active

import (
	//"encoding/json"
	"fmt"
	"net/url"
	"sync"
	"time"

	jsoniter "github.com/json-iterator/go"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
)

const (
	baseURL = "https://path-api.active-servers.com"
)

var client *Client

type Client struct {
	baseURL    string
	httpClient *fasthttp.Client
	token      string
	username   string
	password   string
	mu         sync.Mutex
}

func NewClient() *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: &fasthttp.Client{},
		username:   viper.GetString("app.activepathusername"),
		password:   viper.GetString("app.activepathpassword"),
	}
}

func Init() error {
	client = NewClient()
	_, err := client.refreshToken()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(30 * time.Minute)
	go func() {
		for range ticker.C {
			_, err := client.refreshToken()
			if err != nil {
				log.Error().Err(err).Msg("cannot refresh active-servers token")
			}
		}
	}()

	return nil
}

func (c *Client) Token() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.token
}

func (c *Client) refreshToken() (string, error) {
	token, err := c.auth()
	if err != nil {
		return "", err
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	log.Info().
		Str("module", "active-servers").
		Str("token", token).
		Msg("setting new auth token")
	c.token = token

	return "", nil
}

func (c *Client) auth() (string, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	form := url.Values{}
	form.Add("username", c.username)
	form.Add("password", c.password)
	form.Add("grant_type", "password")

	req.SetRequestURI(c.baseURL + "/token")
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBody([]byte(form.Encode()))

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := c.httpClient.Do(req, resp)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		log.Error().
			Bytes("body", resp.Body()).
			Int("status", resp.StatusCode()).
			Str("url", req.URI().String()).
			Msg("request failed")
		return "", fmt.Errorf("request failed with status code: %d", resp.StatusCode())
	}

	var response struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
	}

	if err := jsoniter.Unmarshal(resp.Body(), &response); err != nil {
		return "", err
	}

	return response.AccessToken, nil
}

func (c *Client) do(method, endpoint string, payload any) ([]byte, error) {
	url := c.baseURL + endpoint

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetMethod(method)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Token())

	if payload != nil {
		data, err := jsoniter.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal JSON: %w", err)
		}

		req.Header.Set("Content-Type", "application/json")
		req.SetBody(data)
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := c.httpClient.Do(req, resp); err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if resp.StatusCode() == fasthttp.StatusInternalServerError {
		_, err := c.refreshToken()
		if err != nil {
			return nil, fmt.Errorf("token refresh failed: %w", err)
		}
		log.Error().
			Bytes("body", resp.Body()).
			Int("status", resp.StatusCode()).
			Str("url", req.URI().String()).
			Str("method", method).
			Msg("request failed")
		return resp.Body(), fmt.Errorf("upstream request failed with status 500, please try again")
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		log.Error().
			Bytes("body", resp.Body()).
			Int("status", resp.StatusCode()).
			Str("url", req.URI().String()).
			Str("method", method).
			Msg("request failed")
		return resp.Body(), fmt.Errorf("request failed with status code: %d", resp.StatusCode())
	}

	return resp.Body(), nil
}

func (c *Client) get(endpoint string) ([]byte, error) {
	return c.do(fasthttp.MethodGet, endpoint, nil)
}

func (c *Client) post(endpoint string, payload interface{}) ([]byte, error) {
	return c.do(fasthttp.MethodPost, endpoint, payload)
}

func (c *Client) delete(endpoint string) ([]byte, error) {
	return c.do(fasthttp.MethodDelete, endpoint, nil)
}

func (c *Client) put(endpoint string, payload interface{}) ([]byte, error) {
	return c.do(fasthttp.MethodPut, endpoint, payload)
}
