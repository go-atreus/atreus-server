package user

import (
	"context"
	"gorm.io/gorm"
)

// AfterToPB implements the posthook interface for the User type. This allows
// us to customize conversion behavior. In this example, we set the User's Age
// based on the Birthday, instead of storing it separately in the DB
func (m *SysUserORM) AfterToPB(ctx context.Context, user *SysUser) error {
	user.Password = ""
	return nil
}

func (m *SysUser) AfterToORM(ctx context.Context, orm *SysUserORM) error {
	//orm.Password = "md5: " + orm.Password
	return nil
}

func (m *SysUserORM) BeforeCreate_(ctx context.Context, db *gorm.DB, orm SysUserORM) (SysUserORM, *gorm.DB, error) {
	orm.Password = "md5: " + orm.Password
	return orm, db, nil
}
