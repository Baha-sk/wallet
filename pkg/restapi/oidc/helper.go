/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package oidc

import (
	"errors"
	"fmt"
	"github.com/trustbloc/edge-core/pkg/storage"
)

func openStore(p storage.Provider, name string) (storage.Store, error) {
	err := p.CreateStore(name)
	if err != nil && !errors.Is(err, storage.ErrDuplicateStore) {
		return nil, fmt.Errorf("failed to create store: %w", err)
	}

	return p.OpenStore(name)
}

// validation rules on the received user claims from the OIDC provider go here
func evaluateClaims(u *endUser) error {
	if u.Sub == "" {
		return fmt.Errorf("empty 'sub' in end user claims")
	}

	return nil
}
