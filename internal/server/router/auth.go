package router

import (
	context "context"
	"github.com/go-atreus/protocol/admin"
	"github.com/go-atreus/protocol/admin/auth"
	"github.com/go-atreus/protocol/admin/menu"
	"github.com/go-kratos/kratos/contrib/registry/zookeeper/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-zookeeper/zk"
	"time"
)

type AuthApi struct {
	*auth.AuthImpl
	*menu.MenuImpl
}

var _ auth.AuthHTTPServer = (*AuthApi)(nil)
var _ menu.MenuHTTPServer = (*AuthApi)(nil)

func NewAuthApi(logger log.Logger) *AuthApi {
	address := "discovery:///" + admin.ServerName

	cli, _, err := zk.Connect([]string{"127.0.0.1:12181"}, time.Second*15)
	if err != nil {
		panic(err)
	}
	dis := zookeeper.New(cli)

	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(address),
		grpc.WithDiscovery(dis),
		grpc.WithMiddleware(
			tracing.Client(),
			logging.Client(logger),
		),
	)

	if err != nil {
		log.NewHelper(logger).Error(err)
		return nil
	}
	client := auth.NewAuthClient(conn)
	menuClient := menu.NewMenuClient(conn)
	return &AuthApi{
		AuthImpl: auth.NewAuthImpl(client),
		MenuImpl: menu.NewMenuImpl(menuClient)}
}
