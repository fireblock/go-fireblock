package main

// CardInfoProvider provider
type CardInfoProvider struct {
	UID    string `json:"uid,omitempty"`
	Proof  string `json:"proof,omitempty"`
	Status string `json:"status,omitempty"`
}

// CardInfo card info
type CardInfo struct {
	UID       string `json:"uid"`
	Txt       string `json:"txt"`
	Signature string `json:"signature"`
	Date      int64  `json:"date"`
	Status    string `json:"status,omitempty"`
	Twitter   CardInfoProvider
	Github    CardInfoProvider
	Linkedin  CardInfoProvider
	HTTPS     CardInfoProvider
}

// KeyInfo key
type KeyInfo struct {
	KeyUID    string `json:"keyuid"`
	KType     string `json:"ktype"`
	Pubkey    string `json:"pubkey"`
	State     int64  `json:"state"`
	Signature string `json:"signature,omitemtpy"`
	Date      int64  `json:"date,omitempty"`
	Status    string `json:"status,omitempty"`
}

// CertificateInfo certificate
type CertificateInfo struct {
	Hash              string `json:"hash"`
	Signature         string `json:"signature"`
	Metadata          string `json:"metadata"`
	MetadataSignature string `json:"metadataSignature"`
	Batch             string `json:"batch"`
	Date              int64  `json:"date"`
	Status            string `json:"status,omitempty"`
}
