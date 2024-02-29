package synlinq

import (
	"api/strukt"
)

func AddPtr(ip string, rdns any) (*strukt.Outputipv4rdns, error) {
	var output strukt.Outputipv4rdns
	err := client.postWithParams("/rdns/by-ip", map[string]any{"ip": ip, "rdns": rdns},
		&output)
	if err != nil {
		return nil, err
	}

	return &output, nil
}
