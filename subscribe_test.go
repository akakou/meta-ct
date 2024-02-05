package metactapi

import (
	"fmt"
	"os"
	"testing"
)

func TestActualSubscribe(t *testing.T) {
	appId := os.Getenv("META_APP_ID")
	accessToken := os.Getenv("META_ACCESS_TOKEN")

	api := NewCT(appId, accessToken)

	resp, err := api.Subscribe("example.com")

	if err != nil {
		t.Errorf("error subscribing: %v", err)
	}

	fmt.Printf("%v", resp)
}
