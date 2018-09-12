package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

func extractFrom(txt string) (string, string, string, string) {
	// I,ellis2323 am cjFDSTdYZ3pN on Fireblock and use the 0x99090eae43316b2ba65ec52bcd5834a3e07edb2c pgp key (https://fireblock.io)
	reg, err := regexp.Compile(`I,([\w\-]+) am ([\w]+) on Fireblock and use the (0x\w+) (ecdsa|pgp|eth) key`)
	if err != nil {
		return "", "", "", ""
	}
	k := reg.FindStringSubmatch(txt)
	if len(k) != 5 {
		return "", "", "", ""
	}
	return k[2], k[1], k[4], k[3]
}

func extractFromHTTPS(txt string) (string, string, string) {
	// I am rJpoNe2zm on Fireblock and use the 0x99090eae43316b2ba65ec52bcd5834a3e07edb2c pgp key (https://fireblock.io )
	reg, err := regexp.Compile(`I am ([\w]+) on Fireblock and use the (0x\w+) (ecdsa|pgp|eth) key`)
	if err != nil {
		return "", "", ""
	}
	k := reg.FindStringSubmatch(txt)
	if len(k) != 4 {
		return "", "", ""
	}
	return k[1], k[3], k[2]
}

func getURLContent(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		msg := fmt.Sprintf("invalid url: %s", url)
		return "", NewFBKError(msg, NetworkError)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	s := buf.String()
	return s, nil
}

// CheckTwitter check twitter proof
func CheckTwitter(url, twUID, useruid, fingerprint string) bool {
	reg, err := regexp.Compile(`https://twitter.com/(\w{1,15})/status/`)
	if err != nil {
		return false
	}
	if !reg.MatchString(url) {
		return false
	}
	k := reg.FindStringSubmatch(url)
	if k[1] != twUID {
		return false
	}
	proof, err2 := getURLContent(url)
	if err2 != nil {
		return false
	}
	luid, twuid, _, fp := extractFrom(proof)
	if luid != useruid || twuid != twUID || fp != fingerprint {
		return false
	}
	return true
}

// OwnerData login
type OwnerData struct {
	Login string `json:"login"`
}

// GistData Gist information
type GistData struct {
	Owner  OwnerData       `json:"owner"`
	ForkOf json.RawMessage `json:"fork_of"`
}

// CheckGithub Check github proof
func CheckGithub(url, ghUID, useruid, fingerprint string) bool {
	reg, err := regexp.Compile(`^https://gist.github.com/(\w{1,15})/(\w+)`)
	if err != nil {
		return false
	}
	k := reg.FindStringSubmatch(url)
	if k[1] != ghUID {
		return false
	}
	// content
	newURL := `https://gist.githubusercontent.com/` + k[1] + `/` + k[2] + `/raw`
	proof, err4 := getURLContent(newURL)
	if err4 != nil {
		return false
	}
	luid, ghuid, _, fp := extractFrom(proof)
	if luid != useruid || ghuid != ghUID || fp != fingerprint {
		return false
	}
	return true
}

// CheckLinkedin check a linkedin proof
func CheckLinkedin(url, lkUID, useruid, fingerprint string) bool {
	proof, err := getURLContent(url)
	if err != nil {
		return false
	}
	reg, err2 := regexp.Compile(`https://\w+.linkedin.com/in/` + lkUID)
	if err2 != nil {
		return false
	}
	res := reg.FindAllString(proof, -1)
	if len(res) == 0 {
		return false
	}
	// check
	luid, lkuid, _, fp := extractFrom(proof)
	if luid != useruid || lkuid != lkUID || fp != fingerprint {
		return false
	}
	return true
}

// CheckHTTPS check website with a https proof
func CheckHTTPS(url, dnsUID, useruid, fingerprint string) bool {
	url2 := `^https://` + dnsUID + `/.well-known/` + fingerprint + `.txt$`
	reg, err2 := regexp.Compile(url2)
	if err2 != nil {
		return false
	}
	resR := reg.MatchString(url)
	if !resR {
		return false
	}
	proof, err := getURLContent(url)
	if err != nil {
		return false
	}
	luid, ktype, fp := extractFromHTTPS(proof)
	if ktype != "eth" && ktype != "ecdsa" && ktype != "pgp" {
		return false
	}
	if luid != useruid || fp != fingerprint {
		return false
	}
	return true
}
