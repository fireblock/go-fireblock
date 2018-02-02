package common

import (
	"encoding/json"
	"fmt"

	"github.com/fireblock/go-fireblock/common/errors"
)

type ProviderData struct {
	Uid      string `json:"uid"`
	Provider string `json:"provider"`
	Date     string `json:"date"`
	Proof    string `json:"proof"`
}

func getCardElement(selector string, providers []ProviderData) *ProviderData {
	for i := 0; i < len(providers); i++ {
		if providers[i].Provider == selector {
			return &providers[i]
		}
	}
	return nil
}

// CheckAllCard verify sha3(card)==cardId then the proofs and return useruid, pubkey, ktype
func CheckAllCard(card, cardId string) (string, string, string, error) {
	// decode json
	var providers []ProviderData
	err := json.Unmarshal([]byte(card), &providers)
	if err != nil {
		msg := fmt.Sprintf(`Invalid Json: %s`, card)
		return "", "", "", errors.NewFBKError(msg, errors.InvalidCard)
	}
	// check useruid
	fireblockCE := getCardElement("fireblock", providers)
	if fireblockCE == nil {
		return "", "", "", errors.NewFBKError(`missing fireblock provider`, errors.InvalidCard)
	}
	useruid := fireblockCE.Uid
	if len(fireblockCE.Uid) < 4 || len(fireblockCE.Uid) > 12 {
		msg := fmt.Sprintf(`invalid fireblock provider: %s`, fireblockCE.Uid)
		return "", "", "", errors.NewFBKError(msg, errors.InvalidCard)
	}
	// check if it's the expected card
	cardUID := Keccak256(card)
	if cardUID != cardId {
		return "", "", "", errors.NewFBKError(`Fake card detected! Contact Us!`, errors.InvalidCard)
	}
	pgpCE := getCardElement("pgp", providers)
	ecdsaCE := getCardElement("ecdsa", providers)
	ethCE := getCardElement("eth", providers)
	count := 0
	if pgpCE != nil {
		count += 1
	}
	if ecdsaCE != nil {
		count += 1
	}
	if ethCE != nil {
		count += 1
	}
	var fingerprint string
	var keyuid string
	var ktype string
	if count != 1 {
		return "", "", "", errors.NewFBKError(`No key in card!`, errors.InvalidCard)
	} else if pgpCE != nil {
		fingerprint = pgpCE.Uid
		keyuid = PGPToB32(pgpCE.Uid)
		ktype = "pgp"
	} else if ecdsaCE != nil {
		fingerprint = ecdsaCE.Uid
		keyuid = ECDSAToB32(ecdsaCE.Uid)
		ktype = "ecdsa"
	} else if ethCE != nil {
		fingerprint = ethCE.Uid
		keyuid = ethCE.Uid
		ktype = "eth"
	} else {
		return "", "", "", errors.NewFBKError(`Multiple keys in card!`, errors.InvalidCard)
	}
	if len(keyuid) != 66 {
		msg := fmt.Sprintf(`Invalid keyuid %s!`, keyuid)
		return "", "", "", errors.NewFBKError(msg, errors.InvalidCard)
	}
	pubkey, errk := HTTPKey(keyuid)
	if errk != nil {
		return "", "", "", errors.NewFBKError(`No public key found!`, errors.InvalidCard)
	}
	// check twitter
	twitterCE := getCardElement("twitter", providers)
	if twitterCE != nil {
		if !CheckTwitter(twitterCE.Proof, twitterCE.Uid, useruid, fingerprint) {
			return "", "", "", errors.NewFBKError(`Invalid twitter!`, errors.InvalidCard)
		}
	}
	// check github
	githubCE := getCardElement("github", providers)
	if githubCE != nil {
		if !CheckGithub(githubCE.Proof, githubCE.Uid, useruid, fingerprint) {
			return "", "", "", errors.NewFBKError(`Invalid github!`, errors.InvalidCard)
		}
	}
	// check HTTPS
	httpsCE := getCardElement("https", providers)
	if httpsCE != nil {
		if !CheckHTTPS(httpsCE.Proof, httpsCE.Uid, useruid, fingerprint) {
			return "", "", "", errors.NewFBKError(`Invalid website!`, errors.InvalidCard)
		}
	}
	// check Linkedin
	linkedinCE := getCardElement("linkedin", providers)
	if linkedinCE != nil {
		if !CheckLinkedin(linkedinCE.Proof, linkedinCE.Uid, useruid, fingerprint) {
			return "", "", "", errors.NewFBKError(`Invalid linkedin`, errors.InvalidCard)
		}
	}
	return useruid, pubkey, ktype, nil
}
