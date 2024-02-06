package metact

type Certificate struct {
	Id                     string     `json:"id,omitempty"`
	AuthorityKeyIdentifier string     `json:"authority_key_identifier,omitempty"`
	BasicConstraints       string     `json:"basic_constraints,omitempty"`
	CertHashMd5            string     `json:"cert_hash_md5,omitempty"`
	CertHashSha1           string     `json:"cert_hash_sha1,omitempty"`
	CertHashSha256         string     `json:"cert_hash_sha256,omitempty"`
	CertificatePem         string     `json:"certificate_pem,omitempty"`
	Domains                []string   `json:"domains,omitempty"`
	ExtendedKeyUsage       string     `json:"extended_key_usage,omitempty"`
	Extensions             []KeyValue `json:"extensions,omitempty"`
	IssuerName             string     `json:"issuer_name,omitempty"`
	KeyUsage               string     `json:"key_usage,omitempty"`
	NotValidAfter          string     `json:"not_valid_after,omitempty"`
	NotValidBefore         string     `json:"not_valid_before,omitempty"`
	PublicKeyAlgorithm     string     `json:"public_key_algorithm,omitempty"`
	PublicKeyHashSha256    string     `json:"public_key_hash_sha256,omitempty"`
	PublicKeyPem           string     `json:"public_key_pem,omitempty"`
	PublicKeySize          int        `json:"public_key_size,omitempty"`
	PublicKeyValues        []KeyValue `json:"public_key_values,omitempty"`
	SerialNumber           string     `json:"serial_number,omitempty"`
	SignatureAlgorithm     string     `json:"signature_algorithm,omitempty"`
	SignatureValue         string     `json:"signature_value,omitempty"`
	SubjectKeyIdentifier   string     `json:"subject_key_identifier,omitempty"`
	SubjectName            string     `json:"subject_name,omitempty"`
	Version                int        `json:"version,omitempty"`
}

type Paging struct {
	Cursors struct {
		Before string `json:"before"`
		After  string `json:"after"`
	} `json:"causers"`
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type SubscribeResp struct {
	Success bool `json:"success"`
}

type SubscribeListResp struct {
	Data []struct {
		Id     string `json:"id"`
		Domain string `json:"domain"`
	} `json:"data"`
	Paging `json:"paging"`
}

type CertificatesResp struct {
	Data   []Certificate `json:"data"`
	Paging `json:"paging"`
}
