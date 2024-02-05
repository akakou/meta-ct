package metactapi

import (
	"fmt"
	"os"
	"testing"
)

var debugPrint = false

func debugPrintf(tag string, data any) {
	if debugPrint {
		fmt.Printf("%v: %v\n", tag, data)
	}
}

const example_domain = "hoge.ochano.co"

func TestActual(t *testing.T) {
	appId := os.Getenv("META_APP_ID")
	accessToken := os.Getenv("META_ACCESS_TOKEN")

	api := NewCT(appId, accessToken)

	resp, err := api.Subscribe(example_domain)

	if err != nil {
		t.Fatalf("error subscribe: %v", err)
	}

	debugPrintf("subscribe", resp)

	list, err := api.SubscribeList()
	if err != nil {
		t.Fatalf("error subscribe list: %v", err)
	}

	debugPrintf("list", list)

	certs, err := api.Certificates(example_domain, []string{
		"authority_key_identifier",
		"basic_constraints",
		"cert_hash_md5",
		"cert_hash_sha1",
		"cert_hash_sha256",
		"certificate_pem",
		"domains",
		"extended_key_usage",
		"extensions",
		"issuer_name",
		"key_usage",
		"not_valid_after",
		"not_valid_before",
		"public_key_algorithm",
		"public_key_hash_sha256",
		"public_key_pem",
		"public_key_size",
		"public_key_values",
		"serial_number",
		"signature_algorithm",
		"signature_value",
		"subject_key_identifier",
		"subject_name",
		"version",
	})

	if err != nil {
		t.Fatalf("error certs: %v", err)
	}

	debugPrintf("certs", certs)

	resp, err = api.Unsubscribe(example_domain)

	if err != nil {
		t.Fatalf("error unsubscribe: %v", err)
	}

	debugPrintf("resp", resp)
}
