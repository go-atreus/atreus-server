package service

import (
	context "context"
	"errors"
	"github.com/go-atreus/atreus-server/app/admin/api/auth"
	"github.com/go-atreus/atreus-server/app/admin/internal/biz"
	"github.com/go-atreus/atreus-server/app/admin/internal/conf"
	"github.com/go-atreus/tools/utils"
	"github.com/golang-jwt/jwt/v5"
)

var _ auth.AuthServer = (*AuthServer)(nil)

type AuthServer struct {
	config   *conf.Auth
	userRepo biz.UserRepo
}

func (s *AuthServer) UserLogin(ctx context.Context, req *auth.UserLoginReq) (res *auth.UserTokenResp, err error) {
	if req.Type == "mobile" {
		return nil, errors.New("订购功能，QQ：420978249")
	}
	if req.Username == "" || req.Password == "" {
		return nil, errors.New("username or password is empty")
	}
	findUser, err := s.userRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	req.Password = utils.Md5(req.Password, s.config.PwdSecret)
	if req.Password != findUser.Password {
		return nil, errors.New("密码错误")
	}
	// generate token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": findUser.Id,
	})
	signedString, err := claims.SignedString([]byte(s.config.Key))
	if err != nil {
		return nil, err
	}
	return &auth.UserTokenResp{Token: signedString}, nil
}

func NewAuthServer(config *conf.Auth, userRepo biz.UserRepo) *AuthServer {
	PwdSecret = config.PwdSecret
	return &AuthServer{config: config, userRepo: userRepo}
}

func (s *AuthServer) UserToken(ctx context.Context, req *auth.UserTokenReq) (*auth.UserTokenResp, error) {
	return &auth.UserTokenResp{Token: "token-test"}, nil
}

func (s *AuthServer) GetUserToken(ctx context.Context, req *auth.GetUserTokenReq) (*auth.GetUserTokenResp, error) {
	// generate token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": 1,
	})
	signedString, err := claims.SignedString([]byte(s.config.Key))
	if err != nil {
		return nil, err
	}
	return &auth.GetUserTokenResp{Token: signedString}, nil
}

func (s *AuthServer) ForceLogout(ctx context.Context, req *auth.ForceLogoutReq) (*auth.ForceLogoutResp, error) {
	return &auth.ForceLogoutResp{}, nil
}

func (s *AuthServer) ParseToken(ctx context.Context, req *auth.ParseTokenReq) (*auth.ParseTokenResp, error) {
	//TODO implement me
	panic("implement me")
}
