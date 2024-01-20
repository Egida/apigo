package combahton

import (
	"api/strukt"
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

type Routing struct {
	Prefix string `json:"prefix"`

	Blackhole   *bool `json:"blackhole"`
	Flowspec    *bool `json:"flowspec"`
	L4Dynamic   *bool `json:"l4_dynamic"`
	L4Permanent *bool `json:"l4_permanent"`
	L7Only      any   `json:"l7_only"`
	L7Permanent *bool `json:"l7_permanent"`
	LastChanged struct {
		Time int `json:"time"`
	} `json:"last_changed"`
}

func GetRouting(ip string) ([]Routing, error) {
	body, err := client.get("/antiddos/routing?prefix=" + ip)
	if err != nil {
		return nil, err
	}

	var response Response[[]Routing]
	err = jsoniter.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %s", err)
	}

	return response.Result, nil
}

func AddRouting(input strukt.DDOSLayer4) (Routing, error) {
	body, err := client.put("/antiddos/routing", input)
	if err != nil {
		return Routing{}, err
	}

	var response Response[Routing]
	err = jsoniter.Unmarshal(body, &response)
	if err != nil {
		return Routing{}, fmt.Errorf("failed to unmarshal response: %s", err)
	}

	return response.Result, nil
}
