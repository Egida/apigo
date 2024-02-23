package pdns

import (
	"api/model"
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

type Zone struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Type             string   `json:"type,omitempty"`
	URL              string   `json:"url"`
	Kind             string   `json:"kind"`
	RRsets           []RRSet  `json:"rrsets,omitempty"`
	Serial           int      `json:"serial"`
	NotifiedSerial   int      `json:"notified_serial"`
	EditedSerial     int      `json:"edited_serial"`
	Masters          []string `json:"masters"`
	DNSSEC           bool     `json:"dnssec"`
	NSEC3Param       string   `json:"nsec3param,omitempty"`
	NSEC3Narrow      bool     `json:"nsec3narrow,omitempty"`
	Presigned        bool     `json:"presigned,omitempty"`
	SOAEdit          string   `json:"soa_edit,omitempty"`
	SOAEditAPI       string   `json:"soa_edit_api,omitempty"`
	APIRectify       bool     `json:"api_rectify,omitempty"`
	Zone             string   `json:"zone,omitempty"`
	Catalog          string   `json:"catalog,omitempty"`
	Account          string   `json:"account,omitempty"`
	NameServers      []string `json:"nameservers,omitempty"`
	MasterTSIGKeyIDs []string `json:"master_tsig_key_ids,omitempty"`
	SlaveTSIGKeyIDs  []string `json:"slave_tsig_key_ids,omitempty"`
}
type ZoneCreate struct {
	Account          string        `json:"account"`            
	APIRectify       bool          `json:"api_rectify"`        
	Dnssec           bool          `json:"dnssec"`             
	EditedSerial     int64         `json:"edited_serial"`      
	ID               string        `json:"id"`                 
	Kind             string        `json:"kind"`               
	LastCheck        int64         `json:"last_check"`         
	MasterTsigKeyIDS []interface{} `json:"master_tsig_key_ids"`
	Masters          []interface{} `json:"masters"`            
	Name             string        `json:"name"`               
	NotifiedSerial   int64         `json:"notified_serial"`    
	Nsec3Narrow      bool          `json:"nsec3narrow"`        
	Nsec3Param       string        `json:"nsec3param"`         
	Rrsets           []Rrset       `json:"rrsets"`             
	Serial           int64         `json:"serial"`             
	SlaveTsigKeyIDS  []interface{} `json:"slave_tsig_key_ids"` 
	SOAEdit          string        `json:"soa_edit"`           
	SOAEditAPI       string        `json:"soa_edit_api"`       
	URL              string        `json:"url"`                
}

type Rrset struct {
	Comments []interface{} `json:"comments"`
	Name     string        `json:"name"`    
	Records  []Record      `json:"records"` 
	TTL      int64         `json:"ttl"`     
	Type     string        `json:"type"`    
}

type Record struct {
	Content  string `json:"content"` 
	Disabled bool   `json:"disabled"`
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

func Add(input model.AddZoneInput) (ZoneCreate, error) {
	body, err := client.post("servers/localhost/zones", input)
	if err != nil {
		fmt.Println("error_body:", string(body))
		return ZoneCreate{}, err
	}

	var response Zone
	err = jsoniter.Unmarshal(body, &response)
	if err != nil {
		return ZoneCreate{}, err
	}

	return response, nil
}

func RemoveZone(zoneID string) error {
	body, err := client.delete("/servers/localhost/zones/" + zoneID)
	if err != nil {
		fmt.Println("error_body:", string(body))
		return err
	}

	return nil
}
