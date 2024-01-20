package active

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"

	"api/strukt"
)

type Rule struct {
	ID            string `json:"id"`
	Source        string `json:"source"`
	SourceASN     any    `json:"source_asn,omitempty"`
	Destination   string `json:"destination"`
	Protocol      string `json:"protocol"`
	DstPort       int    `json:"dst_port"`
	SrcPort       any    `json:"src_port"`
	RateLimiter   bool   `json:"rate_limiter"`
	RateLimiterID any    `json:"rate_limiter_id"`
	Block         bool   `json:"block"`
	Allow         bool   `json:"allow"`
	Priority      bool   `json:"priority"`
	Comment       string `json:"comment"`
}

func GetRules(ip string) ([]*Rule, error) {
	body, err := client.get("/rules?destination=" + ip)
	if err != nil {
		return nil, err
	}

	var response struct {
		Rules []*Rule `json:"rules"`
	}

	err = jsoniter.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return response.Rules, nil
}

func GetRule(id string) (*Rule, error) {
	body, err := client.get("/rules/" + id)
	if err != nil {
		return nil, err
	}

	var response Rule

	err = jsoniter.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}

func AddRule(input strukt.AddRule) (*Rule, error) {
	body, err := client.post("/rules", input)
	if err != nil {
		return nil, err
	}

	var response Rule

	err = jsoniter.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}
func DeleteRule(id string) (*strukt.DeleteRule, error) {
	body, err := client.delete("/rules/" + id)
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
