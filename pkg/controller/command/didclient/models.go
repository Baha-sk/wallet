/*
Copyright SecureKey Technologies Inc. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package didclient

// CreateDIDRequest model
//
// This is used for creating DID
//
type CreateDIDRequest struct {
	PublicKeys []PublicKey `json:"publicKeys,omitempty"`
}

// PublicKey public key
type PublicKey struct {
	ID       string   `json:"id,omitempty"`
	Type     string   `json:"type,omitempty"`
	Encoding string   `json:"encoding,omitempty"`
	KeyType  string   `json:"keyType,omitempty"`
	Usage    []string `json:"usage,omitempty"`
	Recovery bool     `json:"recovery,omitempty"`
	Value    string   `json:"value,omitempty"`
}

type CreateDIDResponse struct {
	DID map[string]interface{}
	PrivateKey string
}
