// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/trustbloc/edge-agent/cmd/wallet-server

go 1.16

require (
	github.com/cenkalti/backoff/v4 v4.1.2
	github.com/coreos/go-oidc v2.2.1+incompatible
	github.com/duo-labs/webauthn v0.0.0-20200714211715-1daaee874e43
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.8.0
	github.com/hyperledger/aries-framework-go v0.1.8-0.20220113090443-b751c0ae7b22
	github.com/hyperledger/aries-framework-go-ext/component/storage/couchdb v0.0.0-20210909220549-ce3a2ee13e22
	github.com/hyperledger/aries-framework-go-ext/component/storage/mongodb v0.0.0-20211219215001-23cd75276fdc
	github.com/hyperledger/aries-framework-go-ext/component/storage/mysql v0.0.0-20210909220549-ce3a2ee13e22
	github.com/hyperledger/aries-framework-go-ext/component/vdr/orb v0.1.4-0.20211229160742-30796de12ba6
	github.com/hyperledger/aries-framework-go/component/storage/leveldb v0.0.0-20210916154931-0196c3a2d102
	github.com/hyperledger/aries-framework-go/component/storageutil v0.0.0-20210916154931-0196c3a2d102
	github.com/hyperledger/aries-framework-go/spi v0.0.0-20211206182816-9cdcbcd09dc2
	github.com/piprate/json-gold v0.4.1-0.20210813112359-33b90c4ca86c
	github.com/rs/cors v1.8.2
	github.com/spf13/cobra v1.3.0
	github.com/stretchr/testify v1.7.0
	github.com/trustbloc/edge-agent v0.0.0-00010101000000-000000000000
	github.com/trustbloc/edge-core v0.1.8-0.20220113141450-e19ffd091d98
)

replace github.com/trustbloc/edge-agent => ../..
