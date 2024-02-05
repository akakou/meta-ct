package metactapi

import (
	"fmt"
	"net/url"
)

type OnlyStatusResp struct {
	Success bool `json:"success"`
}

func (api *MetaCTApi) Subscribe(domain string) (*OnlyStatusResp, error) {
	u := fmt.Sprintf(SUBSCRIBED_DOMAINS_API, api.AppId)

	data := url.Values{}
	data.Set("subscribe", domain)
	data.Set("access_token", api.AccessToken)

	result, err := fetchAPI[OnlyStatusResp]("POST", u, data)

	if err != nil {
		return nil, fmt.Errorf("error subscribing: %v", err)
	}

	if !result.Success {
		return nil, fmt.Errorf("error subscribing: %v", result)
	}

	return result, err
}
