package metactapi

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func debugPrintf(tag string, data any) {
	debugPrint := os.Getenv("DEBUG_PRINT")
	if debugPrint == "1" {
		fmt.Printf("%v: %v\n", tag, data)
	}
}

const example_domain = "hoge.ochano.co"

func TestActual(t *testing.T) {
	appId := os.Getenv("META_APP_ID")
	accessToken := os.Getenv("META_ACCESS_TOKEN")

	api := NewCT(appId, accessToken)

	t.Run("subscribe", func(t *testing.T) {
		err := api.Subscribe(example_domain)
		assert.NoError(t, err)
	})

	t.Run("list", func(t *testing.T) {
		list, err := api.SubscribeList()

		assert.NoError(t, err)
		debugPrintf("list", list)
		assert.Equal(t, example_domain, list.Data[0].Domain)
	})

	t.Run("certificates", func(t *testing.T) {
		certs, err := api.Certificates(example_domain, ALL_CERTIFICATE_FIELDS)

		assert.NoError(t, err)
		debugPrintf("certs", certs)
		assert.Equal(t, example_domain, certs.Data[0].Domains[0])
	})

	t.Run("unsubscribe", func(t *testing.T) {
		err := api.Unsubscribe(example_domain)

		if err != nil {
			t.Fatalf("error unsubscribe: %v", err)
		}
	})
}
