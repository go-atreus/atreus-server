package data

import (
	context "context"
	"errors"
	"fmt"
	"github.com/go-atreus/atreus-server/app/admin/api/user"
	"github.com/go-atreus/atreus-server/app/admin/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/singleflight"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
	sg   *singleflight.Group
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user")),
		sg:   &singleflight.Group{},
	}
}

func (u userRepo) FindByUsername(ctx context.Context, username string) (res *user.SysUserORM, err error) {
	result, err, _ := u.sg.Do(fmt.Sprintf("find_user_by_name_%s", username), func() (interface{}, error) {
		var findUser *user.SysUserORM
		err = u.data.ORM.Where("username = ?", username).First(&findUser).Error
		if err != nil {
			return nil, errors.New("用户不存在")
		}
		return findUser, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*user.SysUserORM), nil
}
