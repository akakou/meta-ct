package metactapi

import (
	"fmt"
	"net/url"
)

type SubscribeList struct {
	Data []struct {
		Id     string `json:"id"`
		Domain string `json:"domain"`
	} `json:"data"`
	Paging struct {
		Cursors struct {
			Before string `json:"before"`
			After  string `json:"after"`
		} `json:"cursors"`
	} `json:"paging"`
}

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

	return result, nil
}

func (api *MetaCTApi) Unsubscribe(domain string) (*OnlyStatusResp, error) {
	u := fmt.Sprintf(SUBSCRIBED_DOMAINS_API, api.AppId)

	data := url.Values{}
	data.Set("unsubscribe", domain)
	data.Set("access_token", api.AccessToken)

	result, err := fetchAPI[OnlyStatusResp]("POST", u, data)

	if err != nil {
		return nil, fmt.Errorf("error unsubscribing: %v", err)
	}

	if !result.Success {
		return nil, fmt.Errorf("error unsubscribing: %v", result)
	}

	return result, nil
}

func (api *MetaCTApi) SubscribeList() (*SubscribeList, error) {
	u := fmt.Sprintf(SUBSCRIBED_DOMAINS_API, api.AppId)

	data := url.Values{}
	data.Set("fields", "domain")
	data.Set("access_token", api.AccessToken)

	result, err := fetchAPI[SubscribeList]("GET", u, data)

	if err != nil {
		return nil, fmt.Errorf("error getting certificates: %v", err)
	}

	return result, nil
}
