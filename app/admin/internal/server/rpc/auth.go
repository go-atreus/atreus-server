package rpc

import (
	context "context"
	"github.com/go-atreus/protocol/admin/auth"
)

var _ auth.AuthServer = (*AuthServer)(nil)

type AuthServer struct{}

func NewAuthServer() *AuthServer {
	return &AuthServer{}
}

func (s *AuthServer) UserToken(ctx context.Context, req *auth.UserTokenReq) (*auth.UserTokenResp, error) {
	return &auth.UserTokenResp{Token: "token-test"}, nil
}

func (s *AuthServer) GetUserToken(ctx context.Context, req *auth.GetUserTokenReq) (*auth.GetUserTokenResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *AuthServer) ForceLogout(ctx context.Context, req *auth.ForceLogoutReq) (*auth.ForceLogoutResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *AuthServer) ParseToken(ctx context.Context, req *auth.ParseTokenReq) (*auth.ParseTokenResp, error) {
	//TODO implement me
	panic("implement me")
}
