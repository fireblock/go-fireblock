package common

import (
	"encoding/json"
	"fmt"
)

// ProviderData provider struct
type ProviderData struct {
	UID      string `json:"uid"`
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

func errorCAC(msg string) (string, string, string, error) {
	return "", "", "", NewFBKError(msg, InvalidCard)
}

// CheckAllCard verify sha3(card)==cardId then the proofs and return useruid, pubkey, ktype
// check the key in the blockchain
func CheckAllCard(card, cardID string) (string, string, string, error) {
	// decode json
	var providers []ProviderData
	err := json.Unmarshal([]byte(card), &providers)
	if err != nil {
		msg := fmt.Sprintf(`Invalid Json: %s`, card)
		return errorCAC(msg)
	}
	// check useruid
	fireblockCE := getCardElement("fireblock", providers)
	if fireblockCE == nil {
		return errorCAC(`missing fireblock provider`)
	}
	useruid := fireblockCE.UID
	if len(fireblockCE.UID) < 4 || len(fireblockCE.UID) > 12 {
		msg := fmt.Sprintf(`invalid fireblock provider: %s`, fireblockCE.UID)
		return errorCAC(msg)
	}
	// check if it's the expected card
	cardUID := Keccak256(card)
	if cardUID != cardID {
		return errorCAC(`Fake card detected! Contact Us!`)
	}
	pgpCE := getCardElement("pgp", providers)
	ecdsaCE := getCardElement("ecdsa", providers)
	ethCE := getCardElement("eth", providers)
	count := 0
	if pgpCE != nil {
		count++
	}
	if ecdsaCE != nil {
		count++
	}
	if ethCE != nil {
		count++
	}
	var fingerprint string
	var keyuid string
	var ktype string
	if count != 1 {
		return errorCAC(`No key in card!`)
	} else if pgpCE != nil {
		fingerprint = pgpCE.UID
		keyuid = PGPToB32(pgpCE.UID)
		ktype = "pgp"
	} else if ecdsaCE != nil {
		fingerprint = ecdsaCE.UID
		keyuid = ECDSAToB32(ecdsaCE.UID)
		ktype = "ecdsa"
	} else if ethCE != nil {
		fingerprint = ethCE.UID
		keyuid = ethCE.UID
		ktype = "eth"
	} else {
		return errorCAC(`Multiple keys in card!`)
	}
	if len(keyuid) != 66 {
		msg := fmt.Sprintf(`Invalid keyuid %s!`, keyuid)
		return errorCAC(msg)
	}
	pubkey, errk := HTTPKey(keyuid)
	if errk != nil {
		return errorCAC(`No public key found!`)
	}
	// check twitter
	if twitterCE := getCardElement("twitter", providers); twitterCE != nil {
		if !CheckTwitter(twitterCE.Proof, twitterCE.UID, useruid, fingerprint) {
			return errorCAC(`Invalid twitter!`)
		}
	}
	// check github
	if githubCE := getCardElement("github", providers); githubCE != nil {
		if !CheckGithub(githubCE.Proof, githubCE.UID, useruid, fingerprint) {
			return errorCAC(`Invalid github!`)
		}
	}
	// check HTTPS
	if httpsCE := getCardElement("https", providers); httpsCE != nil {
		if !CheckHTTPS(httpsCE.Proof, httpsCE.UID, useruid, fingerprint) {
			return errorCAC(`Invalid website!`)
		}
	}
	// check Linkedin
	if linkedinCE := getCardElement("linkedin", providers); linkedinCE != nil {
		if !CheckLinkedin(linkedinCE.Proof, linkedinCE.UID, useruid, fingerprint) {
			return errorCAC(`Invalid linkedin`)
		}
	}
	return useruid, pubkey, ktype, nil
}

// CheckCard verify sha3(card)==cardId then the proofs and return useruid, keyuid, ktype
// Don't check the key in the blockchain
func CheckCard(card, cardID string) (string, string, string, error) {
	// decode json
	var providers []ProviderData
	err := json.Unmarshal([]byte(card), &providers)
	if err != nil {
		msg := fmt.Sprintf(`Invalid Json: %s`, card)
		return errorCAC(msg)
	}
	// check useruid
	fireblockCE := getCardElement("fireblock", providers)
	if fireblockCE == nil {
		return errorCAC(`missing fireblock provider`)
	}
	useruid := fireblockCE.UID
	if len(fireblockCE.UID) < 4 || len(fireblockCE.UID) > 12 {
		msg := fmt.Sprintf(`invalid fireblock provider: %s`, fireblockCE.UID)
		return errorCAC(msg)
	}
	// check if it's the expected card
	cardUID := Keccak256(card)
	if cardUID != cardID {
		return errorCAC(`Fake card detected! Contact Us!`)
	}
	pgpCE := getCardElement("pgp", providers)
	ecdsaCE := getCardElement("ecdsa", providers)
	ethCE := getCardElement("eth", providers)
	count := 0
	if pgpCE != nil {
		count++
	}
	if ecdsaCE != nil {
		count++
	}
	if ethCE != nil {
		count++
	}
	var fingerprint string
	var keyuid string
	var ktype string
	if count != 1 {
		return errorCAC(`No key in card!`)
	} else if pgpCE != nil {
		fingerprint = pgpCE.UID
		keyuid = PGPToB32(pgpCE.UID)
		ktype = "pgp"
	} else if ecdsaCE != nil {
		fingerprint = ecdsaCE.UID
		keyuid = ECDSAToB32(ecdsaCE.UID)
		ktype = "ecdsa"
	} else if ethCE != nil {
		fingerprint = ethCE.UID
		keyuid = ethCE.UID
		ktype = "eth"
	} else {
		return errorCAC(`Multiple keys in card!`)
	}
	if len(keyuid) != 66 {
		msg := fmt.Sprintf(`Invalid keyuid %s!`, keyuid)
		return errorCAC(msg)
	}
	// check twitter
	if twitterCE := getCardElement("twitter", providers); twitterCE != nil {
		if !CheckTwitter(twitterCE.Proof, twitterCE.UID, useruid, fingerprint) {
			return errorCAC(`Invalid twitter!`)
		}
	}
	// check github
	if githubCE := getCardElement("github", providers); githubCE != nil {
		if !CheckGithub(githubCE.Proof, githubCE.UID, useruid, fingerprint) {
			return errorCAC(`Invalid github!`)
		}
	}
	// check HTTPS
	if httpsCE := getCardElement("https", providers); httpsCE != nil {
		if !CheckHTTPS(httpsCE.Proof, httpsCE.UID, useruid, fingerprint) {
			return errorCAC(`Invalid website!`)
		}
	}
	// check Linkedin
	if linkedinCE := getCardElement("linkedin", providers); linkedinCE != nil {
		if !CheckLinkedin(linkedinCE.Proof, linkedinCE.UID, useruid, fingerprint) {
			return errorCAC(`Invalid linkedin`)
		}
	}
	return useruid, keyuid, ktype, nil
}
