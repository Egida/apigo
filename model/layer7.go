package model

type Layer7Input struct {
	IPAddress   string `json:"ipaddress" form:"ipaddress" validate:"required"`
	Domain      string `json:"domain" form:"domain" validate:"required"`
	Certificate string `json:"certificate"`
	PrivateKey  string `json:"privatekey"`
}
