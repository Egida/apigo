package pdns

import (
	"api/model"
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

type Zone struct {
	Account          string        `json:"account"`
	APIRectify       bool          `json:"api_rectify"`
	Catalog          string        `json:"catalog"`
	Dnssec           bool          `json:"dnssec"`
	EditedSerial     int           `json:"edited_serial"`
	ID               string        `json:"id"`
	Kind             string        `json:"kind"`
	LastCheck        int           `json:"last_check"`
	MasterTsigKeyIds []interface{} `json:"master_tsig_key_ids"`
	Masters          []interface{} `json:"masters"`
	Name             string        `json:"name"`
	NotifiedSerial   int           `json:"notified_serial"`
	Nsec3Narrow      bool          `json:"nsec3narrow"`
	Nsec3Param       string        `json:"nsec3param"`
	Rrsets           []struct {
		Comments []interface{} `json:"comments"`
		Name     string        `json:"name"`
		Records  []struct {
			Content  string `json:"content"`
			Disabled bool   `json:"disabled"`
		} `json:"records"`
		TTL  int    `json:"ttl"`
		Type string `json:"type"`
	} `json:"rrsets"`
	Serial          int           `json:"serial"`
	SlaveTsigKeyIds []interface{} `json:"slave_tsig_key_ids"`
	SoaEdit         string        `json:"soa_edit"`
	SoaEditAPI      string        `json:"soa_edit_api"`
	URL             string        `json:"url"`
}

func ListZones() ([]Zone, error) {
	body, err := client.get("/servers/localhost/zones")
	if err != nil {
		fmt.Println("error_body:", string(body))
		return nil, err
	}

	var response []Zone
	err = jsoniter.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func GetZone(zoneID string) (Zone, error) {
	body, err := client.get("/servers/localhost/zones/" + zoneID)
	if err != nil {
		fmt.Println("error_body:", string(body))
		return Zone{}, err
	}

	var response Zone
	err = jsoniter.Unmarshal(body, &response)
	if err != nil {
		return Zone{}, err
	}

	return response, nil
}

func Add(input model.AddZoneInput) (Zone, error) {
	body, err := client.post("/servers/localhost/zones", input)
	if err != nil {
		fmt.Println("error_body:", string(body))
		return Zone{}, err
	}

	var response Zone
	err = jsoniter.Unmarshal(body, &response)
	if err != nil {
		return Zone{}, fmt.Errorf("failed to unmarshal response: %s", err)
	}

	return response, err
}

func RemoveZone(zoneID string) error {
	body, err := client.delete("/servers/localhost/zones/" + zoneID)
	if err != nil {
		fmt.Println("error_body:", string(body))
		return err
	}

	return nil
}
