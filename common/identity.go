package common

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/fireblock/go-fireblock/common/errors"
)

func extract(txt string) (string, string) {
	// I am registered on fireblock.io cjFDSTdYZ3pN and use the 0x99090eae43316b2ba65ec52bcd5834a3e07edb2c pgp key (https://fireblock.io)
	reg, err := regexp.Compile(`I am registered on (Fireblock|fireblock.io) ([=\w]+) and use the (0x\w+)`)
	if err != nil {
		return "", ""
	}
	k := reg.FindStringSubmatch(txt)
	if len(k) != 4 {
		return "", ""
	}
	useruid := k[2]
	fp := k[3]
	b64, err2 := base64.StdEncoding.DecodeString(useruid)
	if err2 != nil {
		return "", ""
	}
	//	if (k === null) { return null }
	//	let useruid = utilFcts.btoa(res[1])
	//	return [useruid, res[2]]
	return string(b64[:]), fp
}

func getUrlContent(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		msg := fmt.Sprintf("invalid url: %s", url)
		return "", errors.NewFBKError(msg, errors.NetworkError)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	s := buf.String()
	return s, nil
}

func CheckTwitter(url, twUid, useruid, fingerprint string) bool {
	reg, err := regexp.Compile(`https://twitter.com/(\w{1,15})/status/`)
	if err != nil {
		return false
	}
	if !reg.MatchString(url) {
		return false
	}
	k := reg.FindStringSubmatch(url)
	if k[1] != twUid {
		return false
	}
	proof, err2 := getUrlContent(url)
	if err2 != nil {
		return false
	}
	luid, lfp := extract(proof)
	if luid != useruid || lfp != fingerprint {
		return false
	}
	return true
}

type OwnerData struct {
	Login string `json:"login"`
}

type GistData struct {
	Owner  OwnerData       `json:"owner"`
	ForkOf json.RawMessage `json:"fork_of"`
}

func CheckGithub(url, ghUid, useruid, fingerprint string) bool {
	reg, err := regexp.Compile(`^https://gist.github.com/(\w{1,15})/(\w+)`)
	if err != nil {
		return false
	}
	k := reg.FindStringSubmatch(url)
	if k[1] != ghUid {
		return false
	}
	// gist
	gistUrl := `https://api.github.com/gists/` + k[2]
	gistInfo, err3 := getUrlContent(gistUrl)
	if err3 != nil {
		return false
	}
	// check
	var data GistData
	err5 := json.Unmarshal([]byte(gistInfo), &data)
	if err5 != nil {
		return false
	}
	if data.Owner.Login != ghUid {
		return false
	}
	if len(data.ForkOf) > 0 {
		return false
	}
	// content
	newURL := `https://gist.githubusercontent.com/` + k[1] + `/` + k[2] + `/raw`
	proof, err4 := getUrlContent(newURL)
	if err4 != nil {
		return false
	}
	luid, lfp := extract(proof)
	if luid != useruid || lfp != fingerprint {
		return false
	}
	return true
}

func CheckLinkedin(url, lkUid, useruid, fingerprint string) bool {
	proof, err := getUrlContent(url)
	if err != nil {
		return false
	}
	reg, err2 := regexp.Compile(`href="https://www.linkedin.com/in/` + lkUid + `"`)
	if err2 != nil {
		return false
	}
	res := reg.FindAllString(proof, -1)
	if len(res) == 0 {
		return false
	}
	// check
	luid, lfp := extract(proof)
	if luid != useruid || lfp != fingerprint {
		return false
	}
	return true
}

func CheckHTTPS(url, dnsUid, useruid, fingerprint string) bool {
	reg, err2 := regexp.Compile(`^https://` + dnsUid + `/.fireblock/` + fingerprint + `.txt$`)
	if err2 != nil {
		return false
	}
	resR := reg.MatchString(url)
	if !resR {
		// check the old version
		reg, err2 = regexp.Compile(`^https://` + dnsUid + `/.fireblock/` + fingerprint + `$`)
		if err2 != nil {
			return false
		}
		resR := reg.MatchString(url)
		if !resR {
			return false
		}
	}
	proof, err := getUrlContent(url)
	if err != nil {
		return false
	}
	luid, lfp := extract(proof)
	if luid != useruid || lfp != fingerprint {
		return false
	}
	return true
}
