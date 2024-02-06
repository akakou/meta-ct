package main

import (
	"fmt"
	"net/http"

	"os"

	metact "github.com/akakou/meta-ct"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	appId := os.Getenv("META_APP_ID")
	accessToken := os.Getenv("META_ACCESS_TOKEN")
	testDomain := os.Getenv("TEST_DOMAIN")

	ct := metact.NewCT(appId, accessToken)
	err := ct.Subscribe(testDomain)

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	e.GET("/", func(c echo.Context) error {
		certs, err := ct.WebHookCertificates(c)
		if err != nil {
			c.Error(err)
		}

		fmt.Printf("certs: %v", certs)
		return c.String(http.StatusOK, "Hello, World!")
	})

	err = e.Start(":1323")
	e.Logger.Fatal(err)
}
