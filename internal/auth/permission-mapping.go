package auth

import (
	"echo-starter/internal/wellknown"

	middleware_oidc "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/oidc"
	services_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/claimsprincipal"
)

// BuildGrpcEntrypointPermissionsClaimsMap ...
func BuildGrpcEntrypointPermissionsClaimsMap() map[string]*middleware_oidc.EntryPointConfig {
	entryPointClaimsBuilder := services_claimsprincipal.NewEntryPointClaimsBuilder()
	// HEALTH SERVICE START
	//---------------------------------------------------------------------------------------------------
	// health check is open to anyone
	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimsMapOpen(wellknown.HealthzPath)
	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimsMapOpen(wellknown.WellKnownOpenIDCOnfiguationPath)
	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimsMapOpen(wellknown.WellKnownJWKS)
	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimsMapOpen(wellknown.OAuth2TokenPath)
	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimsMapOpen(wellknown.OAuth2RevokePath)

	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimsMapOpen(wellknown.ReadyPath)
	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimsMapOpen(wellknown.LoginPath)
	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimsMapOpen(wellknown.LogoutPath)
	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimsMapOpen(wellknown.OIDCCallbackPath)
	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimsMapOpen(wellknown.OAuth2CallbackPath)
	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimsMapOpen(wellknown.HomePath)
	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimsMapOpen(wellknown.AboutPath)
	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimsMapOpen(wellknown.ErrorPath)
	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimsMapOpen(wellknown.UnauthorizedPath)
	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimsMapOpen("/static*")

	cMap := entryPointClaimsBuilder.GrpcEntrypointClaimsMap
	return cMap
}
