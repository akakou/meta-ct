package metact

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Notification struct {
	Entry []struct {
		ID            string   `json:"id"`
		ChangedFields []string `json:"changed_fields"`
		Time          int      `json:"time"`
		Certificate   `json:"certificate"`
	} `json:"entry"`
	Object string `json:"object"`
}

func (api *MetaCT) WebHookCertificates(c echo.Context) ([]Certificate, error) {
	var notification Notification
	if err := c.Bind(&notification); err != nil {
		return nil, fmt.Errorf("error parsing payload: %v", err)
	}

	if notification.Object != "subscribed_domain" {
		return nil, fmt.Errorf("invalid object: %s", notification.Object)
	}

	certs := []Certificate{}

	for _, entry := range notification.Entry {
		certs = append(certs, entry.Certificate)
	}

	return certs, nil
}
