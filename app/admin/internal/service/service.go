package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwt2 "github.com/golang-jwt/jwt/v5"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewRoleService, NewOrganization, NewDictService)

var PwdSecret = ""

func GetUserId(ctx context.Context) float64 {
	var sub float64
	if claims, ok := jwt.FromContext(ctx); ok {
		sub = claims.(jwt2.MapClaims)["user_id"].(float64)
	}
	return sub
}
