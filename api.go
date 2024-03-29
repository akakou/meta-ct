package metact

import (
	"fmt"
	"net/url"
	"strings"
)

func (ct *MetaCT) Subscribe(domain string) error {
	u := fmt.Sprintf(SUBSCRIBED_DOMAINS_API, ct.AppId)

	data := url.Values{}
	data.Set("subscribe", domain)
	data.Set("access_token", ct.AccessToken)

	result, err := fetchAPI[SubscribeResp]("POST", u, data)

	if err != nil {
		return fmt.Errorf("error subscribing: %v", err)
	}

	if !result.Success {
		return fmt.Errorf("error subscribing: %v", result)
	}

	return nil
}

func (ct *MetaCT) Unsubscribe(domain string) error {
	u := fmt.Sprintf(SUBSCRIBED_DOMAINS_API, ct.AppId)

	data := url.Values{}
	data.Set("unsubscribe", domain)
	data.Set("access_token", ct.AccessToken)

	result, err := fetchAPI[SubscribeResp]("POST", u, data)

	if err != nil {
		return fmt.Errorf("error unsubscribing: %v", err)
	}

	if !result.Success {
		return fmt.Errorf("error unsubscribing: %v", result)
	}

	return nil
}

func (ct *MetaCT) SubscribeList() (*SubscribeListResp, error) {
	u := fmt.Sprintf(SUBSCRIBED_DOMAINS_API, ct.AppId)

	data := url.Values{}
	data.Set("fields", "domain")
	data.Set("access_token", ct.AccessToken)

	result, err := fetchAPI[SubscribeListResp]("GET", u, data)

	if err != nil {
		return nil, fmt.Errorf("error getting certificates: %v", err)
	}

	return result, nil
}

func (ct *MetaCT) Certificates(query string, fields []string) (*CertificatesResp, error) {
	f := strings.Join(fields, ",")

	data := url.Values{}
	data.Set("access_token", ct.AccessToken)
	data.Set("query", query)
	data.Set("fields", f)

	result, err := fetchAPI[CertificatesResp]("GET", CERTFICATES_API, data)

	if err != nil {
		return nil, fmt.Errorf("error getting certificates: %v", err)
	}

	return result, nil
}
