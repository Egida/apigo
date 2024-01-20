package active

import (
	"api/strukt"
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

type Rate struct {
	RateLimiters []struct {
		PacketsPerSecond int    `json:"packets_per_second"`
		PerDestination   bool   `json:"per_destination"`
		Comment          string `json:"comment"`
		ID               string `json:"id"`
	} `json:"rate_limiters"`
}
type RateLimiter struct {
	ID               string `json:"id"`
	PacketsPerSecond int    `json:"packets_per_second"`
	PerDestination   bool   `json:"per_destination"`
	Comment          string `json:"comment"`
}

func GetRateLimiters() (*Rate, error) {
	body, err := client.get("/rate_limiters")
	if err != nil {
		return nil, err
	}

	var response Rate

	err = jsoniter.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}
func DeleteRatelimit(id string) (*strukt.DeleteRule, error) {
	body, err := client.delete("/rate_limiters/" + id)
	if err != nil {
		return nil, err
	}

	var response strukt.DeleteRule

	err = jsoniter.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}
func AddRatelimit(input strukt.CreateRatelimit) (*RateLimiter, error) {
	body, err := client.post("/rate_limiters", input)
	if err != nil {
		return nil, err
	}

	var response RateLimiter

	err = jsoniter.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}
