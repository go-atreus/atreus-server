package service

import (
	context "context"
	"github.com/go-atreus/atreus-server/app/admin/api/user"
	"github.com/go-atreus/atreus-server/app/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwt2 "github.com/golang-jwt/jwt/v5"
)

var _ user.UserServer = (*UserService)(nil)

func NewUserServer(logger log.Logger, data *data.Data) *UserService {
	return &UserService{
		UserDefaultServer: user.UserDefaultServer{DB: data.ORM},
	}
}

type UserService struct {
	user.UserDefaultServer
}

func (UserService) GetUserInfo(ctx context.Context, req *user.UserInfoReq) (*user.SysUser, error) {
	var sub float64
	if claims, ok := jwt.FromContext(ctx); ok {
		sub = claims.(jwt2.MapClaims)["user_id"].(float64)
	}
	return &user.SysUser{Id: int32(sub),
		NickName: "admin",
		Avatar:   "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
	}, nil
}
