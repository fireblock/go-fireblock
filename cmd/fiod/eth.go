package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	fcommon "github.com/fireblock/go-fireblock/common"
	"github.com/fireblock/go-fireblock/contracts"
)

// EthKey key in store contract
type EthKey struct {
	Key     string
	Revoked bool
	Begin   int64
	End     int64
}

// EthCard card in store contract
type EthCard struct {
	Card    string
	KeyUID  string
	Upgrade string
}

// SuccessReturn success data
type SuccessReturn struct {
	Filename string `json:"filename"`
	Verified bool   `json:"verified"`
	Hash     string `json:"hash"`
	UserID   string `json:"user-id"`
	CardID   string `json:"card-id"`
}

// ErrorReturn error return struct
type ErrorReturn struct {
	Error    string `json:"error"`
	Detail   string `json:"detail"`
	Hash     string `json:"hash"`
	Filename string `json:"filename"`
	UserID   string `json:"user-id"`
	CardID   string `json:"card-id"`
}

func fioSuccess(filename, hash, userid, cardid string, verified, verbose bool) {
	if verbose {
		var r SuccessReturn
		r.Filename = filename
		r.Verified = verified
		r.Hash = hash
		r.UserID = userid
		r.CardID = cardid
		export, _ := json.Marshal(r)
		fmt.Printf("%s\n", export)
	} else {
		fmt.Printf("File matched and verified\n")
	}
	os.Exit(0)
}

func fioError(id, detail, filename, hash, userid, cardid string, verbose bool) {
	if verbose {
		var r ErrorReturn
		r.Error = id
		r.Detail = detail
		r.Hash = hash
		r.UserID = userid
		r.CardID = cardid
		export, _ := json.Marshal(r)
		fmt.Printf("%s\n", export)
	} else {
		fmt.Printf("Error: %s (detail: %s filename: %s hash: %s)\n", id, detail, filename, hash)
	}
	os.Exit(1)
}

// HexToHash util function
func HexToHash(hex string) (*common.Hash, error) {
	b, err := hexutil.Decode(hex)
	if err != nil {
		msg := fmt.Sprintf("not a hex: %s", hex)
		return nil, NewFBKError(msg, InvalidType)
	}
	hash := common.BytesToHash(b)
	return &hash, nil
}

// GetString getString1 eth func
func GetString(store *contracts.Store, hash common.Hash) (string, error) {
	return store.GetString1(nil, hash)
}

// GetKey get public key
func GetKey(store *contracts.Store, keyUID string) (*EthKey, error) {
	b, err := hexutil.Decode(keyUID)
	if err != nil {
		return nil, NewFBKError("keyuid invalid", InvalidKey)
	}
	hash := common.BytesToHash(b)
	key, err2 := store.GetKey(nil, hash)
	if err2 != nil {
		return nil, NewFBKError("keyuid invalid", InvalidKey)
	}
	var k EthKey
	k = key
	return &k, nil
}

// GetCard get card
func GetCard(store *contracts.Store, cardID string) (*EthCard, error) {
	c, err := HexToHash(cardID)
	if err != nil {
		return nil, NewFBKError("card-id invalid", InvalidCard)
	}
	card, err := store.GetCard(nil, *c)
	if err != nil {
		return nil, NewFBKError("card-id invalid", InvalidCard)
	}
	keyUID := common.Hash(card.Key)
	var res EthCard
	res.Card = card.Card
	res.KeyUID = keyUID.String()
	return &res, nil
}

// CheckCard check a card and returns useruid, keyuid, ktype and error
func CheckCard(card *EthCard, cardID string) (string, string, error) {
	useruid, keyuid, _, err := fcommon.CheckCard(card.Card, cardID)
	return useruid, keyuid, err
}

