package service

import (
	context "context"
	"github.com/go-atreus/atreus-server/app/admin/api/auth"
	"github.com/go-atreus/atreus-server/app/admin/internal/conf"
	"github.com/golang-jwt/jwt/v5"
)

var _ auth.AuthServer = (*AuthServer)(nil)

type AuthServer struct {
	key string
}

func (s *AuthServer) UserLogin(ctx context.Context, req *auth.UserLoginReq) (*auth.UserTokenResp, error) {
	// generate token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": 1,
	})
	signedString, err := claims.SignedString([]byte(s.key))
	if err != nil {
		return nil, err
	}
	return &auth.UserTokenResp{Token: signedString}, nil
}

func NewAuthServer(config *conf.Auth) *AuthServer {
	return &AuthServer{key: config.Key}
}

func (s *AuthServer) UserToken(ctx context.Context, req *auth.UserTokenReq) (*auth.UserTokenResp, error) {
	return &auth.UserTokenResp{Token: "token-test"}, nil
}

func (s *AuthServer) GetUserToken(ctx context.Context, req *auth.GetUserTokenReq) (*auth.GetUserTokenResp, error) {
	// generate token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": 1,
	})
	signedString, err := claims.SignedString([]byte(s.key))
	if err != nil {
		return nil, err
	}
	return &auth.GetUserTokenResp{Token: signedString}, nil
}

func (s *AuthServer) ForceLogout(ctx context.Context, req *auth.ForceLogoutReq) (*auth.ForceLogoutResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *AuthServer) ParseToken(ctx context.Context, req *auth.ParseTokenReq) (*auth.ParseTokenResp, error) {
	//TODO implement me
	panic("implement me")
}
