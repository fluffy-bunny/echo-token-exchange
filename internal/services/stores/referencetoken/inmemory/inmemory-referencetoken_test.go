package inmemory

import (
	"testing"

	"echo-starter/tests"

	"echo-starter/internal/services/stores/referencetoken"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/golang/mock/gomock"
)

func TestStore(t *testing.T) {
	tests.RunTest(t, func(ctrl *gomock.Controller) {
		builder, _ := di.NewBuilder(di.App, di.Request, "transient")
		AddSingletonIReferenceTokenStore(builder)
		ctn := builder.Build()
		referencetoken.RunTestSuite(t, ctn)

	})
}
