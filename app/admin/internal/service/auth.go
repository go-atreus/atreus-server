package service

import (
	context "context"
	"errors"
	"github.com/go-atreus/atreus-server/app/admin/api/auth"
	"github.com/go-atreus/atreus-server/app/admin/api/user"
	"github.com/go-atreus/atreus-server/app/admin/internal/conf"
	"github.com/go-atreus/atreus-server/app/admin/internal/data"
	"github.com/go-atreus/tools/utils"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var _ auth.AuthServer = (*AuthServer)(nil)

type AuthServer struct {
	config *conf.Auth
	data   *data.Data
}

func (s *AuthServer) UserLogin(ctx context.Context, req *auth.UserLoginReq) (res *auth.UserTokenResp, err error) {
	if req.Type == "mobile" {
		return nil, errors.New("订购功能，QQ：420978249")
	}
	if req.Username == "" || req.Password == "" {
		return nil, errors.New("username or password is empty")
	}
	req.Password = utils.Md5(req.Password, s.config.PwdSecret)
	var findUser user.SysUserORM
	err = s.data.ORM.Where("username = ? AND password = ?", req.Username, req.Password).First(&findUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户名或密码错误")
		}
		return nil, err
	}
	// generate token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": 1,
	})
	signedString, err := claims.SignedString([]byte(s.config.Key))
	if err != nil {
		return nil, err
	}
	return &auth.UserTokenResp{Token: signedString}, nil
}

func NewAuthServer(config *conf.Auth, data *data.Data) *AuthServer {
	return &AuthServer{config: config, data: data}
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
