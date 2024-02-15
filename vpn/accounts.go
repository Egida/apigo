package vpn

import (
	"api/strukt"
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

func GetAccounts() (*strukt.GetAccounts, error) {
	body, err := client.get("/accounts")
	if err != nil {
		return nil, err
	}

	var response *strukt.GetAccounts
	err = jsoniter.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %s", err)
	}

	return response, nil
}