// Verify check
func Verify(fireblock *contracts.Fireblock, store *contracts.Store, filename, hash, cardID string, verbose bool) {
	hashH, err := HexToHash(hash)
	if err != nil {
		msg := fmt.Sprintf("not a valid hash: %s\n", hash)
		fioError("Invalid hash", msg, filename, hash, "", cardID, verbose)
	}
	cardIDH, err := HexToHash(cardID)
	if err != nil {
		msg := fmt.Sprintf("not a valid card-uid: %s\n", hash)
		fioError("Invalid card-uid", msg, filename, hash, "", cardID, verbose)
	}
	res, err := fireblock.Verify(nil, *hashH, *cardIDH)
	if err != nil {
		fioError("Contract Error", "Is there a fgeth running", "", hash, "", cardID, verbose)
	}
	if !res.Validity {
		fioError("File not registered", "", filename, hash, "", cardID, verbose)
	}
	// get pub key
	keyUID := common.Hash(res.Key)
	kstr := keyUID.String()
	ktype := fcommon.B32Type(kstr)
	key, err := store.GetKey(nil, keyUID)
	if err != nil {
		fioError("File registered but invalid", "Invalid Key", filename, hash, "", cardID, verbose)
	}
	// get Card
	card, err := GetCard(store, cardID)
	if err != nil {
		fioError("File registered but invalid", "Invalid Card", filename, hash, "", cardID, verbose)
	}
	useruidC, keyuidC, err := CheckCard(card, cardID)
	if err != nil {
		fioError("File registered but invalid", "Invalid Card (Not verified)", filename, hash, "", cardID, verbose)
	}
	if keyuidC != kstr {
		fioError("File registered but invalid", "Invalid Card (Hack detected)", filename, hash, "", cardID, verbose)
	}
	pubkey := key.Key
	// get Signature
	signature, err := GetString(store, res.Sig)
	if err != nil {
		fioError("File registered but invalid", "Invalid Signature", filename, hash, useruidC, cardID, verbose)
	}
	message := hash + `||` + cardID
	if ktype == "pgp" {
		r, err := fcommon.PGPVerify(signature, message, pubkey)
		if err != nil || !r {
			fioError("File registered but invalid", "Invalid Signature", filename, hash, useruidC, cardID, verbose)
		}
		fioSuccess(filename, hash, "", cardID, true, verbose)
	} else if ktype == "ecdsa" {
		jwkPubKey, _, err2 := fcommon.ECDSAReadKeys(pubkey)
		if err2 != nil {
			fioError("File registered but ECDSA signature not verified", "", filename, hash, useruidC, cardID, verbose)
		}
		r, err := fcommon.ECDSAVerify(jwkPubKey, message, signature)
		if err != nil || !r {
			fioError("File registered but ECDSA signature not verified", "", filename, hash, useruidC, cardID, verbose)
		}
		fioSuccess(filename, hash, "", cardID, true, verbose)
	} else {
		fioError("File registered but key format not supported", "", filename, hash, useruidC, cardID, verbose)
	}
	// check signature
	fioError("Error unknown", "", filename, hash, useruidC, "", verbose)
}

// VerifyByUser check if there's a proof linked to this useruid
func VerifyByUser(fireblock *contracts.Fireblock, store *contracts.Store, filename, hash, userUID string, verbose bool) {
	hashH, err := HexToHash(hash)
	if err != nil {
		msg := fmt.Sprintf("not a valid hash: %s\n", hash)
		fioError("Invalid hash", msg, filename, hash, userUID, "", verbose)
	}
	userUIDH, err := HexToHash(userUID)
	if err != nil {
		msg := fmt.Sprintf("not a valid user-uid: %s\n", hash)
		fioError("Invalid user-uid", msg, filename, hash, userUID, "", verbose)
	}
	res, err := fireblock.VerifyByUserUID(nil, *hashH, *userUIDH)
	if err != nil {
		fioError("Contract Error", "Is there a fgeth running", "", hash, userUID, "", verbose)
	}
	for i := 0; i < 16; i++ {
		var cardID common.Hash
		cardID = res[4*i]
		if common.EmptyHash(cardID) {
			continue
		}
		// get key
		keyUID := common.Hash(res[4*i+1])
		kstr := keyUID.String()
		ktype := fcommon.B32Type(kstr)
		key, err := store.GetKey(nil, res[4*i+1])
		if err != nil {
			continue
		}
		// get Card
		card, err := GetCard(store, cardID.String())
		if err != nil {
			continue
		}
		useruidC, keyuidC, err := CheckCard(card, cardID.String())
		if useruidC != userUID || kstr != keyuidC || err != nil {
			continue
		}
		// get Signature
		sigSID := res[4*i+2]
		signature, err := GetString(store, sigSID)
		if err != nil {
			continue
		}
		// check signature
		message := hash + `||` + cardID.String()
		if ktype == "pgp" {
			r, err := fcommon.PGPVerify(signature, message, key.Key)
			if err != nil || !r {
				continue
			}
			// check Card
		} else if ktype == "ecdsa" {
			jwkPubKey, _, err2 := fcommon.ECDSAReadKeys(key.Key)
			if err2 != nil {
				continue
			}
			message := hash + `||` + cardID.String()
			r, err := fcommon.ECDSAVerify(jwkPubKey, message, signature)
			if err != nil || !r {
				continue
			}
			// check Card
		} else {
			continue
		}
	}
	fioError("File not registered", "", filename, hash, userUID, "", verbose)
}
