package data

import (
	zookeeper "github.com/go-kratos/kratos/contrib/registry/zookeeper/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-zookeeper/zk"
	"github.com/google/wire"
	"time"
)

var ProviderSet = wire.NewSet(
	NewDiscovery,
	NewRegistrar,
)

func NewDiscovery() registry.Discovery {
	cli, _, err := zk.Connect([]string{"127.0.0.1:12181"}, time.Second*15)
	if err != nil {
		panic(err)
	}
	r := zookeeper.New(cli)
	return r
}

func NewRegistrar() registry.Registrar {
	cli, _, err := zk.Connect([]string{"127.0.0.1:12181"}, time.Second*15)
	if err != nil {
		panic(err)
	}
	r := zookeeper.New(cli)
	return r
}
