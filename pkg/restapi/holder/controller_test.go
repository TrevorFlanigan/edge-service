/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package holder

import (
	"fmt"
	"testing"

	cryptomock "github.com/hyperledger/aries-framework-go/pkg/mock/crypto"
	vdrimock "github.com/hyperledger/aries-framework-go/pkg/mock/vdri"
	"github.com/stretchr/testify/require"
	"github.com/trustbloc/edge-core/pkg/storage/memstore"
	"github.com/trustbloc/edge-core/pkg/storage/mockstore"

	"github.com/trustbloc/edge-service/pkg/restapi/holder/operation"
)

func TestController_New(t *testing.T) {
	t.Run("test success", func(t *testing.T) {
		controller, err := New(&operation.Config{StoreProvider: memstore.NewProvider(),
			Crypto: &cryptomock.Crypto{}, VDRI: &vdrimock.MockVDRIRegistry{}})
		require.NoError(t, err)
		require.NotNil(t, controller)
	})

	t.Run("test error", func(t *testing.T) {
		controller, err := New(&operation.Config{StoreProvider: &mockstore.Provider{
			ErrOpenStoreHandle: fmt.Errorf("error open store")},
			VDRI: &vdrimock.MockVDRIRegistry{}})
		require.Error(t, err)
		require.Contains(t, err.Error(), "error open store")
		require.Nil(t, controller)
	})
}

func TestController_GetOperations(t *testing.T) {
	controller, err := New(&operation.Config{StoreProvider: memstore.NewProvider(),
		Crypto: &cryptomock.Crypto{}, VDRI: &vdrimock.MockVDRIRegistry{}})

	require.NoError(t, err)
	require.NotNil(t, controller)

	ops := controller.GetOperations()

	require.Equal(t, 3, len(ops))
}
