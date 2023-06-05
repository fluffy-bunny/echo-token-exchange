package templates

import (
	"echo-starter/internal/models"

	fluffycore_contracts_common "github.com/fluffy-bunny/fluffycore/contracts/common"
	core_echo_templates "github.com/fluffy-bunny/fluffycore/echo/templates"
	core_wellknown "github.com/fluffy-bunny/fluffycore/echo/wellknown"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, claimsPrincipal fluffycore_contracts_common.IClaimsPrincipal, code int, name string, data map[string]interface{}) error {
	data["isAuthenticated"] = func() bool {
		return claimsPrincipal.HasClaimType(core_wellknown.ClaimTypeAuthenticated)
	}
	data["paths"] = models.NewPaths()
	data["claims"] = claimsPrincipal.GetClaims()
	return core_echo_templates.Render(c, code, name, data)

}
