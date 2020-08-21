/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package sdscomm

// Error messages
const (
	failureEnsuringDIDVaultExistsErrMsg   = "failure while ensuring that the user's DID vault exists: %w"
	failureEnsuringCredVaultExistsErrMsg  = "failure while ensuring that the user's credential vault exists: %w"
	unexpectedErrorOnCreateVaultCall      = `unexpected error during the "create vault" call to SDS: %w`
	failureStoringDIDDocErrMsg            = "failure while storing DID document: %w"
	failureStoringCredErrMsg              = "failure while storing credential: %w"
	failureGeneratingEncryptedDocIDErrMsg = "failure while generating encrypted document ID: %w"
	failureMarshallingStructuredDocErrMsg = "failure while marshalling structured document: %w"
	failureStoringDocErrMsg               = "failed to store document: %w"
)

// Log messages
const (
	vaultAlreadyExistsLogMsg = "%s vault already exists. Skipping vault creation."
	newVaultCreatedLogMsg    = "%s vault created."
)
