package strukt

type GetAccounts struct {
	Data  []Item `json:"data"` 
	Links Links   `json:"links"`
	Meta  Meta    `json:"meta"` 
}

type Item struct {
	ID           int64   `json:"id"`            
	Username     string  `json:"username"`      
	Status       string  `json:"status"`        
	WgIP         string  `json:"wg_ip"`         
	WgPrivateKey string  `json:"wg_private_key"`
	WgPublicKey  string  `json:"wg_public_key"` 
	ExpiredAt    *string `json:"expired_at"`    
	Updated      string  `json:"updated"`       
	Created      string  `json:"created"`       
}

type Links struct {
	First string      `json:"first"`
	Last  string      `json:"last"` 
	Prev  interface{} `json:"prev"` 
	Next  interface{} `json:"next"` 
}

type Meta struct {
	CurrentPage int64  `json:"current_page"`
	From        int64  `json:"from"`        
	LastPage    int64  `json:"last_page"`   
	Path        string `json:"path"`        
	PerPage     int64  `json:"per_page"`    
	To          int64  `json:"to"`          
	Total       int64  `json:"total"`       
}

