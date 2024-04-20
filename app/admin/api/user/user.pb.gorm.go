package user

import (
	context "context"
	fmt "fmt"
	organization "github.com/go-atreus/atreus-server/app/admin/api/organization"
	role "github.com/go-atreus/atreus-server/app/admin/api/role"
	errors "github.com/infobloxopen/protoc-gen-gorm/errors"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	gorm "gorm.io/gorm"
	strings "strings"
)

type SysUserORM struct {
	Access            string
	Avatar            string
	Country           string
	Group             string
	Id                int32 `gorm:"type:integer;primaryKey;autoIncrement"`
	NickName          string
	NotifyCount       string
	Organization      *organization.SysOrganizationORM `gorm:"foreignKey:SysOrganizationId;references:Id"`
	ParentId          int32
	Password          string
	Role              []*role.SysRoleORM `gorm:"foreignKey:Id;references:Id;many2many:sys_user_role;joinForeignKey:SysUserId;joinReferences:SysRoleId"`
	SysOrganizationId *int32
	Tags              string
	Title             string
	UnreadCount       string
	Username          string `gorm:"unique"`
}

// TableName overrides the default tablename generated by GORM
func (SysUserORM) TableName() string {
	return "sys_users"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *SysUser) ToORM(ctx context.Context) (SysUserORM, error) {
	to := SysUserORM{}
	var err error
	if prehook, ok := interface{}(m).(SysUserWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Username = m.Username
	to.NickName = m.NickName
	to.Password = m.Password
	to.ParentId = m.ParentId
	to.Avatar = m.Avatar
	to.Title = m.Title
	to.Group = m.Group
	to.Tags = m.Tags
	to.NotifyCount = m.NotifyCount
	to.UnreadCount = m.UnreadCount
	to.Country = m.Country
	to.Access = m.Access
	for _, v := range m.Role {
		if v != nil {
			if tempRole, cErr := v.ToORM(ctx); cErr == nil {
				to.Role = append(to.Role, &tempRole)
			} else {
				return to, cErr
			}
		} else {
			to.Role = append(to.Role, nil)
		}
	}
	if m.Organization != nil {
		tempOrganization, err := m.Organization.ToORM(ctx)
		if err != nil {
			return to, err
		}
		to.Organization = &tempOrganization
	}
	if posthook, ok := interface{}(m).(SysUserWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *SysUserORM) ToPB(ctx context.Context) (SysUser, error) {
	to := SysUser{}
	var err error
	if prehook, ok := interface{}(m).(SysUserWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Username = m.Username
	to.NickName = m.NickName
	to.Password = m.Password
	to.ParentId = m.ParentId
	to.Avatar = m.Avatar
	to.Title = m.Title
	to.Group = m.Group
	to.Tags = m.Tags
	to.NotifyCount = m.NotifyCount
	to.UnreadCount = m.UnreadCount
	to.Country = m.Country
	to.Access = m.Access
	for _, v := range m.Role {
		if v != nil {
			if tempRole, cErr := v.ToPB(ctx); cErr == nil {
				to.Role = append(to.Role, &tempRole)
			} else {
				return to, cErr
			}
		} else {
			to.Role = append(to.Role, nil)
		}
	}
	if m.Organization != nil {
		tempOrganization, err := m.Organization.ToPB(ctx)
		if err != nil {
			return to, err
		}
		to.Organization = &tempOrganization
	}
	if posthook, ok := interface{}(m).(SysUserWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type SysUser the arg will be the target, the caller the one being converted from

// SysUserBeforeToORM called before default ToORM code
type SysUserWithBeforeToORM interface {
	BeforeToORM(context.Context, *SysUserORM) error
}

// SysUserAfterToORM called after default ToORM code
type SysUserWithAfterToORM interface {
	AfterToORM(context.Context, *SysUserORM) error
}

// SysUserBeforeToPB called before default ToPB code
type SysUserWithBeforeToPB interface {
	BeforeToPB(context.Context, *SysUser) error
}

// SysUserAfterToPB called after default ToPB code
type SysUserWithAfterToPB interface {
	AfterToPB(context.Context, *SysUser) error
}

// DefaultCreateSysUser executes a basic gorm create call
func DefaultCreateSysUser(ctx context.Context, in *SysUser, db *gorm.DB) (*SysUser, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SysUserORMWithBeforeCreate_); ok {
		if ormObj, db, err = hook.BeforeCreate_(ctx, db, ormObj); err != nil {
			return nil, err
		}
	}
	if err = db.Omit().Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SysUserORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type SysUserORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB, SysUserORM) (SysUserORM, *gorm.DB, error)
}
type SysUserORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadSysUser(ctx context.Context, in *SysUser, db *gorm.DB) (*SysUser, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == 0 {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(SysUserORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(SysUserORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := SysUserORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(SysUserORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type SysUserORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type SysUserORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type SysUserORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteSysUser(ctx context.Context, in *SysUser, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(SysUserORMWithBeforeDelete_); ok {
		if ormObj, db, err = hook.BeforeDelete_(ctx, db, ormObj); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&SysUserORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(SysUserORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type SysUserORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB, SysUserORM) (SysUserORM, *gorm.DB, error)
}
type SysUserORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteSysUserSet(ctx context.Context, in []*SysUser, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	var err error
	keys := []int32{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == 0 {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&SysUserORM{})).(SysUserORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&SysUserORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&SysUserORM{})).(SysUserORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type SysUserORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*SysUser, *gorm.DB) (*gorm.DB, error)
}
type SysUserORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*SysUser, *gorm.DB) error
}

// DefaultStrictUpdateSysUser clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateSysUser(ctx context.Context, in *SysUser, db *gorm.DB) (*SysUser, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateSysUser")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &SysUserORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(SysUserORMWithBeforeStrictUpdateCleanup); ok {
		if ormObj, db, err = hook.BeforeStrictUpdateCleanup(ctx, db, ormObj); err != nil {
			return nil, err
		}
	}
	if err = db.Model(&ormObj).Association("Role").Replace(ormObj.Role); err != nil {
		return nil, err
	}
	ormObj.Role = nil
	if hook, ok := interface{}(&ormObj).(SysUserORMWithBeforeStrictUpdateSave); ok {
		if ormObj, db, err = hook.BeforeStrictUpdateSave(ctx, db, ormObj); err != nil {
			return nil, err
		}
	}
	if err = db.Omit().Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SysUserORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type SysUserORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB, SysUserORM) (SysUserORM, *gorm.DB, error)
}
type SysUserORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB, SysUserORM) (SysUserORM, *gorm.DB, error)
}
type SysUserORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchSysUser executes a basic gorm update call with patch behavior
func DefaultPatchSysUser(ctx context.Context, in *SysUser, updateMask *field_mask.FieldMask, db *gorm.DB) (*SysUser, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj SysUser
	var err error
	if hook, ok := interface{}(&pbObj).(SysUserWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadSysUser(ctx, &SysUser{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(SysUserWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskSysUser(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(SysUserWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateSysUser(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(SysUserWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type SysUserWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *SysUser, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type SysUserWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *SysUser, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type SysUserWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *SysUser, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type SysUserWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *SysUser, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetSysUser executes a bulk gorm update call with patch behavior
func DefaultPatchSetSysUser(ctx context.Context, objects []*SysUser, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*SysUser, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*SysUser, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchSysUser(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskSysUser patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskSysUser(ctx context.Context, patchee *SysUser, patcher *SysUser, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*SysUser, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	var updatedOrganization bool
	for i, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"Username" {
			patchee.Username = patcher.Username
			continue
		}
		if f == prefix+"NickName" {
			patchee.NickName = patcher.NickName
			continue
		}
		if f == prefix+"Children" {
			patchee.Children = patcher.Children
			continue
		}
		if f == prefix+"Password" {
			patchee.Password = patcher.Password
			continue
		}
		if f == prefix+"ParentId" {
			patchee.ParentId = patcher.ParentId
			continue
		}
		if f == prefix+"Avatar" {
			patchee.Avatar = patcher.Avatar
			continue
		}
		if f == prefix+"Title" {
			patchee.Title = patcher.Title
			continue
		}
		if f == prefix+"Group" {
			patchee.Group = patcher.Group
			continue
		}
		if f == prefix+"Tags" {
			patchee.Tags = patcher.Tags
			continue
		}
		if f == prefix+"NotifyCount" {
			patchee.NotifyCount = patcher.NotifyCount
			continue
		}
		if f == prefix+"UnreadCount" {
			patchee.UnreadCount = patcher.UnreadCount
			continue
		}
		if f == prefix+"Country" {
			patchee.Country = patcher.Country
			continue
		}
		if f == prefix+"Access" {
			patchee.Access = patcher.Access
			continue
		}
		if f == prefix+"Permissions" {
			patchee.Permissions = patcher.Permissions
			continue
		}
		if f == prefix+"Roles" {
			patchee.Roles = patcher.Roles
			continue
		}
		if f == prefix+"Role" {
			patchee.Role = patcher.Role
			continue
		}
		if !updatedOrganization && strings.HasPrefix(f, prefix+"Organization.") {
			updatedOrganization = true
			if patcher.Organization == nil {
				patchee.Organization = nil
				continue
			}
			if patchee.Organization == nil {
				patchee.Organization = &organization.SysOrganization{}
			}
			if o, err := organization.DefaultApplyFieldMaskSysOrganization(ctx, patchee.Organization, patcher.Organization, &field_mask.FieldMask{Paths: updateMask.Paths[i:]}, prefix+"Organization.", db); err != nil {
				return nil, err
			} else {
				patchee.Organization = o
			}
			continue
		}
		if f == prefix+"Organization" {
			updatedOrganization = true
			patchee.Organization = patcher.Organization
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListSysUser executes a gorm list call
func DefaultListSysUser(ctx context.Context, db *gorm.DB) ([]*SysUser, error) {
	in := SysUser{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SysUserORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(SysUserORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []SysUserORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SysUserORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*SysUser{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type SysUserORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type SysUserORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type SysUserORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]SysUserORM) error
}
type UserDefaultServer struct {
	DB *gorm.DB
}

// ListSysUser ...
func (m *UserDefaultServer) ListSysUser(ctx context.Context, in *emptypb.Empty) (*ListUser, error) {
	db := m.DB
	if custom, ok := interface{}(in).(UserSysUserWithBeforeListSysUser); ok {
		var err error
		if db, err = custom.BeforeListSysUser(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultListSysUser(ctx, db)
	if err != nil {
		return nil, err
	}
	out := &ListUser{Results: res}
	if custom, ok := interface{}(in).(UserSysUserWithAfterListSysUser); ok {
		var err error
		if err = custom.AfterListSysUser(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// UserSysUserWithBeforeListSysUser called before DefaultListSysUserSysUser in the default ListSysUser handler
type UserSysUserWithBeforeListSysUser interface {
	BeforeListSysUser(context.Context, *gorm.DB) (*gorm.DB, error)
}

// UserSysUserWithAfterListSysUser called before DefaultListSysUserSysUser in the default ListSysUser handler
type UserSysUserWithAfterListSysUser interface {
	AfterListSysUser(context.Context, *ListUser, *gorm.DB) error
}
