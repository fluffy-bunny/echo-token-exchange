package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	startup "echo-starter/internal/startup"
	tests "echo-starter/tests"
	httptest "net/http/httptest"

	di "github.com/dozm/di"
	echo_contracts_startup "github.com/fluffy-bunny/fluffycore/echo/contracts/startup"
	runtime "github.com/fluffy-bunny/fluffycore/echo/runtime"
	gomock "github.com/golang/mock/gomock"
	echo "github.com/labstack/echo/v4"
	require "github.com/stretchr/testify/require"
)

func TestHealthCheck(t *testing.T) {
	tests.RunTest(t, func(ctrl *gomock.Controller) {

		folderChanger := NewFolderChanger("../../cmd/server")
		defer folderChanger.ChangeBack()

		startChan := make(chan bool)

		startup := startup.NewStartup()
		var myEcho *echo.Echo
		hooks := &echo_contracts_startup.Hooks{
			PrebuildHook: func(builder di.ContainerBuilder) error {
				// register a null task engine

				return nil
			},
			PreStartHook: func(echo *echo.Echo) error {
				myEcho = echo
				startChan <- true
				return nil
			},
		}
		startup.AddHooks(hooks)

		r := runtime.New(startup)
		future := tests.ExecuteWithPromiseAsync(r)

		<-startChan

		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		w := httptest.NewRecorder()
		myEcho.ServeHTTP(w, req)

		res := w.Result()
		defer res.Body.Close()
		data, err := ioutil.ReadAll(res.Body)
		require.NoError(t, err)
		fmt.Println("data:", string(data))

		r.Stop()
		future.Join()
	})
}
