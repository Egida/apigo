package pdns

import (
	"api/model"
	"fmt"
)

type RRSet struct {
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	TTL        int       `json:"ttl,omitempty"`
	ChangeType string    `json:"changetype,omitempty"`
	Records    []Record  `json:"records,omitempty"`
	Comments   []Comment `json:"comments,omitempty"`
}

type Record struct {
	Content  string `json:"content"`
	Disabled bool   `json:"disabled,omitempty"`
}

type Comment struct {
	Content    string `json:"content"`
	Account    string `json:"account,omitempty"`
	ModifiedAt int    `json:"modified_at,omitempty"`
}

func AddRRSets(zoneID string, input model.AddRecodInput) error {
	body, err := client.patch("/servers/localhost/zones/"+zoneID, input)
	if err != nil {
		fmt.Println("error_body:", string(body))
		return err
	}

	return nil
}
