package cloudns

import (
	"crypto/tls"
	"fmt"
	"net/url"

	jsoniter "github.com/json-iterator/go"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
)

var client *Client

type Client struct {
	baseUrl    string
	httpClient *fasthttp.Client
	apiKey     string
	apiPass    string
}

func Init() {
	client = &Client{
		baseUrl: viper.GetString("app.cloudnsserver"),
		httpClient: &fasthttp.Client{
			TLSConfig: &tls.Config{InsecureSkipVerify: true},
		},
		apiKey:  viper.GetString("app.cloudnsauthid"),
		apiPass: viper.GetString("app.cloudnsauthpass"),
	}
}
func (c *Client) getWithParams(path string, params map[string]interface{}, output interface{}) error {
	values := url.Values{}
	values.Add("auth-id", c.apiKey)
	values.Add("auth-password", c.apiPass)

	for key, val := range params {
		if arr, ok := val.([]string); ok {
			for _, item := range arr {
				values.Add(key, item)
			}
		} else {
			values.Add(key, fmt.Sprintf("%v", val))
		}
	}

	url, err := url.ParseRequestURI(c.baseUrl + path)
	if err != nil {
		return err
	}
	url.RawQuery = values.Encode()

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url.String())
	req.Header.SetMethod(fasthttp.MethodGet)
	req.Header.Set("Accept", "application/json")

	resp := fasthttp.AcquireResponse()
	err = c.httpClient.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	if err != nil {
		return err
	}
	defer fasthttp.ReleaseResponse(resp)

	if err := jsoniter.Unmarshal(resp.Body(), &output); err != nil {
		log.Error().Err(err).Bytes("body", resp.Body()).Msg("cannot umnarshal json")
		return fmt.Errorf("cannot unmarshal json output: %s", err)
	}

	return nil
}
