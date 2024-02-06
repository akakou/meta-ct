package metact

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Notification struct {
	Entry []struct {
		ID      string `json:"id"`
		Time    int    `json:"time"`
		Changes []struct {
			Field string      `json:"field"`
			Value Certificate `json:"value"`
		} `json:"changes"`
	} `json:"entry"`
	Object string `json:"object"`
}

func (ct *MetaCT) WebHookCertificates(c echo.Context) ([]Certificate, error) {
	var notification Notification

	if err := c.Bind(&notification); err != nil {
		return nil, fmt.Errorf("error parsing payload: %v", err)
	}

	if notification.Object != "certificate_transparency" {
		return nil, fmt.Errorf("invalid object: %s", notification.Object)
	}

	certs := []Certificate{}

	for _, entry := range notification.Entry {
		for _, change := range entry.Changes {
			certs = append(certs, change.Value)
		}
	}

	return certs, nil
}
