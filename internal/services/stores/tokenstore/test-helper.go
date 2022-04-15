package tokenstore

import (
	"context"
	"fmt"
	"testing"
	"time"

	contracts_stores_tokenstore "echo-starter/internal/contracts/stores/tokenstore"
	"echo-starter/internal/models"
	"echo-starter/internal/utils"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/go-test/deep"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/require"
)

func RunTestSuite(t *testing.T, ctn di.Container) {
	now := time.Now()
	store := contracts_stores_tokenstore.GetITokenStoreFromContainer(ctn)
	expectedTokenInfo := &models.TokenInfo{
		Metadata: models.TokenMetadata{
			ClientID:   "client-id",
			Subject:    "subject",
			Expiration: now.Add(time.Hour),
		},
		Data: map[string]interface{}{
			"response-key": "response-value",
		},
	}

	expectedTokenInfo2 := &models.TokenInfo{
		Metadata: models.TokenMetadata{
			ClientID:   "client-id2",
			Subject:    "subject2",
			Expiration: now.Add(time.Hour),
		},
		Data: map[string]interface{}{
			"response-key": "response-value",
		},
	}
	// prove our deep comparison works
	require.NotNil(t, deep.Equal(expectedTokenInfo, expectedTokenInfo2))
	handle := utils.GenerateHandle()

	handle, err := store.StoreToken(context.Background(), handle, expectedTokenInfo)
	require.NoError(t, err)
	require.NotEmpty(t, handle)
	actualTokenInfo, err := store.GetToken(context.Background(), handle)
	require.NoError(t, err)
	require.Nil(t, deep.Equal(expectedTokenInfo, actualTokenInfo))

	actualTokenInfo, err = store.GetToken(context.Background(), "garbage")
	require.NoError(t, err)
	require.Nil(t, actualTokenInfo)

	err = store.RemoveToken(context.Background(), handle)
	require.NoError(t, err)
	actualTokenInfo, err = store.GetToken(context.Background(), handle)
	require.NoError(t, err)
	require.Nil(t, actualTokenInfo)

	handles := make([]string, 0)
	for i := 0; i < 10; i++ {
		handle := utils.GenerateHandle()
		handle, err := store.StoreToken(context.Background(), handle, expectedTokenInfo)
		require.NoError(t, err)
		require.NotEmpty(t, handle)
		handles = append(handles, handle)
	}
	for _, handle := range handles {
		actualTokenInfo, err := store.GetToken(context.Background(), handle)
		require.NoError(t, err)
		require.Nil(t, deep.Equal(expectedTokenInfo, actualTokenInfo))
	}
	err = store.RemoveTokenByClientID(context.Background(), expectedTokenInfo.Metadata.ClientID)
	require.NoError(t, err)

	for _, handle := range handles {
		actualTokenInfo, err := store.GetToken(context.Background(), handle)
		require.NoError(t, err)
		require.Nil(t, actualTokenInfo)
	}
	handles = make([]string, 0)
	// diff client_id, same subject
	for i := 0; i < 10; i++ {
		nrt := &models.TokenInfo{}
		copier.Copy(&nrt, expectedTokenInfo)
		nrt.Metadata.ClientID = "client-id-" + fmt.Sprintf("%d", i)
		handle := utils.GenerateHandle()
		handle, err := store.StoreToken(context.Background(), handle, nrt)
		require.NoError(t, err)
		require.NotEmpty(t, handle)
		handles = append(handles, handle)
	}
	for i := 0; i < 10; i++ {
		handle := handles[i]
		nrt := &models.TokenInfo{}
		copier.Copy(&nrt, expectedTokenInfo)
		nrt.Metadata.ClientID = "client-id-" + fmt.Sprintf("%d", i)
		actualTokenInfo, err := store.GetToken(context.Background(), handle)
		require.NoError(t, err)
		require.Nil(t, deep.Equal(nrt, actualTokenInfo))
	}
	err = store.RemoveTokenBySubject(context.Background(), expectedTokenInfo.Metadata.Subject)
	require.NoError(t, err)
	for _, handle := range handles {
		actualTokenInfo, err := store.GetToken(context.Background(), handle)
		require.NoError(t, err)
		require.Nil(t, actualTokenInfo)
	}
	handles = make([]string, 0)

	for i := 0; i < 2; i++ {
		nrt := &models.TokenInfo{}
		copier.Copy(&nrt, expectedTokenInfo)
		nrt.Metadata.ClientID = "client-id-" + fmt.Sprintf("%d", i)
		handle := utils.GenerateHandle()
		handle, err := store.StoreToken(context.Background(), handle, nrt)
		require.NoError(t, err)
		require.NotEmpty(t, handle)
		handles = append(handles, handle)
	}
	err = store.RemoveTokenByClientIdAndSubject(context.Background(), "client-id-0", expectedTokenInfo.Metadata.Subject)
	require.NoError(t, err)
	actualTokenInfo, err = store.GetToken(context.Background(), handles[0])
	require.NoError(t, err)
	require.Nil(t, actualTokenInfo)
	actualTokenInfo, err = store.GetToken(context.Background(), handles[1])
	require.NoError(t, err)
	require.NotNil(t, actualTokenInfo)
	require.Equal(t, "client-id-1", actualTokenInfo.Metadata.ClientID)
}
