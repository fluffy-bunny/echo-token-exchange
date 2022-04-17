package tokenstore

import (
	"context"
	"encoding/json"
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
	expectedTokenInfoOri := &models.TokenInfo{
		Metadata: models.TokenMetadata{
			Type:       models.TokenTypeReferenceToken,
			ClientID:   "0000",
			Subject:    "1111",
			Expiration: now.Add(time.Hour),
		},
		Data: map[string]interface{}{
			"aud": []string{
				"b2b-client",
				"users",
				"invoices",
			},
			"client_id": "b2b-client",
			"exp":       1650123272,
			"iat":       1650119672,
			"iss":       "http://localhost:1323",
			"jti":       "c9dd7u1ld5lnsc78u3f0",
			"scope": []string{
				"offline_access",
				"a",
				"b",
				"c",
				"users.read",
				"invoices",
			},
		},
	}
	// marshaling to map to do deep comparison
	expectedTokenInfoM := map[string]interface{}{}
	jM, _ := json.Marshal(expectedTokenInfoOri)
	json.Unmarshal(jM, &expectedTokenInfoM)

	expectedTokenInfo2Ori := &models.TokenInfo{
		Metadata: models.TokenMetadata{
			ClientID:   "client-id2",
			Subject:    "subject2",
			Expiration: now.Add(time.Hour),
		},
		Data: map[string]interface{}{
			"response-key": "response-value",
		},
	}
	expectedTokenInfo2 := map[string]interface{}{}
	jM2, _ := json.Marshal(expectedTokenInfo2Ori)
	json.Unmarshal(jM2, &expectedTokenInfo2)

	// prove our deep comparison works
	require.NotNil(t, deep.Equal(expectedTokenInfoM, expectedTokenInfo2))
	handle := utils.GenerateHandle()

	handle, err := store.StoreToken(context.Background(), handle, expectedTokenInfoOri)
	require.NoError(t, err)
	require.NotEmpty(t, handle)
	actualTokenInfo, err := store.GetToken(context.Background(), handle)
	require.NoError(t, err)
	actualTokenInfoM := map[string]interface{}{}
	jM, _ = json.Marshal(actualTokenInfo)
	json.Unmarshal(jM, &actualTokenInfoM)

	require.Nil(t, deep.Equal(expectedTokenInfoM, actualTokenInfoM))

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
		handle, err := store.StoreToken(context.Background(), handle, expectedTokenInfoOri)
		require.NoError(t, err)
		require.NotEmpty(t, handle)
		handles = append(handles, handle)
		defer store.RemoveToken(context.Background(), handle)
	}
	for _, handle := range handles {
		actualTokenInfo, err := store.GetToken(context.Background(), handle)
		require.NoError(t, err)
		actualTokenInfoM := map[string]interface{}{}
		jM, _ = json.Marshal(actualTokenInfo)
		json.Unmarshal(jM, &actualTokenInfoM)

		require.Nil(t, deep.Equal(expectedTokenInfoM, actualTokenInfoM))
	}
	err = store.RemoveTokenByClientID(context.Background(), expectedTokenInfoOri.Metadata.ClientID)
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
		copier.Copy(&nrt, expectedTokenInfoM)
		nrt.Metadata.ClientID = "cli" + fmt.Sprintf("%d", i)
		handle := utils.GenerateHandle()
		handle, err := store.StoreToken(context.Background(), handle, nrt)
		require.NoError(t, err)
		require.NotEmpty(t, handle)
		handles = append(handles, handle)
	}
	for i := 0; i < 10; i++ {
		handle := handles[i]
		nrt := &models.TokenInfo{}
		copier.Copy(&nrt, expectedTokenInfoM)
		nrt.Metadata.ClientID = "cli" + fmt.Sprintf("%d", i)

		nrtM := map[string]interface{}{}
		jM, _ := json.Marshal(nrt)
		json.Unmarshal(jM, &nrtM)

		actualTokenInfo, err := store.GetToken(context.Background(), handle)
		require.NoError(t, err)
		actualTokenInfoM := map[string]interface{}{}
		jM, _ = json.Marshal(actualTokenInfo)
		json.Unmarshal(jM, &actualTokenInfoM)

		require.Nil(t, deep.Equal(nrtM, actualTokenInfoM))

	}
	err = store.RemoveTokenBySubject(context.Background(), expectedTokenInfoOri.Metadata.Subject)
	require.NoError(t, err)
	for _, handle := range handles {
		actualTokenInfo, err := store.GetToken(context.Background(), handle)
		require.NoError(t, err)
		require.Nil(t, actualTokenInfo)
	}
	handles = make([]string, 0)

	for i := 0; i < 2; i++ {
		nrt := &models.TokenInfo{}
		copier.Copy(&nrt, expectedTokenInfoM)
		nrt.Metadata.ClientID = "cli" + fmt.Sprintf("%d", i)
		handle := utils.GenerateHandle()
		handle, err := store.StoreToken(context.Background(), handle, nrt)
		require.NoError(t, err)
		require.NotEmpty(t, handle)
		handles = append(handles, handle)
	}
	err = store.RemoveTokenByClientIdAndSubject(context.Background(), "cli0", expectedTokenInfoOri.Metadata.Subject)
	require.NoError(t, err)
	actualTokenInfo, err = store.GetToken(context.Background(), handles[0])
	require.NoError(t, err)
	require.Nil(t, actualTokenInfo)
	actualTokenInfo, err = store.GetToken(context.Background(), handles[1])
	require.NoError(t, err)
	require.NotNil(t, actualTokenInfo)
	require.Equal(t, "cli1", actualTokenInfo.Metadata.ClientID)
}
