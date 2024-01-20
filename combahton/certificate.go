package combahton

import (
	"encoding/base64"
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

type Certificate struct {
	UUID        string `json:"uuid,omitempty"`
	Customer    string `json:"customer"`
	Service     string `json:"service"`
	IPAddress   string `json:"ipaddress"`
	Domain      string `json:"domain"`
	Certificate string `json:"certificate"`
	PrivateKey  string `json:"privatekey"`
	Hash        struct {
		SHA1 string `json:"sha1"`
	} `json:"hash,omnitempty"`
	Validity int `json:"validity"`
}

func ListCertificates() ([]Certificate, error) {
	body, err := client.get("/antiddos/certificate")
	if err != nil {
		return nil, err
	}
	var response Response[[]Certificate]
	err = jsoniter.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %s", err)
	}

	return response.Result, nil
}

func CreateCertificate(ipAddress, domain, certificate, privateKey string, validity int) (Certificate, error) {
	if !isBase64(certificate) {
		certificate = base64.StdEncoding.EncodeToString([]byte(certificate))
	}

	if !isBase64(privateKey) {
		privateKey = base64.StdEncoding.EncodeToString([]byte(privateKey))
	}

	payload := Certificate{
		IPAddress:   ipAddress,
		Domain:      domain,
		Certificate: certificate,
		PrivateKey:  privateKey,
		Validity:    validity,
	}

	body, err := client.post("/antiddos/certificate", payload)
	if err != nil {
		fmt.Println("error_body:", string(body))
		return Certificate{}, fmt.Errorf("failed to create certificate: %s", err)
	}

	var response Response[Certificate]
	err = jsoniter.Unmarshal(body, &response)
	if err != nil {
		return Certificate{}, fmt.Errorf("failed to unmarshal response: %s", err)
	}

	return response.Result, nil
}

func DeleteCertificate(uuid string) error {
	body, err := client.delete("/antiddos/certificate/" + uuid)
	if err != nil {
		fmt.Println("error_body:", string(body))
		return fmt.Errorf("failed to delete certificate: %s", err)
	}

	return nil
}

func isBase64(s string) bool {
	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}
