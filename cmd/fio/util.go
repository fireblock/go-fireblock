package main

type ProjectInfo struct {
	UID    string
	KType  string
	State  int64
	Status string
	Pubkey string
}

type KeyInfo struct {
	UID       string
	KType     string
	State     int64
	Status    string
	Pubkey    string
	Date      string
	Signature string
}

type CardInfoProvider struct {
	UID    string
	Proof  string
	Status string
}

type CardInfo struct {
	UID       string
	Txt       string
	Signature string
	Date      int64
	Status    string
	Twitter   CardInfoProvider
	Github    CardInfoProvider
	Linkedin  CardInfoProvider
	Https     CardInfoProvider
}

type CertificateInfo struct {
	Hash                       string
	Signature                  string
	Metadata                   string
	MetadataSignature          string
	Date                       int64
	DateISO                    string
	Tx                         int64
	CertificateSignatureStatus string
	CardSignatureStatus        string
	KeySignatureStatus         string
	MetadataSignatureStatus    string
}
