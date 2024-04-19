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

func (u UserService) GetUserInfo(ctx context.Context, req *user.UserInfoReq) (res *user.SysUser, err error) {
	var sub float64
	if claims, ok := jwt.FromContext(ctx); ok {
		sub = claims.(jwt2.MapClaims)["user_id"].(float64)
	}
	var findUser *user.SysUserORM
	if err = u.DB.Where("id = ?", sub).Preload("Role").Preload("Role.Menu", "type = ?", 2).First(&findUser).Error; err != nil {
		return
	}
	var roles []string
	var permissions []string
	for _, role := range findUser.Role {
		roles = append(roles, role.Code)
		for _, menu := range role.Menu {
			permissions = append(permissions, menu.Permission)
		}
	}
	return &user.SysUser{
		Id:          int32(sub),
		NickName:    "admin",
		Avatar:      "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
		Permissions: permissions,
		Roles:       roles,
	}, nil
}
