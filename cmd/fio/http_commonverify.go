// Copyright 2015-2017 Fireblock.
// This file is part of Fireblock.

// Fireblock is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Fireblock is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Fireblock.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"

	"github.com/fireblock/go-fireblock/fireblocklib"
)

func checkAResult(pkey, key KeyInfo, card *CardInfo, certificate CertificateInfo, hash string) bool {
	certificateHash := hash
	if certificate.Hash != hash {
		certificateHash = certificate.Hash
		batch := certificate.Batch
		hash2 := fireblocklib.Sha256(batch)
		if hash2 != certificate.Hash {
			return false
		}
		if !fireblocklib.IsInBatch(certificate.Batch, hash) {
			return false
		}
	}

	// pkey state
	if (pkey.State & 15) != 3 {
		return false
	}

	// check signature + state of the card
	if card.Txt != "" {
		msg := fmt.Sprintf("register card %s at %d", card.UID, card.Date)
		ck, err := fireblocklib.ECDSAVerify(pkey.Pubkey, msg, card.Signature)
		if err != nil || !ck {
			return false
		}
		// check card
		pstates, err3 := fireblocklib.VerifyCard(card.Txt, pkey.KeyUID, pkey.KType)
		if err3 != nil {
			return false
		}
		// update card
		card.Status = "ok"
		if pstates.Twitter.Status == "ok" {
			card.Twitter.Status = "ok"
			card.Twitter.Proof = pstates.Twitter.Proof
			card.Twitter.UID = pstates.Twitter.UID
		} else {
			card.Twitter.Status = "none"
		}
		if pstates.Github.Status == "ok" {
			card.Github.Status = "ok"
			card.Github.Proof = pstates.Github.Proof
			card.Github.UID = pstates.Github.UID
		} else {
			card.Github.Status = "none"
		}
		if pstates.Linkedin.Status == "ok" {
			card.Linkedin.Status = "ok"
			card.Linkedin.Proof = pstates.Linkedin.Proof
			card.Linkedin.UID = pstates.Linkedin.UID
		} else {
			card.Linkedin.Status = "none"
		}
		if pstates.HTTPS.Status == "ok" {
			card.HTTPS.Status = "ok"
			card.HTTPS.Proof = pstates.HTTPS.Proof
			card.HTTPS.UID = pstates.HTTPS.UID
		} else {
			card.HTTPS.Status = "none"
		}
	}

	// key state
	if (key.State & 7) != 3 {
		return false
	}
	// check certificate
	message := fmt.Sprintf("%s||%s", certificateHash, key.KeyUID)
	ck, err := fireblocklib.VerifySignature(key.KType, key.Pubkey, message, certificate.Signature)
	if err != nil {
		return false
	}
	if !ck {
		return false
	}
	// check delegation
	message2 := fmt.Sprintf("approved key is %s at %d", key.KeyUID, key.Date)
	ck2, err2 := fireblocklib.VerifySignature("ecdsa", pkey.Pubkey, message2, key.Signature)
	if err2 != nil {
		return false
	}
	if !ck2 {
		return false
	}
	// check metadataSignature
	if certificate.MetadataSignature != "" {
		metadataSID := fireblocklib.Keccak256(certificate.Metadata)
		message3 := fmt.Sprintf("%s||%s||%s", metadataSID, certificateHash, key.KeyUID)
		ck3, err3 := fireblocklib.VerifySignature(key.KType, key.Pubkey, message3, certificate.MetadataSignature)
		if err3 != nil {
			return false
		}
		if !ck3 {
			return false
		}
	}
	return true
}
