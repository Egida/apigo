package active

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

type Attack struct {
	Data []interface{} `json:"data"`
}

func GetIncidents(ip string) (*Attack, error) {
	body, err := client.get("/attack_history?host=" + ip)
	if err != nil {
		return nil, err
	}

	var response *Attack
	err = jsoniter.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return response, nil
}
