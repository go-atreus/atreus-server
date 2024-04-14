package data

import (
	"database/sql"
	"github.com/go-atreus/atreus-server/app/admin/api/dict"
	"github.com/go-atreus/atreus-server/app/admin/api/menu"
	"github.com/go-atreus/atreus-server/app/admin/api/organization"
	"github.com/go-atreus/atreus-server/app/admin/api/role"
	"github.com/go-atreus/atreus-server/app/admin/api/user"
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
)

type Data struct {
	ORM      *gorm.DB
	redisCli redis.Cmdable
}

func NewData(logger log.Logger) *Data {
	type source struct {
		Database struct {
			Driver      string
			Source      string
			Idle        int
			Active      int
			IdleTimeout string
		}
	}
	c := &source{}
	c.Database.Driver = "mysql"
	c.Database.Source = "root:root@tcp(127.0.0.1:3306)/atreus?charset=utf8mb4&parseTime=True&loc=Local"
	c.Database.Idle = 10
	c.Database.Active = 10
	c.Database.IdleTimeout = "10s"

	sqlDB, err := sql.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		//log.Error("sql dsn(%v) error: %v", c.Database.Source, err)
		panic(err)
	}
	sqlDB.SetMaxIdleConns(int(c.Database.Idle))
	sqlDB.SetMaxOpenConns(int(c.Database.Active))
	idleTimeout, _ := time.ParseDuration(c.Database.IdleTimeout)
	sqlDB.SetConnMaxLifetime(idleTimeout)

	orm, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger:                                   logger2.Default.LogMode(logger2.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Errorf("db dsn(%v) error: %v", c.Database.Source, err)
		panic(err)
	}

	// migrate
	err = orm.AutoMigrate(
		&user.SysUserORM{},
		&menu.SysMenuORM{},
		&dict.SysDictORM{},
		&organization.SysOrganizationORM{},
		&role.SysRoleORM{})
	if err != nil {
		panic(err)
	}

	return &Data{ORM: orm}
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
