/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package msp

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/hyperledger-labs/fabric-smart-client/platform/common/utils"
	config2 "github.com/hyperledger-labs/fabric-smart-client/platform/fabric/core/generic/config"
	"github.com/hyperledger-labs/fabric-smart-client/platform/fabric/core/generic/msp/driver/mock"
	fdriver "github.com/hyperledger-labs/fabric-smart-client/platform/fabric/driver"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/services/sig"
	mem "github.com/hyperledger-labs/fabric-smart-client/platform/view/services/storage/driver/memory"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/services/storage/kvs"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/view"
)

func TestServiceExtended(t *testing.T) {
	t.Parallel()
	cp := &mock.ConfigProvider{}
	cp.IsSetReturns(false)
	cp.GetStringReturns("default_msp")

	kvss, err := kvs.New(utils.MustGet(mem.NewDriver().NewKVS("")), "", kvs.DefaultCacheSize)
	require.NoError(t, err)

	des := sig.NewMultiplexDeserializer()

	config, err := config2.NewService(cp, "default", true)
	require.NoError(t, err)

	signerService := &mock.SignerService{}
	binderService := &mock.BinderService{}
	defaultViewIdentity := view.Identity("default_view_identity")

	mspService := NewLocalMSPManager(config, kvss, signerService, binderService, defaultViewIdentity, des, 100)
	mspService.defaultMSP = "default_msp"

	// Test DefaultMSP
	require.Equal(t, "default_msp", mspService.DefaultMSP())

	// Test SetDefaultIdentity and DefaultIdentity
	id := view.Identity("id1")
	sid := &mock.SigningIdentity{}
	mspService.SetDefaultIdentity("default_msp", id, sid)
	require.Equal(t, id, mspService.DefaultIdentity())
	require.Equal(t, sid, mspService.DefaultSigningIdentity())

	// Test SignerService()
	require.Equal(t, signerService, mspService.SignerService())

	// Test CacheSize()
	require.Equal(t, 100, mspService.CacheSize())

	// Test Config()
	require.Equal(t, config, mspService.Config())

	// Test IsMe
	signerService.IsMeReturns(true)
	require.True(t, mspService.IsMe(context.Background(), id))
	signerService.IsMeReturns(false)
	require.False(t, mspService.IsMe(context.Background(), id))

	// Test AddMSP and GetIdentityByID
	require.NoError(t, mspService.AddMSP("apple", BccspMSP, "enrollment1", func(opts *fdriver.IdentityOptions) (view.Identity, []byte, error) {
		return id, nil, nil
	}))

	resId, err := mspService.GetIdentityByID("apple")
	require.NoError(t, err)
	require.Equal(t, id, resId)

	resId, err = mspService.GetIdentityByID("enrollment1")
	require.NoError(t, err)
	require.Equal(t, id, resId)

	// Test GetIdentityByID via BinderService
	binderId := view.Identity("binder_id")
	binderService.GetIdentityReturns(binderId, nil)
	resId, err = mspService.GetIdentityByID("non_existent")
	require.NoError(t, err)
	require.Equal(t, binderId, resId)

	// Test Identity
	resId, err = mspService.Identity("apple")
	require.NoError(t, err)
	require.Equal(t, id, resId)

	// Test GetIdentityInfoByIdentity for BCCSP
	info := mspService.GetIdentityInfoByIdentity(BccspMSP, id)
	require.NotNil(t, info)
	require.Equal(t, "apple", info.ID)

	// Test GetIdentityInfoByIdentity for non-existent BCCSP
	require.Nil(t, mspService.GetIdentityInfoByIdentity(BccspMSP, view.Identity("unknown")))

	// Test GetIdentityInfoByIdentity for Idemix (requires scan)
	idemixId := view.Identity("idemix_id")
	require.NoError(t, mspService.AddMSP("orange", IdemixMSP, "enrollment2", func(opts *fdriver.IdentityOptions) (view.Identity, []byte, error) {
		return idemixId, nil, nil
	}))

	info = mspService.GetIdentityInfoByIdentity(IdemixMSP, idemixId)
	require.NotNil(t, info)
	require.Equal(t, "orange", info.ID)

	// Test AnonymousIdentity
	// This requires an "idemix" MSP to be registered
	require.NoError(t, mspService.AddMSP("idemix", IdemixMSP, "idemix_enrollment", func(opts *fdriver.IdentityOptions) (view.Identity, []byte, error) {
		return idemixId, nil, nil
	}))
	anonId, err := mspService.AnonymousIdentity()
	require.NoError(t, err)
	require.Equal(t, idemixId, anonId)
	require.Equal(t, 2, binderService.BindCallCount())
}
