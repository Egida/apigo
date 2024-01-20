package combahton

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"

	"github.com/spf13/viper"
)

type Vhost struct {
	UUID      string    `json:"uuid,omitempty"`
	Ipaddress string    `json:"ipaddress"`
	Domain    string    `json:"domain"`
	TLS       TLS       `json:"tls"`
	Challenge Challenge `json:"challenge"`
}

type TLS struct {
	Certificate string `json:"certificate"`
}
type Requests struct {
	Limit    int    `json:"limit"`
	Mode     string `json:"mode"`
	Template string `json:"template"`
	Timeout  int    `json:"timeout"`
}
type Challenge struct {
	Mode     string   `json:"mode"`
	Template string   `json:"template"`
	Requests Requests `json:"requests"`
}

func ListVhost() ([]Vhost, error) {
	body, err := client.get("/antiddos/vhost")
	if err != nil {
		return nil, err
	}

	var response Response[[]Vhost]
	err = jsoniter.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %s", err)
	}

	return response.Result, nil
}

func CreateVhost(ipAddress, domain, certificateUUID string) (Vhost, error) {
	payload := Vhost{
		Ipaddress: ipAddress,
		Domain:    domain,
		TLS: TLS{
			Certificate: certificateUUID,
		},
		Challenge: Challenge{
			Mode:     "captcha",
			Template: viper.GetString("app.layer7template"),
			Requests: Requests{
				Limit:    25,
				Mode:     "captcha",
				Template: viper.GetString("app.layer7template"),
				Timeout:  300,
			},
		},
	}

	body, err := client.post("/antiddos/vhost", payload)
	if err != nil {
		fmt.Println("error_body:", string(body))
		return Vhost{}, fmt.Errorf("failed to create vhost: %s", err)
	}

	var response Response[Vhost]
	err = jsoniter.Unmarshal(body, &response)
	if err != nil {
		return Vhost{}, fmt.Errorf("failed to unmarshal response: %s", err)
	}

	return response.Result, nil
}

func DeleteVhost(uuid string) error {
	body, err := client.delete("/antiddos/vhost/" + uuid)
	if err != nil {
		fmt.Println("error_body:", string(body))
		return fmt.Errorf("failed to delete vhost: %s", err)
	}

	return nil
}
