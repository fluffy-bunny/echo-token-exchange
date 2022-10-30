package jwt

import (
	"context"
	"echo-starter/internal/models"
	"reflect"

	contracts_stores_tokenstore "echo-starter/internal/contracts/stores/tokenstore"

	"google.golang.org/grpc/codes"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/gogo/status"
)

type (
	service struct {
		JwtTokenStore contracts_stores_tokenstore.IJwtTokenStore `inject:""`
	}
)

var reflectType = reflect.TypeOf((*service)(nil))

func assertImplementation() {
	var _ contracts_stores_tokenstore.ITokenStore = (*service)(nil)
}
func init() {
	assertImplementation()
}
func AddSingletonITokenStore(builder *di.Builder) {
	contracts_stores_tokenstore.AddScopedITokenStore(builder, reflectType)

}
func (s *service) StoreToken(ctx context.Context, handle string, info *models.TokenInfo) (string, error) {

	return "", status.Error(codes.Unimplemented, "not implemented")
}
func (s *service) GetToken(ctx context.Context, handle string) (*models.TokenInfo, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}
func (s *service) UpdateToken(ctx context.Context, handle string, info *models.TokenInfo) error {
	return status.Error(codes.Unimplemented, "not implemented")
}
func (s *service) RemoveToken(ctx context.Context, handle string) error {
	return status.Error(codes.Unimplemented, "not implemented")
}
func (s *service) RemoveTokenByClientID(ctx context.Context, clientID string) error {
	return status.Error(codes.Unimplemented, "not implemented")
}
func (s *service) RemoveTokenBySubject(ctx context.Context, subject string) error {
	return status.Error(codes.Unimplemented, "not implemented")
}
func (s *service) RemoveTokenByClientIdAndSubject(ctx context.Context, clientID string, subject string) error {
	return status.Error(codes.Unimplemented, "not implemented")
}
