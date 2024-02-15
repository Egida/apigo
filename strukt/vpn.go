package strukt

type GetAccounts struct {
	Data []struct {
		ID           int    `json:"id"`
		Username     string `json:"username"`
		Status       string `json:"status"`
		WgIP         string `json:"wg_ip"`
		WgPrivateKey string `json:"wg_private_key"`
		WgPublicKey  string `json:"wg_public_key"`
		ExpiredAt    string `json:"expired_at"`
		Updated      string `json:"updated"`
		Created      string `json:"created"`
	} `json:"data"`
	Links struct {
		First string      `json:"first"`
		Last  string      `json:"last"`
		Prev  interface{} `json:"prev"`
		Next  interface{} `json:"next"`
	} `json:"links"`
	Meta struct {
		CurrentPage int    `json:"current_page"`
		From        int    `json:"from"`
		LastPage    int    `json:"last_page"`
		Path        string `json:"path"`
		PerPage     int    `json:"per_page"`
		To          int    `json:"to"`
		Total       int    `json:"total"`
	} `json:"meta"`
}
