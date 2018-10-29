package main

type CardInfoProvider struct {
	UID    string
	Proof  string
	Status string
}

type CardInfo struct {
	UID       string `json:"uid"`
	Txt       string `json:"txt"`
	Signature string `json:"signature"`
	Date      int64  `json:"date"`
	Status    string
	Twitter   CardInfoProvider
	Github    CardInfoProvider
	Linkedin  CardInfoProvider
	HTTPS     CardInfoProvider
}

type KeyInfo struct {
	KeyUID    string `json:"keyuid"`
	KType     string `json:"ktype"`
	Pubkey    string `json:"pubkey"`
	State     int64  `json:"state"`
	Signature string `json:"signature"`
	Date      int64  `json:"date"`
	Status    string
}

type CertificateInfo struct {
	Hash              string `json:"hash"`
	Signature         string `json:"signature"`
	Metadata          string `json:"metadata"`
	MetadataSignature string `json:"metadataSignature"`
	Batch             string `json:"batch"`
	Date              int64  `json:"date"`
	Status            string
}
