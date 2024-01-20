package strukt

import (
	"time"
)

type DdosResponse struct {
	Status string `json:"status"`
	Result []struct {
		Analyzer    string      `json:"analyzer"`
		L4Dynamic   interface{} `json:"l4_dynamic"`
		L4Permanent bool        `json:"l4_permanent"`
		L7Only      bool        `json:"l7_only"`
		L7Permanent bool        `json:"l7_permanent"`
		LastChanged struct {
			Timestamp time.Time `json:"@timestamp"`
			Time      int       `json:"time"`
		} `json:"last_changed"`
	} `json:"routing"`
}
type DdosResponse2 struct {
	Status  string `json:"status"`
	Routing struct {
		Prefix      string      `json:"prefix"`
		Analyzer    string      `json:"analyzer"`
		Blackhole   interface{} `json:"blackhole"`
		Flowspec    interface{} `json:"flowspec"`
		L4Dynamic   interface{} `json:"l4_dynamic"`
		L4Permanent bool        `json:"l4_permanent"`
		L7Only      bool        `json:"l7_only"`
		L7Permanent bool        `json:"l7_permanent"`
		LastChanged struct {
			Time      int       `json:"time"`
			Timestamp time.Time `json:"@timestamp"`
		} `json:"last_changed"`
	} `json:"routing"`
}
type Routing struct {
	L4permanent bool `json:"l4_permanent"`
	L7permanent bool `json:"l7_permanent"`
	L7only      bool `json:"l7_only"`
}
type Incidents struct {
	Items []struct {
		Timestamp   time.Time `json:"@timestamp"`
		Cluster     string    `json:"cluster"`
		IP          string    `json:"ip"`
		Mbps        string    `json:"mbps"`
		Method      string    `json:"method"`
		Mode        string    `json:"mode"`
		Packetsize  string    `json:"packetsize"`
		Pps         string    `json:"pps"`
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
		UUID string `json:"uuid"`
	} `json:"items"`
}
type AddThreshold struct {
	Prefix string `json:"prefix"`
	Mbit   int    `json:"mbit"`
	Kpps   int    `json:"kpps"`
}
type GetThreshold struct {
	Status string `json:"status"`
	Result []struct {
		Kpps   int    `json:"kpps"`
		Mbit   int    `json:"mbit"`
		Prefix string `json:"prefix"`
		UUID   string `json:"uuid"`
	} `json:"result"`
}
type PathIncident struct {
	Status string `json:"status"`
	Items  []struct {
		Host    string    `json:"host"`
		Reason  string    `json:"reason"`
		Start   time.Time `json:"start"`
		End     time.Time `json:"end"`
		PeakBps struct {
			Value     int       `json:"value"`
			Timestamp time.Time `json:"timestamp"`
		} `json:"peak_bps"`
		PeakPps struct {
			Value     int       `json:"value"`
			Timestamp time.Time `json:"timestamp"`
		} `json:"peak_pps"`
	} `json:"items"`
}
type DDOSLayer4 struct {
	Prefix      string `json:"prefix" validate:"required"`
	L4permanent bool   `json:"l4_permanent"`
	L7permanent bool   `json:"l7_permanent"`
	L7only      bool   `json:"l7_only"`
}
type Pref struct {
	Prefix string `json:"prefix" validate:"required"`
}

type AddRule struct {
	Source        string      `json:"source"`
	SourceASN     any         `json:"source_asn,omitempty"`
	Destination   string      `json:"destination"`
	Protocol      string      `json:"protocol"`
	DstPort       int         `json:"dst_port"`
	SrcPort       interface{} `json:"src_port"`
	RateLimiter   bool        `json:"rate_limiter"`
	RateLimiterID string      `json:"rate_limiter_id"`
	Block         bool        `json:"block"`
	Allow         bool        `json:"allow"`
	Priority      string      `json:"priority"`
	Comment       string      `json:"comment"`
}
type DeleteRule struct {
	Acknowledged bool `json:"acknowledged"`
}
type GetFilters struct {
	Filters []struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Settings struct {
			Addr             string `json:"addr,omitempty"`
			Port             int    `json:"port,omitempty"`
			MaxConnPps       int    `json:"max_conn_pps,omitempty"`
			VoicePort        int    `json:"voice_port,omitempty"`
			FileTransferPort int    `json:"file_transfer_port,omitempty"`
			OverridePort     int    `json:"override_port,omitempty"`
			Strict           bool   `json:"strict,omitempty"`
			Cache            bool   `json:"cache,omitempty"`
			AcceptQueries    bool   `json:"accept_queries,omitempty"`
			TCPPort          int    `json:"tcp_port,omitempty"`
			MultiIPSupport   bool   `json:"multi_ip_support,omitempty"`
			Protocol         string `json:"protocol,omitempty"`
			LoginPort        int    `json:"login_port,omitempty"`
		} `json:"settings"`
	} `json:"filters"`
}
type AvailableFilter struct {
	Filters []struct {
		Name        string        `json:"name"`
		Label       string        `json:"label"`
		Description string        `json:"description"`
		Fields      []interface{} `json:"fields"`
	} `json:"filters"`
}
type AddFilter struct {
	Addr             string `json:"addr"`
	Port             string `json:"port,omitempty"`
	VoicePort        string `json:"voice_port,omitempty"`
	FileTransferPort string `json:"file_transfer_port,omitempty"`
	OverridePort     string `json:"override_port,omitempty"`
	Strict           bool   `json:"strict,omitempty"`
	Cache            bool   `json:"cache,omitempty"`
	AcceptQueries    bool   `json:"accept_queries,omitempty"`
	MaxConnPps       string `json:"max_conn_pps,omitempty"`
	TCPPort          string `json:"tcp_port,omitempty"`
	MultiIPSupport   bool   `json:"multi_ip_support,omitempty"`
	Protocol         string `json:"protocol,omitempty"`
	LoginPort        string `json:"login_port,omitempty"`
}
type CreateRatelimit struct {
	PacketsPerSecond int    `json:"packets_per_second"`
	PerDestination   bool   `json:"per_destination"`
	Comment          string `json:"comment"`
}
