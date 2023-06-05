package tests

import (
	"testing"

	fluffycore_async "github.com/fluffy-bunny/fluffycore/async"
	runtime "github.com/fluffy-bunny/fluffycore/echo/runtime"

	"github.com/golang/mock/gomock"
	"github.com/reugn/async"
)

const bufSize = 1024 * 1024

func CreateForEach(setUp func(*testing.T), tearDown func()) func(*testing.T, func(*gomock.Controller)) {
	return func(t *testing.T, testFunc func(ctrl *gomock.Controller)) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		setUp(t)
		testFunc(ctrl)
		tearDown()
	}
}

var RunTest = CreateForEach(setUp, tearDown)

func setUp(t *testing.T) {
	// SETUP METHOD WHICH IS REQUIRED TO RUN FOR EACH TEST METHOD
	// your code here
	t.Setenv("APPLICATION_ENVIRONMENT", "Test")
}

func tearDown() {
	// TEAR DOWN METHOD WHICH IS REQUIRED TO RUN FOR EACH TEST METHOD
	// your code here
}

func ExecuteWithPromiseAsync(runtime *runtime.Runtime) async.Future[interface{}] {
	future := fluffycore_async.ExecuteWithPromiseAsync(func(promise async.Promise[interface{}]) {
		var err error
		defer func() {
			promise.Success(&fluffycore_async.AsyncResponse{
				Message: "End Serve - echo Server",
				Error:   err,
			})
		}()
		runtime.Run()
	})
	return future
}
