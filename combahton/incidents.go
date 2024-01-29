package combahton

import (
	"fmt"
	"time"

	jsoniter "github.com/json-iterator/go"
)

type Incident struct {
	UUID        string    `json:"uuid"`
	IP          string    `json:"ip"`
	Mode        string    `json:"mode"`
	Custom      string    `json:"custom"`
	Cluster     string    `json:"cluster"`
	Method      string    `json:"method"`
	Mbps        string    `json:"mbps"`
	Pps         string    `json:"pps"`
	Packetsize  string    `json:"packetsize"`
	Time        int       `json:"time"`
	Timestamp   time.Time `json:"@timestamp"`
	SampleCount int       `json:"sample_count"`
	Samples     []struct {
		Bytes      int    `json:"bytes"`
		IcmpCode   int    `json:"icmp_code"`
		IcmpType   int    `json:"icmp_type"`
		IPDst      string `json:"ip_dst"`
		IPProtocol int    `json:"ip_protocol"`
		IPSrc      string `json:"ip_src"`
		IPTTL      int    `json:"ip_ttl"`
		PortDst    int    `json:"port_dst"`
		PortSrc    int    `json:"port_src"`
		TCPFlags   int    `json:"tcp_flags"`
		Time       int    `json:"time"`
		VlanDst    int    `json:"vlan_dst"`
		VlanSrc    int    `json:"vlan_src"`
	} `json:"samples"`
}

func GetIncidents(ip string) ([]Incident, error) {
	body, err := client.get("/antiddos/incidents/" + ip)
	if err != nil {
		return nil, err
	}

	var response Response[[]Incident]
	err = jsoniter.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %s", err)
	}

	return response.Result, nil
}
