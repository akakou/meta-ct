package metactapi

import (
	"fmt"
	"net/url"
)

func (api *MetaCTApi) Subscribe(domain string) (*SubscribeResp, error) {
	u := fmt.Sprintf(SUBSCRIBED_DOMAINS_API, api.AppId)

	data := url.Values{}
	data.Set("subscribe", domain)
	data.Set("access_token", api.AccessToken)

	result, err := fetchAPI[SubscribeResp]("POST", u, data)

	if err != nil {
		return nil, fmt.Errorf("error subscribing: %v", err)
	}

	if !result.Success {
		return nil, fmt.Errorf("error subscribing: %v", result)
	}

	return result, nil
}

func (api *MetaCTApi) Unsubscribe(domain string) (*SubscribeResp, error) {
	u := fmt.Sprintf(SUBSCRIBED_DOMAINS_API, api.AppId)

	data := url.Values{}
	data.Set("unsubscribe", domain)
	data.Set("access_token", api.AccessToken)

	result, err := fetchAPI[SubscribeResp]("POST", u, data)

	if err != nil {
		return nil, fmt.Errorf("error unsubscribing: %v", err)
	}

	if !result.Success {
		return nil, fmt.Errorf("error unsubscribing: %v", result)
	}

	return result, nil
}

func (api *MetaCTApi) SubscribeList() (*SubscribeListResp, error) {
	u := fmt.Sprintf(SUBSCRIBED_DOMAINS_API, api.AppId)

	data := url.Values{}
	data.Set("fields", "domain")
	data.Set("access_token", api.AccessToken)

	result, err := fetchAPI[SubscribeListResp]("GET", u, data)

	if err != nil {
		return nil, fmt.Errorf("error getting certificates: %v", err)
	}

	return result, nil
}
