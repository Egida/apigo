package active

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"

	"api/strukt"
)

func GetFilter() (*strukt.GetFilters, error) {
	body, err := client.get("/filters")
	if err != nil {
		return nil, err
	}

	var response strukt.GetFilters

	err = jsoniter.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}
func AvailableFilters() (*strukt.AvailableFilter, error) {
	body, err := client.get("/filters/available")
	if err != nil {
		return nil, err
	}

	var response strukt.AvailableFilter

	err = jsoniter.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}
func AddFilter(input strukt.AddFilter, ft string) (*strukt.GetFilters, error) {
	body, err := client.post("/filters/"+ft, input)
	if err != nil {
		return nil, err
	}

	var response strukt.GetFilters

	err = jsoniter.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}
func DeleteFilter(id string, ft string) (*strukt.DeleteRule, error) {
	body, err := client.delete("/filters/" + ft + "/" + id)
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
