package role

import (
	"context"
	"gorm.io/gorm"
)

func (o *SysRoleORM) BeforeStrictUpdateCleanup(ctx context.Context, db *gorm.DB, orm SysRoleORM) (SysRoleORM, *gorm.DB, error) {
	var res SysRoleORM
	if err := db.Model(&orm).Where("id=?", orm.Id).Preload("Menu").First(&res).Error; err != nil {
		return orm, db, err
	}
	orm.Menu = res.Menu
	return orm, db, nil
}
