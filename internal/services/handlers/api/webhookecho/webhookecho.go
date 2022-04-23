package webhookecho

import (
	"echo-starter/internal/wellknown"
	"io/ioutil"
	"net/http"
	"reflect"

	contracts_handler "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/handler"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/labstack/echo/v4"
)

type (
	service struct{}
)

func assertImplementation() {
	var _ contracts_handler.IHandler = (*service)(nil)
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddScopedIHandler registers the *service as a singleton.
func AddScopedIHandler(builder *di.Builder) {
	contracts_handler.AddScopedIHandlerEx(builder,
		reflectType,
		[]contracts_handler.HTTPVERB{
			contracts_handler.GET,
			contracts_handler.POST,
		},
		wellknown.WebhookEchoPath)
}

func (s *service) GetMiddleware() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{}
}

func (s *service) Do(c echo.Context) error {
	switch c.Request().Method {
	case http.MethodGet:
		return s.post(c)
	case http.MethodPost:
		return s.post(c)
	default:
		return echo.NewHTTPError(http.StatusMethodNotAllowed)
	}

}

type echoResponse struct {
	Path   string      `json:"path"`
	Method string      `json:"method"`
	Header http.Header `json:"header"`
	Body   interface{} `json:"body"`
}

func (s *service) post(c echo.Context) error {
	request := c.Request()
	response := &echoResponse{
		Path:   request.URL.Path,
		Header: request.Header,
		Method: request.Method,
	}

	b, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}
	response.Body = string(b)

	return c.JSON(http.StatusOK, response)

}
