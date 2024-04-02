package service

import (
	context "context"
	"github.com/go-atreus/atreus-server/app/admin/api/user"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwt2 "github.com/golang-jwt/jwt/v5"
)

func NewUserServer(logger log.Logger) *UserServer {
	return &UserServer{}
}

type UserServer struct{}

func (UserServer) GetUserInfo(ctx context.Context, req *user.UserInfoReq) (*user.UserInfoResp, error) {
	var sub float64
	if claims, ok := jwt.FromContext(ctx); ok {
		sub = claims.(jwt2.MapClaims)["user_id"].(float64)
	}
	return &user.UserInfoResp{Userid: int32(sub),
		Name:   "admin",
		Avatar: "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
	}, nil
}
