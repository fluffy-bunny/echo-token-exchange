package inmemory

import (
	"context"
	"fmt"
	"testing"
	"time"

	"echo-starter/tests"

	contracts_stores_tokenstore "echo-starter/internal/contracts/stores/tokenstore"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/go-test/deep"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/require"
)

func TestStore(t *testing.T) {
	tests.RunTest(t, func(ctrl *gomock.Controller) {
		now := time.Now()
		builder, _ := di.NewBuilder(di.App, di.Request, "transient")
		AddSingletonIReferenceTokenStore(builder)
		ctn := builder.Build()
		store := contracts_stores_tokenstore.GetIReferenceTokenStoreFromContainer(ctn)
		expectedReferenceTokenInfo := &contracts_stores_tokenstore.ReferenceTokenInfo{
			ClientID:   "client-id",
			Subject:    "subject",
			Expiration: now.Add(time.Hour),
			Response: map[string]interface{}{
				"response-key": "response-value",
			},
		}

		expectedReferenceTokenInfo2 := &contracts_stores_tokenstore.ReferenceTokenInfo{
			ClientID:   "client-id2",
			Subject:    "subject2",
			Expiration: now.Add(time.Hour),
			Response: map[string]interface{}{
				"response-key": "response-value",
			},
		}
		// prove our deep comparison works
		require.NotNil(t, deep.Equal(expectedReferenceTokenInfo, expectedReferenceTokenInfo2))

		handle, err := store.StoreReferenceToken(context.Background(), expectedReferenceTokenInfo)
		require.NoError(t, err)
		require.NotEmpty(t, handle)
		actualReferenceTokenInfo, err := store.GetReferenceToken(context.Background(), handle)
		require.NoError(t, err)
		require.Nil(t, deep.Equal(expectedReferenceTokenInfo, actualReferenceTokenInfo))

		actualReferenceTokenInfo, err = store.GetReferenceToken(context.Background(), "garbage")
		require.Error(t, err)
		require.Nil(t, actualReferenceTokenInfo)

		err = store.RemoveReferenceToken(context.Background(), handle)
		require.NoError(t, err)
		actualReferenceTokenInfo, err = store.GetReferenceToken(context.Background(), "garbage")
		require.Error(t, err)
		require.Nil(t, actualReferenceTokenInfo)

		handles := make([]string, 0)
		for i := 0; i < 10; i++ {
			handle, err := store.StoreReferenceToken(context.Background(), expectedReferenceTokenInfo)
			require.NoError(t, err)
			require.NotEmpty(t, handle)
			handles = append(handles, handle)
		}
		for _, handle := range handles {
			actualReferenceTokenInfo, err := store.GetReferenceToken(context.Background(), handle)
			require.NoError(t, err)
			require.Nil(t, deep.Equal(expectedReferenceTokenInfo, actualReferenceTokenInfo))
		}
		err = store.RemoveReferenceTokenByClientID(context.Background(), expectedReferenceTokenInfo.ClientID)
		require.NoError(t, err)

		for _, handle := range handles {
			actualReferenceTokenInfo, err := store.GetReferenceToken(context.Background(), handle)
			require.Error(t, err)
			require.Nil(t, actualReferenceTokenInfo)
		}
		handles = make([]string, 0)
		// diff client_id, same subject
		for i := 0; i < 10; i++ {
			nrt := &contracts_stores_tokenstore.ReferenceTokenInfo{}
			copier.Copy(&nrt, expectedReferenceTokenInfo)
			nrt.ClientID = "client-id-" + fmt.Sprintf("%d", i)

			handle, err := store.StoreReferenceToken(context.Background(), nrt)
			require.NoError(t, err)
			require.NotEmpty(t, handle)
			handles = append(handles, handle)
		}
		for i := 0; i < 10; i++ {
			handle := handles[i]
			nrt := &contracts_stores_tokenstore.ReferenceTokenInfo{}
			copier.Copy(&nrt, expectedReferenceTokenInfo)
			nrt.ClientID = "client-id-" + fmt.Sprintf("%d", i)
			actualReferenceTokenInfo, err := store.GetReferenceToken(context.Background(), handle)
			require.NoError(t, err)
			require.Nil(t, deep.Equal(nrt, actualReferenceTokenInfo))
		}
		err = store.RemoveReferenceTokenBySubject(context.Background(), expectedReferenceTokenInfo.Subject)
		require.NoError(t, err)
		for _, handle := range handles {
			actualReferenceTokenInfo, err := store.GetReferenceToken(context.Background(), handle)
			require.Error(t, err)
			require.Nil(t, actualReferenceTokenInfo)
		}
		handles = make([]string, 0)

		for i := 0; i < 2; i++ {
			nrt := &contracts_stores_tokenstore.ReferenceTokenInfo{}
			copier.Copy(&nrt, expectedReferenceTokenInfo)
			nrt.ClientID = "client-id-" + fmt.Sprintf("%d", i)

			handle, err := store.StoreReferenceToken(context.Background(), nrt)
			require.NoError(t, err)
			require.NotEmpty(t, handle)
			handles = append(handles, handle)
		}
		err = store.RemoveReferenceTokenByClientIdAndSubject(context.Background(), "client-id-0", expectedReferenceTokenInfo.Subject)
		require.NoError(t, err)
		actualReferenceTokenInfo, err = store.GetReferenceToken(context.Background(), handles[0])
		require.Error(t, err)
		require.Nil(t, actualReferenceTokenInfo)
		actualReferenceTokenInfo, err = store.GetReferenceToken(context.Background(), handles[1])
		require.NoError(t, err)
		require.NotNil(t, actualReferenceTokenInfo)
		require.Equal(t, "client-id-1", actualReferenceTokenInfo.ClientID)

	})
}
