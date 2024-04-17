package data

import (
	"context"
	"database/sql"
	"github.com/go-atreus/atreus-server/app/admin/api/dict"
	"github.com/go-atreus/atreus-server/app/admin/api/menu"
	"github.com/go-atreus/atreus-server/app/admin/api/organization"
	"github.com/go-atreus/atreus-server/app/admin/api/role"
	"github.com/go-atreus/atreus-server/app/admin/api/user"
	"github.com/go-atreus/atreus-server/app/admin/internal/conf"
	zookeeper "github.com/go-kratos/kratos/contrib/registry/zookeeper/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-redis/redis/v8"
	"github.com/go-zookeeper/zk"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
	// database driver
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var ProviderSet = wire.NewSet(
	NewDiscovery,
	NewRegistrar,
	NewData,
	NewUserRepo,
	NewOrm,
	NewRedisCmd,
)

type Data struct {
	log      *log.Helper
	ORM      *gorm.DB
	redisCli redis.Cmdable
}

// NewData .
func NewData(orm *gorm.DB, redisCmd redis.Cmdable, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "user-service/data"))

	d := &Data{
		ORM:      orm,
		redisCli: redisCmd,
		log:      log,
	}
	return d, func() {
	}, nil
}

func NewOrm(logger log.Logger, bootstrap *conf.Bootstrap) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "data"))
	dataConf := bootstrap.Data

	sqlDB, err := sql.Open(dataConf.Database.Driver, dataConf.Database.Source)
	if err != nil {
		log.Fatalf("sql dsn(%v) error: %v", dataConf.Database.Source, err)
	}
	sqlDB.SetMaxIdleConns(int(dataConf.Database.Idle))
	sqlDB.SetMaxOpenConns(int(dataConf.Database.Active))
	idleTimeout, _ := time.ParseDuration(dataConf.Database.IdleTimeout)
	sqlDB.SetConnMaxLifetime(idleTimeout)

	orm, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger:                                   logger2.Default.LogMode(logger2.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("db dsn(%v) error: %v", dataConf.Database.Source, err)
	}

	// migrate
	err = orm.AutoMigrate(
		&user.SysUserORM{},
		&menu.SysMenuORM{},
		&dict.SysDictORM{},
		&organization.SysOrganizationORM{},
		&role.SysRoleORM{})
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return orm
}

func NewRedisCmd(conf *conf.Data, logger log.Logger) redis.Cmdable {
	log := log.NewHelper(log.With(logger, "module", "data/redis"))
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     "openIM123",
		Username:     "",
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  time.Second * 2,
		PoolSize:     10,
	})
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()
	err := client.Ping(timeout).Err()
	if err != nil {
		log.Fatalf("redis connect error: %v", err)
	}
	return client
}

func NewDiscovery() registry.Discovery {
	cli, _, err := zk.Connect([]string{"127.0.0.1:12181"}, time.Second*15)
	if err != nil {
		return nil
	}
	if err != nil {
		panic(err)
	}
	r := zookeeper.New(cli)
	return r
}

func NewRegistrar() registry.Registrar {
	cli, _, err := zk.Connect([]string{"127.0.0.1:12181"}, time.Second*15)
	if err != nil {
		return nil
	}
	if err != nil {
		panic(err)
	}
	r := zookeeper.New(cli)
	return r
}
