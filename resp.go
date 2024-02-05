package metactapi

type SubscribeResp struct {
	Success bool `json:"success"`
}

type Paging struct {
	Cursors struct {
		Before string `json:"before"`
		After  string `json:"after"`
	} `json:"causers"`
}

type SubscribeListResp struct {
	Data []struct {
		Id     string `json:"id"`
		Domain string `json:"domain"`
	} `json:"data"`
	Paging `json:"paging"`
}

type CertificatesResp struct {
	Data []struct {
		Id                     string `json:"id"`
		AuthorityKeyIdentifier string `json:"authority_key_identifier"`
		BasicConstraints       string `json:"basic_constraints"`
		CertHashMd5            string `json:"cert_hash_md5"`
		CertHashSha1           string `json:"cert_hash_sha1"`
		CertHashSha256         string `json:"cert_hash_sha256"`
		CertificatePem         string `json:"certificate_pem"`
		Domains                string `json:"domains"`
		ExtendedKeyUsage       string `json:"extended_key_usage"`
		Extensions             string `json:"extensions"`
		IssuerName             string `json:"issuer_name"`
		KeyUsage               string `json:"key_usage"`
		NotValidAfter          string `json:"not_valid_after"`
		NotValidBefore         string `json:"not_valid_before"`
		PublicKeyAlgorithm     string `json:"public_key_algorithm"`
		PublicKeyHashSha256    string `json:"public_key_hash_sha256"`
		PublicKeyPem           string `json:"public_key_pem"`
		PublicKeySize          string `json:"public_key_size"`
		PublicKeyValues        string `json:"public_key_values"`
		SerialNumber           string `json:"serial_number"`
		SignatureAlgorithm     string `json:"signature_algorithm"`
		SignatureValue         string `json:"signature_value"`
		SubjectKeyIdentifier   string `json:"subject_key_identifier"`
		SubjectName            string `json:"subject_name"`
		Version                string `json:"version"`
	} `json:"data"`
	Paging `json:"paging"`
}
