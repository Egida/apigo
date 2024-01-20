package cloudns

type Zoneresponse struct {
	Status            string `json:"status"`
	StatusDescription string `json:"statusDescription"`
}

func AddZone(domain string) (*Zoneresponse, error) {
	var output Zoneresponse
	//map[string]any{"ns[]": []string{"ns1.dnic.icu", "ns2.dnic.icu"}, "domain-name": domain, "zone-type": "master"}
	err := client.getWithParams("/dns/register.json", map[string]any{"ns[]": []string{"ns1.dnic.icu", "ns2.dnic.icu"}, "domain-name": domain, "zone-type": "master"},
		&output)
	if err != nil {
		return nil, err
	}

	return &output, nil
}
func DeleteCloudZone(domain string) (*Zoneresponse, error) {
	var output Zoneresponse
	//map[string]any{"ns[]": []string{"ns1.dnic.icu", "ns2.dnic.icu"}, "domain-name": domain, "zone-type": "master"}
	err := client.getWithParams("/dns/delete.json", map[string]any{"domain-name": domain},
		&output)
	if err != nil {
		return nil, err
	}

	return &output, nil
}
func AddCloudrecord(domain string, r_type string, host string, ttl int, record string, prio int) (*Zoneresponse, error) {
	var output Zoneresponse
	//map[string]any{"ns[]": []string{"ns1.dnic.icu", "ns2.dnic.icu"}, "domain-name": domain, "zone-type": "master"}
	err := client.getWithParams("/dns/add-record.json", map[string]any{"domain-name": domain, "record-type": r_type, "host": host, "record": record, "ttl": ttl, "priority": prio},
		&output)
	if err != nil {
		return nil, err
	}

	return &output, nil
}
