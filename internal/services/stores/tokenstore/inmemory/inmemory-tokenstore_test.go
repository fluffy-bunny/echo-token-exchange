package inmemory

import (
	"testing"

	"echo-starter/tests"

	"echo-starter/internal/services/stores/tokenstore"

	di "github.com/dozm/di"
	"github.com/golang/mock/gomock"
)

func TestStore(t *testing.T) {
	tests.RunTest(t, func(ctrl *gomock.Controller) {
		builder, _ := di.NewBuilder(di.App, di.Request, "transient")
		AddSingletonITokenStore(builder)
		ctn := builder.Build()
		tokenstore.RunTestSuite(t, ctn)

	})
}
