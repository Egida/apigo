package dynadot

type SearchResponse struct {
	SearchResponse struct {
		ResponseCode  string
		SearchResults []struct {
			DomainName string
			Available  string
		}
	}
}

func Search(domain string) (*SearchResponse, error) {
	var resp SearchResponse
	err := client.sendParams("search", &resp, map[string]string{"domain0": domain})
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
