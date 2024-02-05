package metact

const API_BASE = "https://graph.facebook.com"

const (
	SUBSCRIBED_DOMAINS_API = API_BASE + "/%v/subscribed_domains"
	CERTFICATES_API        = API_BASE + "/certificates"
)

var ALL_CERTIFICATE_FIELDS = []string{
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
}
